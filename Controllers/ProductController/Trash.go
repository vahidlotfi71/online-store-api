package ProductController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Models/Product"
	"github.com/vahidlotfi71/online-store-api/Resources/ProductResource"
)

func Trash(c *fiber.Ctx) error {
	// ایجاد query builder با Model مشخص
	tx := Config.DB.Model(&Models.Product{}).Unscoped().Where("deleted_at IS NOT NULL").Order("id")

	products, meta, err := Product.Paginate(tx, c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"data":     ProductResource.Collection(products),
		"metadata": meta,
	})
}
