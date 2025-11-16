package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/config"
	"github.com/vahidlotfi71/online-store-api.git/internal/models"
	"github.com/vahidlotfi71/online-store-api.git/internal/routes"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err := db.AutoMigrate(
		&models.Order{},
		&models.OrderItem{},
		&models.Product{},
		&models.User{}); err != nil {
		log.Fatal("Migration error:", err)
	}

	seedAdmin(db, cfg)

	app := fiber.New()
	routes.SetupRoutes(app, db, cfg)

	log.Fatal(app.Listen(":" + cfg.Port))
}

func seedAdmin(db *gorm.DB, cfg *config.Config) {
	var admin models.User

	// Go to the users table and search for a record that meets the condition.
	// If found, return that same record.
	// If not found, create a new record with the new data you provided.
	db.FirstOrCreate(
		&admin,
		models.User{Phone: "09123456789"},
		models.User{
			FirstName:  "vahid",
			LastName:   "lotfi",
			Phone:      "09123456789",
			Address:    "Tehran",
			NationalID: "0000000000",
			Password:   "$2a$10$3TslMTzXGq6TggHbF2BbqOkhMuhSuQeMY29upAwo9eJ54eOzrd4nm", //The password must be hashed.
			Role:       "admin",
			IsVerified: true,
		},
	)
}
