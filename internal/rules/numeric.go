package Rules

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Numeric() ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)
		// اگر فیلد خالی باشد، بررسی نکن (با Required چک شود)
		if value == "" {
			return true, "", nil, nil
		}
		// برسی عدد بودن
		if _, err := strconv.Atoi(value); err != nil {
			message := fmt.Sprintf("The %s field must be a number", field_name)
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
