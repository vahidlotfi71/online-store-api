package Rules

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func BooleanStrict() ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)

		if value == "" {
			return true, "", nil, nil
		}

		// حذف فاصله‌های اضافی و تبدیل به حروف کوچک
		normalized := strings.ToLower(strings.TrimSpace(value))

		// فقط "true" و "false" قابل قبول هستند
		if normalized != "true" && normalized != "false" {
			message = fmt.Sprintf("The %s field must be either 'true' or 'false'", field_name)
			return false, message, nil, nil
		}

		return true, "", nil, nil
	}
}
