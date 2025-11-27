package Rules

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NationalID() ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		// Try to get value from JSON body first
		var jsonBody map[string]interface{}
		var value string

		if err := c.BodyParser(&jsonBody); err == nil {
			if val, exists := jsonBody[field_name]; exists {
				if strVal, ok := val.(string); ok {
					value = strVal
				}
			}
		}

		// If not found in JSON, try form value
		if value == "" {
			value = c.FormValue(field_name)
		}

		// If field is empty, skip validation (Required rule will check it)
		if value == "" {
			return true, "", nil, nil
		}

		// Check 10-digit format
		matched, _ := regexp.MatchString(`^\d{10}$`, value)
		if !matched {
			message := fmt.Sprintf("The %s field must be exactly 10 digits", field_name)
			return false, message, nil, nil
		}

		// Calculate control digit (checksum)
		sum := 0
		for i := 0; i < 9; i++ {
			d, _ := strconv.Atoi(string(value[i]))
			sum += d * (10 - i)
		}

		remainder := sum % 11
		checkDigit, _ := strconv.Atoi(string(value[9]))

		// Validate control digit
		if (remainder < 2 && checkDigit != remainder) ||
			(remainder >= 2 && checkDigit != (11-remainder)) {
			message := fmt.Sprintf("The %s field is not a valid national ID", field_name)
			return false, message, nil, nil
		}

		// If valid
		return true, "", nil, nil
	}
}
