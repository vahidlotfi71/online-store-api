package AuthController

import "github.com/gofiber/fiber/v2"

func AdminLogout(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "You have been logged out",
	})
}
