package OrderController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Resources/OrderResource"
)

func Index(c *fiber.Ctx) error {
	// دریافت کاربر از locals
	user, ok := c.Locals("user").(Models.User)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "User not found in context"})
	}

	// فقط سفارشات کاربر جاری
	tx := Config.DB.Table("orders").Where("deleted_at IS NULL AND user_id = ?", user.ID).Order("id DESC")

	var orders []Models.Order
	if err := tx.Preload("Items.Product").Find(&orders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": OrderResource.Collection(orders),
	})
}
