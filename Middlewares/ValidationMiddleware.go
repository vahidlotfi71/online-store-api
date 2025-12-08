package Middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Providers"
	"github.com/vahidlotfi71/online-store-api/Rules"
)

func ValidationMiddleware(schema []Rules.FieldRules) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Printf("=== VALIDATION MIDDLEWARE START ===\n")
		fmt.Printf("Content-Type: %s\n", c.Get("Content-Type"))

		// Get multipart form to see all fields
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Printf("MultipartForm Error: %v\n", err)
		} else {
			fmt.Printf("MultipartForm Values: %+v\n", form.Value)
			fmt.Printf("MultipartForm Files: %+v\n", form.File)
		}

		// Check each field individually
		fmt.Printf("=== CHECKING INDIVIDUAL FIELDS ===\n")
		for _, field_rules := range schema {
			formValue := c.FormValue(field_rules.FieldName)
			queryValue := c.Query(field_rules.FieldName)
			paramValue := c.Params(field_rules.FieldName)

			fmt.Printf("Field: %s\n", field_rules.FieldName)
			fmt.Printf("  - FormValue: '%s'\n", formValue)
			fmt.Printf("  - Query: '%s'\n", queryValue)
			fmt.Printf("  - Params: '%s'\n", paramValue)

			if formValue == "" {
				fmt.Printf("FORM VALUE IS EMPTY!\n")
			}
		}

		body := c.Request().Body()
		if len(body) == 0 {
			fmt.Printf("REQUEST BODY IS EMPTY\n")
			return c.Status(400).JSON(fiber.Map{
				"message": "Empty request body",
			})
		}

		fmt.Printf("Raw Body: %s\n", string(body))

		for _, field_rules := range schema {
			fmt.Printf("Validating field: %s\n", field_rules.FieldName)
			fieldValue := c.FormValue(field_rules.FieldName)
			fmt.Printf("   Field value: '%s'\n", fieldValue)

			for i, rule := range field_rules.Rules {
				fmt.Printf("   Rule %d...\n", i+1)
				passed, message, flags, err := rule(c, field_rules.FieldName)
				if err != nil {
					fmt.Printf("   Rule error: %v\n", err)
					return c.Status(500).JSON(fiber.Map{
						"message": Providers.ErrorProvider(err),
					})
				}

				if passed && (flags != nil && flags.IsNull) {
					fmt.Printf("    Rule passed (is null)\n")
					break
				} else if !passed {
					fmt.Printf("    Validation FAILED: %s\n", message)
					return c.Status(400).JSON(fiber.Map{
						"message": message,
					})
				}
				fmt.Printf("    Rule passed\n")
			}
			fmt.Printf("   All rules passed for %s\n", field_rules.FieldName)
		}

		fmt.Printf("=== ALL VALIDATIONS PASSED ===\n")
		return c.Next()
	}
}
