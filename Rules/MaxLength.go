package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MaxLength(max uint) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)

		message = fmt.Sprintf("The %s field must not exceed %d characters", field_name, max)

		if len(value) > int(max) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
