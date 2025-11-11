package ProductValidation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/internal/Middlewares"
	"github.com/vahidlotfi71/online-store-api.git/internal/Rules"
)

var CreateProduct = createProduct
var UpdateProduct = createProduct

func createProduct() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "name",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(2, 100)},
		},
		{
			FieldName: "brand",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(2, 100)},
		},
		{
			FieldName: "price",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Numeric()},
		},
		{
			FieldName: "description",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(10, 255)},
		},
		{
			FieldName: "stock",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Numeric()},
		},
	})
}
