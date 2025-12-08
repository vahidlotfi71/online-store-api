// cmd/migrate.go
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Utils"
	"gorm.io/gorm"
)

var force bool

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().BoolVarP(&force, "force", "f", false, "drop all tables before migrating")
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Create / re-create MySQL tables",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Config.Getenv(); err != nil {
			log.Fatalf("env load: %v", err)
		}

		err := Config.Connect()
		if err != nil {
			log.Fatalf("db connect: %v", err)
		}

		models := []interface{}{
			&Models.User{},
			&Models.Admin{},
			&Models.Product{},
			&Models.Order{},
			&Models.OrderItem{},
		}

		if force {
			fmt.Println("Dropping tables ...")
			for _, m := range models {
				if err := Config.DB.Migrator().DropTable(m); err != nil {
					log.Fatalf("drop: %v", err)
				}
			}
			fmt.Println("Tables dropped")
		}

		fmt.Println("Migrating ...")
		for _, m := range models {
			if err := Config.DB.AutoMigrate(m); err != nil {
				log.Fatalf("migrate: %v", err)
			}
		}

		// seed super-admin
		seedSuperAdmin(Config.DB)
		fmt.Println("Migration & seed completed")
	},
}

// cmd/migrate.go
func seedSuperAdmin(db *gorm.DB) {
	fmt.Println("Seeding super admin...")

	hash, err := Utils.GenerateHashPassword("12345678")
	if err != nil {
		fmt.Printf("Admin password hash error: %v\n", err)
		return
	}

	admin := Models.Admin{
		FirstName:  "Super",
		LastName:   "Admin",
		Phone:      "09123456789",
		Address:    "Tehran",
		NationalID: "1111111111",
		Password:   hash,
		Role:       "admin",
		IsVerified: true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	db.Unscoped().Where("phone = ?", "09123456789").Or("national_id = ?", "1111111111").Delete(&Models.Admin{})

	if err := db.Create(&admin).Error; err != nil {
		fmt.Printf("Failed to create admin: %v\n", err)
	} else {
		fmt.Printf("Super Admin CREATED!\n")
		fmt.Printf(" 	Phone: 09123456789\n")
		fmt.Printf("   	Password: 12345678\n")
		fmt.Printf("   	National ID: 1111111111\n")
	}
}
