package OrderValidation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Middlewares"
	"github.com/vahidlotfi71/online-store-api/Rules"
)

func CreateOrderValidation() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "user_id",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Numeric()},
		},
		{
			FieldName: "total_price",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Numeric()},
		},
		{
			FieldName: "items",
			Rules:     []Rules.ValidationRule{Rules.Required},
		},
	})
}
