package UserController

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Resources/UserResource"
	"gorm.io/gorm"
)

// Restore بازگردانی نرم (Soft-Delete) بر اساس شناسه
func Restore(c *fiber.Ctx) error {
	//  استخراج و اعتبارسنجی شناسه
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "id param is required"})
	}
	num, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "id must be a positive integer"})
	}
	id := uint(num)

	//  جستجوی کاربر با Unscoped (شامل رکوردهای حذف شده)
	var user Models.User
	if err := Config.DB.Unscoped().
		Where("id = ?", id).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": err.Error()})
	}

	//  بررسی اینکه کاربر واقعاً حذف شده باشد
	if user.DeletedAt.Time.IsZero() { // استفاده از متد IsZero()
		log.Printf("User with ID %d is not deleted (DeletedAt is zero)", id)
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "User is not deleted"})
	}

	//  بازگردانی با Unscoped
	result := Config.DB.Unscoped().
		Model(&Models.User{}).
		Where("id = ?", id).
		Update("deleted_at", nil)

	if result.Error != nil {
		log.Printf("Failed to restore user: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		log.Printf("No rows affected when restoring user with ID %d", id)
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"message": "User not found"})
	}

	//  خواندن دوباره مدل برای پاسخ
	if err := Config.DB.First(&user, id).Error; err != nil {
		log.Printf("Failed to reload user after restore: %v", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": err.Error()})
	}

	//  پاسخ موفق
	log.Printf("User with ID %d restored successfully", id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User restored successfully",
		"data":    UserResource.Single(user),
	})
}
