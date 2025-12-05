package Rules

import "github.com/gofiber/fiber/v2"

func Optional() ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {

		// برای Form Data از c.FormValue استفاده می‌کنیم
		value := c.FormValue(field_name)

		// اگر مقدار فیلد خالی بود
		if value == "" {
			// IsNull را برمی‌گردانیم تا میدل‌ور اعتبارسنجی‌های بعدی را متوقف کند.
			return true, "", &Flags{IsNull: true}, nil
		}

		// اگر مقدار خالی نبود، اعتبارسنجی ادامه می‌یابد
		return true, "", nil, nil
	}
}
