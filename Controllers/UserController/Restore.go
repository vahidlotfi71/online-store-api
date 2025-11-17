// file: Controllers/UserController.go
package UserController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models"
	"github.com/vahidlotfi71/online-store-api.git/Resources/UserResource"
	"gorm.io/gorm"
)

// Restore بازگردانی نرم (Soft-Delete) بر اساس شناسه
func Restore(c *fiber.Ctx) error {
	// ۱) استخراج و اعتبارسنجی شناسه
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

	// ۲) چک وجود رکورد حذف‌شده
	var user Models.User
	if err := Config.DB.
		Where("id = ? AND deleted_at IS NOT NULL", id).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"message": "User not found or not deleted"})
		}
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": err.Error()})
	}

	// ۳) بازگردانی با ORM + RowsAffected
	result := Config.DB.
		Model(&Models.User{}).
		Where("id = ? AND deleted_at IS NOT NULL", id).
		Update("deleted_at", gorm.Expr("NULL"))

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"message": "User not found or not deleted"})
	}

	// ۴) خواندن دوباره مدل برای پاسخ
	if err := Config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": err.Error()})
	}

	// ۵) پاسخ موفق
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User restored successfully",
		"data":    UserResource.Single(user),
	})
}
