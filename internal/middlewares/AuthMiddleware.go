package Middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models"
	"github.com/vahidlotfi71/online-store-api.git/internal/Utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tokenSplit := strings.Split(token, " ")

	if len(tokenSplit) != 2 || strings.ToLower(tokenSplit[0]) != "bearer" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token type, expected bearer",
		})
	}
	userID, _, err := Utils.VerifyToken(tokenSplit[1])

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token",
		})

	}

	var user Models.User
	if err := Config.DB.First(&user, userID).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "User not found"})
	}

	c.Locals("user", user)
	return c.Next()
}
