// file: Controllers/UserController.go
package UserController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models"
)

// ClearTrash حذف فیزیکی دسته‌ای کاربران حذف‌شده (بدون تصویر)
func ClearTrash(c *fiber.Ctx) error {
	// ۱) خواندن تعداد درخواستی (با سقف امن)
	limitStr := c.Query("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 50 {
		limit = 10
	}

	// ۲) شروع تراکنش
	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": "DB connection error"})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ۳) خواندن رکوردهای حذف‌شده
	var users []Models.User
	if err := tx.Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("deleted_at ASC").
		Limit(limit).
		Find(&users).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": err.Error()})
	}

	// ۴) حذف فیزیکی دسته‌ای
	result := tx.Unscoped().Delete(&Models.User{}, "deleted_at IS NOT NULL AND id IN (?)",
		tx.Model(&Models.User{}).Select("id").Where("deleted_at IS NOT NULL").Limit(limit))
	if result.Error != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": result.Error.Error()})
	}

	// ۵) کامیت موفق
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": "Commit failed"})
	}

	// ۶) پاسخ موفق
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Trash cleared successfully",
		"cleared_count": result.RowsAffected,
	})
}
