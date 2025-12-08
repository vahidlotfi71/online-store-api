// cmd/seed.go
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

var users uint

func init() {
	rootCmd.AddCommand(seedCmd)
	seedCmd.Flags().UintVarP(&users, "users", "u", 100, "number of fake users")
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database with fake users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting seed process...")

		if err := Config.Getenv(); err != nil {
			log.Fatalf("env: %v", err)
		}

		err := Config.Connect()
		if err != nil {
			log.Fatalf("db: %v", err)
		}

		fmt.Printf("Attempting to seed %d users...\n", users)
		seedUsers(Config.DB, users)
		fmt.Println("✅ Seeding done")
	},
}

func seedUsers(db *gorm.DB, n uint) {
	// تولید هش پسورد
	hash, err := Utils.GenerateHashPassword("password")
	if err != nil {
		fmt.Printf("Password hash error: %v\n", err)
		return
	}

	successCount := 0
	for i := 1; i <= int(n); i++ {
		// ساخت شماره تلفن منحصر به فرد
		phone := fmt.Sprintf("0912%07d", i) // تغییر به 0912 برای اطمینان از جدید بودن

		user := Models.User{
			FirstName:  fmt.Sprintf("User%d", i),
			LastName:   "Test",
			Phone:      phone,
			Address:    fmt.Sprintf("Address %d", i),
			NationalID: fmt.Sprintf("001%09d", i), // کد ملی منحصر به فرد
			Password:   hash,
			Role:       "user",
			IsVerified: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		// ایجاد کاربر
		result := db.Create(&user)
		if result.Error != nil {
			fmt.Printf("Failed to create user %d: %v\n", i, result.Error)
		} else {
			successCount++
			if successCount%10 == 0 {
				fmt.Printf("Created %d users so far...\n", successCount)
			}
		}
	}

	fmt.Printf("Seeding completed: %d users created successfully\n", successCount)

	// نمایش نمونه از کاربران ایجاد شده
	var sampleUsers []Models.User
	db.Limit(3).Find(&sampleUsers)
	fmt.Println("Sample users created:")
	for _, u := range sampleUsers {
		fmt.Printf("   - %s %s (%s)\n", u.FirstName, u.LastName, u.Phone)
	}
}
