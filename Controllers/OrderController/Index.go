package OrderController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models/Order"
	"github.com/vahidlotfi71/online-store-api/Resources/OrderResource"
)

func Index(c *fiber.Ctx) error {
	orders, meta, err := Order.Paginate(Config.DB.Where("deleted_at IS NULL").Order("id"), c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"data":     OrderResource.Collection(orders),
		"metadata": meta,
	})
}
