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

var users uint

func init() {
	rootCmd.AddCommand(seedCmd)
	seedCmd.Flags().UintVarP(&users, "users", "u", 100, "number of fake users")
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database with fake users",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Config.Getenv(); err != nil {
			log.Fatalf("env: %v", err)
		}
		err := Config.Connect()
		if err != nil {
			log.Fatalf("db: %v", err)
		}
		seedUsers(Config.DB, users)
		fmt.Println("✅ seeding done")
	},
}

func seedUsers(db *gorm.DB, n uint) {
	hash, _ := Utils.GenerateHashPassword("password")
	for i := 1; i <= int(n); i++ {
		u := Models.User{
			FirstName:  fmt.Sprintf("User%d", i),
			LastName:   "Fake",
			Phone:      fmt.Sprintf("0911%07d", i),
			Address:    "Tehran",
			NationalID: fmt.Sprintf("%010d", i),
			Password:   hash,
			Role:       "user",
			IsVerified: true,
		}
		db.FirstOrCreate(&u, Models.User{Phone: u.Phone})
	}
}
