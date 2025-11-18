package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models"
	"github.com/vahidlotfi71/online-store-api.git/Utils"
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
			fmt.Println("🗑️  Dropping tables ...")
			for _, m := range models {
				if err := Config.DB.Migrator().DropTable(m); err != nil {
					log.Fatalf("drop: %v", err)
				}
			}
			fmt.Println("✅ Tables dropped")
		}

		fmt.Println("🔨 Migrating ...")
		for _, m := range models {
			if err := Config.DB.AutoMigrate(m); err != nil {
				log.Fatalf("migrate: %v", err)
			}
		}

		// seed super-admin
		seedSuperAdmin(Config.DB)
		fmt.Println("✅ Migration & seed completed")
	},
}

func seedSuperAdmin(db *gorm.DB) {
	var admin Models.Admin
	hash, _ := Utils.GenerateHashPassword("12345678")
	db.FirstOrCreate(&admin, Models.Admin{Phone: "09123456789"},
		Models.Admin{
			FirstName:  "Super",
			LastName:   "Admin",
			Phone:      "09123456789",
			Address:    "Tehran",
			NationalID: "0000000000",
			Password:   hash,
			Role:       "admin",
			IsVerified: true,
		})
}
