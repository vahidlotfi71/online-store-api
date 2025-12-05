package Utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// تابع کمکی برای تشخیص نقش admin از JWT
func IsAdmin(c *fiber.Ctx) bool {
	token := c.Get("Authorization")

	if token == "" {
		return false
	}

	tokenSplit := strings.Split(token, " ")

	if len(tokenSplit) != 2 || strings.ToLower(tokenSplit[0]) != "bearer" {
		return false
	}

	claims, err := VerifyToken(tokenSplit[1])
	if err != nil {
		return false
	}

	// بررسی نقش
	return claims.Role == "admin"
}
