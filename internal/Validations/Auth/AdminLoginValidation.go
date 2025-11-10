package Auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/internal/Middlewares"
	"github.com/vahidlotfi71/online-store-api.git/internal/Rules"
)

func AdminLogin() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "phone",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.PhoneNumber()},
		},
		{
			FieldName: "password",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.MinLength(8)},
		},
	})
}
