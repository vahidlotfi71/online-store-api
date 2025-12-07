package Rules

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Numeric() ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)

		if value == "" {
			return true, "", nil, nil // فیلد اختیاری
		}

		num, err := strconv.Atoi(value)
		if err != nil {
			return false, fmt.Sprintf("The %s field must be a number", field_name), nil, nil
		}

		if num < 0 {
			message := fmt.Sprintf("The %s field must be greater than or equal to 0", field_name)
			return false, message, nil, nil
		}

		return true, "", nil, nil
	}
}
