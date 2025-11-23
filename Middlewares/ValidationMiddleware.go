package Middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Providers"
	"github.com/vahidlotfi71/online-store-api/Rules"
)

func ValidationMiddleware(schema []Rules.FieldRules) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		body := c.Request().Body()
		if len(body) == 0 {
			return c.Status(400).JSON(fiber.Map{
				"message": "Empty form-data request",
			})
		}

		for _, field_rules := range schema {
			for _, rule := range field_rules.Rules {
				passed, message, flags, err := rule(c, field_rules.FieldName)
				if err != nil {
					return c.Status(500).JSON(fiber.Map{
						"message": Providers.ErrorProvider(err),
					})
				}

				if passed && (flags != nil && flags.IsNull) {
					break
				} else if !passed {
					return c.Status(400).JSON(fiber.Map{
						"message": message,
					})
				}
			}
		}
		return c.Next()
	}
}
