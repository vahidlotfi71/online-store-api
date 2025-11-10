package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LengthBetween(min, max uint) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)

		message = fmt.Sprintf("The %s field must be between %d  and %d characters", field_name, min, max)

		if len(value) > int(max) || len(value) < int(min) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
