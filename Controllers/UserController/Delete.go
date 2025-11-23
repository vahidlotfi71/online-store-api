// file: Controllers/UserController.go
package UserController

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"gorm.io/gorm"
)

// Delete انجام Soft-Delete بر اساس شناسه
func Delete(c *fiber.Ctx) error {
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

	// ۲) چک وجود رکورد (فقط حذف‌نشده‌ها)
	var user Models.User
	if err := Config.DB.Where("deleted_at IS NULL").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": err.Error()})
	}

	// ۳) Soft-Delete با ORM + RowsAffected
	result := Config.DB.Model(&Models.User{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", time.Now())

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"message": "User not found"})
	}

	// ۴) پاسخ موفق
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
