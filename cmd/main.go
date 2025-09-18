package main

import (
	"log"

	"github.com/vahidlotfi71/online-store-api.git/config"
)

func main() {
	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
}
