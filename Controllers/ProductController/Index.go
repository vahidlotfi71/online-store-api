package ProductController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Resources/ProductResource"
	"github.com/vahidlotfi71/online-store-api/Utils/Http"
)

func Index(c *fiber.Ctx) error {
	var products []Models.Product
	tx := Config.DB.Table("products").Where("deleted_at IS NULL")
	tx.Order("id")

	var metadata Http.PaginationMetadata
	tx, metadata = Http.Paginate(tx, c)
	tx.Find(&products)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":     ProductResource.Collection(products),
		"metadata": metadata,
	})
}
