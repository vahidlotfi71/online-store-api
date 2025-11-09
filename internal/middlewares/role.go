package middlewares

import "github.com/gofiber/fiber/v2"

func RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// در این خط نقش کاربر فعلی (که از JWT در AuthMiddleware استخراج شده بود) از Context خوانده می‌شود.
		if c.Locals("role") != role {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"success": false, "message": "You do not have permission to access this section"})
		}
		return c.Next()
	}
}
