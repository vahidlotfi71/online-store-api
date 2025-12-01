// debug_migrate.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Utils"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("🔍 DEBUG MIGRATE AND SEED")

	// 1. Load env
	fmt.Println("1. Loading environment...")
	if err := Config.Getenv(); err != nil {
		log.Fatalf("❌ ENV ERROR: %v", err)
	}
	fmt.Println("✅ Environment loaded")

	// 2. Connect to DB
	fmt.Println("2. Connecting to database...")
	if err := Config.Connect(); err != nil {
		log.Fatalf("❌ DB ERROR: %v", err)
	}
	fmt.Println("✅ Database connected")

	// 3. Show all tables
	fmt.Println("3. Checking existing tables...")
	var tables []string
	Config.DB.Raw("SHOW TABLES").Scan(&tables)
	fmt.Printf("📊 Existing tables: %v\n", tables)

	// 4. Check if admin table exists
	fmt.Println("4. Checking admin table...")
	if Config.DB.Migrator().HasTable(&Models.Admin{}) {
		fmt.Println("✅ Admin table exists")
	} else {
		fmt.Println("❌ Admin table does NOT exist")
	}

	// 5. Try to create admin directly
	fmt.Println("5. Creating admin directly...")
	createAdminDirectly(Config.DB)

	// 6. Check if admin was created
	fmt.Println("6. Verifying admin creation...")
	var adminCount int64
	Config.DB.Model(&Models.Admin{}).Count(&adminCount)
	fmt.Printf("📊 Total admins in database: %d\n", adminCount)

	if adminCount > 0 {
		var admins []Models.Admin
		Config.DB.Find(&admins)
		fmt.Println("📋 Admin details:")
		for i, admin := range admins {
			fmt.Printf("   %d. ID: %d, Phone: %s, Name: %s %s\n", i+1, admin.ID, admin.Phone, admin.FirstName, admin.LastName)
		}
	}
}

func createAdminDirectly(db *gorm.DB) {
	hash, err := Utils.GenerateHashPassword("12345678")
	if err != nil {
		fmt.Printf("❌ Password hash error: %v\n", err)
		return
	}

	admin := Models.Admin{
		FirstName:  "Super",
		LastName:   "Admin",
		Phone:      "09123456789",
		Address:    "Tehran",
		NationalID: "0000000000",
		Password:   hash,
		Role:       "admin",
		IsVerified: true,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}

	fmt.Printf("🔨 Attempting to create admin: %s %s (%s)\n", admin.FirstName, admin.LastName, admin.Phone)

	// Try to create
	result := db.Create(&admin)
	if result.Error != nil {
		fmt.Printf("❌ CREATE ERROR: %v\n", result.Error)
	} else {
		fmt.Printf("✅ ADMIN CREATED SUCCESSFULLY! Rows affected: %d, Admin ID: %d\n",
			result.RowsAffected, admin.ID)
	}
}
