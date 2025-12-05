package Order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Middlewares"
	"github.com/vahidlotfi71/online-store-api/Rules"
)

func StoreOrder() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "items",
			// فقط چک می کنیم که فیلد items حتماً وجود داشته باشد.
			Rules: []Rules.ValidationRule{Rules.Required},
		},
	})
}
