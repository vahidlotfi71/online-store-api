package OrderController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Resources/OrderResource"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models"
	"gorm.io/gorm"
)

func Show(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id param is required"})
	}
	num, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id must be a positive integer"})
	}
	id := uint(num)

	var order Models.Order
	if err := Config.DB.Where("deleted_at IS NULL").Preload("Items.Product").First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Order not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": OrderResource.Single(order)})
}
