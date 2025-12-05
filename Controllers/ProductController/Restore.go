package ProductController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"gorm.io/gorm"
)

func Restore(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id param is required"})
	}
	num, _ := strconv.ParseUint(idStr, 10, 32)
	id := uint(num)

	var product Models.Product
	if err := Config.DB.Unscoped().Where("deleted_at IS NOT NULL").First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Product not found or not deleted"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := Config.DB.Unscoped().Model(&product).Update("deleted_at", nil).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product restored successfully"})
}
