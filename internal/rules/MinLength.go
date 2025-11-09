package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MinLength(min uint) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)

		message = fmt.Sprintf("The %s field must be at least %d characters long", field_name, min)

		if len(value) < int(min) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
