package Middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models"
	"github.com/vahidlotfi71/online-store-api.git/Utils"
)

func AdminAuthMiddleware(c *fiber.Ctx) error {
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

	claims, err := Utils.VerifyToken(tokenSplit[1])
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// بررسی نقش ادمین
	if claims.Role != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"message": "Access denied. Admin role required",
		})
	}

	var admin Models.Admin
	if err := Config.DB.First(&admin, claims.ID).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Admin not found"})
	}

	c.Locals("admin", admin)
	return c.Next()
}
