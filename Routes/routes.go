package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vahidlotfi71/online-store-api.git/config"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	// Allow requests from different addresses (CORS is enabled).
	// Without this line, some requests from the browser may give access errors.
	// This line applies the CORS middleware to all routes.
	app.Use(cors.New())
	api := app.Group("/api")
	setupUserRoutes(api, db, cfg)
	setupAdminRoutes(api, db, cfg)
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "message": "سرور در حال اجراست"})
	})
}
