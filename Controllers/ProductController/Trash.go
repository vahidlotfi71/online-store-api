package ProductController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Resources/ProductResource"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models/Product"
)

func Trash(c *fiber.Ctx) error {
	product, meta, err := Product.Paginate(
		Config.DB.Where("deleted_at IS NOT NULL").Order("id"),
		c,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"data":     ProductResource.Collection(product),
		"metadata": meta,
	})
}
