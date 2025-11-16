package Rules

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NationalID() ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)

		// اگر فیلد خالی بود، بررسی نشود (با Required چک می‌شود)
		if value == "" {
			return true, "", nil, nil
		}

		// بررسی فرمت ۱۰ رقمی
		matched, _ := regexp.MatchString(`^\d{10}$`, value)
		if !matched {
			message := fmt.Sprintf("The %s field must be exactly 10 digits", field_name)
			return false, message, nil, nil
		}

		// محاسبه رقم کنترل (checksum)
		sum := 0
		for i := 0; i < 9; i++ {
			d, _ := strconv.Atoi(string(value[i]))
			sum += d * (10 - i)
		}

		remainder := sum % 11
		checkDigit, _ := strconv.Atoi(string(value[9]))

		// بررسی صحت رقم کنترلی
		if (remainder < 2 && checkDigit != remainder) ||
			(remainder >= 2 && checkDigit != (11-remainder)) {
			message := fmt.Sprintf("The %s field is not a valid national ID", field_name)
			return false, message, nil, nil
		}

		// در صورت معتبر بودن
		return true, "", nil, nil
	}
}
