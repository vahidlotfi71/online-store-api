package UserController

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Resources/UserResource"
	"gorm.io/gorm"
)

// Show اطلاعات یک کاربر بر اساس شناسه
func Show(c *fiber.Ctx) error {
	//  خواندن و اعتبارسنجی شناسه از مسیر
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

	//  جستجوی رکورد (فقط حذف‌نشده‌ها)
	var user Models.User
	if dbErr := Config.DB.Where("deleted_at IS NULL").First(&user, id).Error; dbErr != nil {
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"message": dbErr.Error()})
	}

	//  پاسخ موفق
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": UserResource.Single(user),
	})
}
