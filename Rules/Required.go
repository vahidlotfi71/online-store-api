package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Required(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
	// Try to get value from JSON body first
	var jsonBody map[string]interface{}
	if err := c.BodyParser(&jsonBody); err == nil {
		if value, exists := jsonBody[field_name]; exists && value != "" {
			return true, "", nil, nil
		}
	}

	// Fallback to form value for form-data
	value := c.FormValue(field_name)
	message = fmt.Sprintf("The %s field is required", field_name)

	if value == "" {
		return false, message, nil, nil
	}

	return true, "", nil, nil
}
