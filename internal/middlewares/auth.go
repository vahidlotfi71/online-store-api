package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vahidlotfi71/online-store-api.git/config"
)

func AuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization") //Check for the presence of the Authorization header
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "توکن ارسال نشده است"})
		}
		//The auth value is similar to this:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...
		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		//Token parsing and validation
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(cfg.JWT.Secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "توکن نامعتبر است"})
		}
		// Extracting information from token (Claims)
		claims := token.Claims.(jwt.MapClaims)
		// It stores the token information in the Context.
		c.Locals("userID", uint(claims["user_id"].(float64)))
		c.Locals("phone", claims["phone"].(string))
		c.Locals("role", claims["role"].(string))
		return c.Next()
	}
}
