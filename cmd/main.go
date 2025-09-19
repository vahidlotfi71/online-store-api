package main

import (
	"log"

	"github.com/vahidlotfi71/online-store-api.git/config"
	"github.com/vahidlotfi71/online-store-api.git/internal/models"
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
}
