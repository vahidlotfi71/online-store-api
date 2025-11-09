package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/internal/rules"
)

// ValidationMiddleware بررسی می‌کند داده‌های ورودی با قوانین اعتبارسنجی مطابقت دارند یا خیر
func ValidationMiddleware(fieldRules []rules.FieldRules) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// خواندن بدنه درخواست
		body := make(map[string]interface{})
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Data format is not valid",
			})
		}

		// چک کردن خالی بودن داده
		if len(body) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "No data sent",
			})
		}

		// بررسی قوانین اعتبارسنجی
		var errs []rules.ValidationError
		for _, fr := range fieldRules {
			val := ""
			if v, ok := body[fr.Field].(string); ok {
				val = v
			}

			for _, rule := range fr.Rules {
				ok, msg, err := rule(val, fr.Field)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"success": false,
						"message": "خطای سرور",
					})
				}
				if !ok {
					errs = append(errs, rules.ValidationError{
						Field:   fr.Field,
						Message: msg,
					})
					break
				}
			}
		}

		// اگر خطایی واحد داشت
		if len(errs) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"errors":  errs,
			})
		}

		// ادامه
		return c.Next()
	}
}
