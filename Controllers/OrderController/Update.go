package OrderController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models"
	"github.com/vahidlotfi71/online-store-api.git/Resources/OrderResource"
	"gorm.io/gorm"
)

type OrderUpdateRequest struct {
	Status *string `json:"status"`
}

func Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id param is required"})
	}
	num, _ := strconv.ParseUint(idStr, 10, 32)
	id := uint(num)

	var req OrderUpdateRequest
	if err := c.BodyParser(&req); err != nil || req.Status == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	var order Models.Order
	if err := Config.DB.Where("deleted_at IS NULL").First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Order not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := Config.DB.Model(&order).Update("status", *req.Status).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order status updated",
		"data":    OrderResource.Single(order),
	})
}
