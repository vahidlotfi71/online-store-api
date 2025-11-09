package Rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func PhoneNumber() ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)

		if value == "" {
			return true, "", nil, nil
		}

		cleaned := cleanPhoneNumber(value)

		// الگوهای مختلف شماره تلفن ایران
		patterns := []string{
			`^09[0-9]{9}$`,    // 09123456789
			`^9[0-9]{9}$`,     // 9123456789 (بدون صفر)
			`^\+989[0-9]{9}$`, // +989123456789
			`^00989[0-9]{9}$`, // 00989123456789
		}

		for _, pattern := range patterns {
			matched, _ := regexp.MatchString(pattern, cleaned)
			if matched {
				return true, "", nil, nil
			}
		}

		message = fmt.Sprintf("شماره موبایل وارد شده معتبر نیست. فرمت صحیح: 09123456789")
		return false, message, nil, nil
	}
}

// تابع کمکی برای پاکسازی شماره تلفن
func cleanPhoneNumber(phone string) string {
	// حذف فاصله، خط تیره و پیش شماره +98
	cleaned := strings.ReplaceAll(phone, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	cleaned = strings.ReplaceAll(cleaned, "+", "")
	cleaned = strings.ReplaceAll(cleaned, "(", "")
	cleaned = strings.ReplaceAll(cleaned, ")", "")

	// تبدیل پیش شماره 989 به 09
	if strings.HasPrefix(cleaned, "989") && len(cleaned) == 11 {
		cleaned = "09" + cleaned[3:]
	}

	// تبدیل پیش شماره به 09
	if strings.HasPrefix(cleaned, "989") && len(cleaned) == 12 {
		cleaned = "09" + cleaned[3:]
	}

	return cleaned
}
