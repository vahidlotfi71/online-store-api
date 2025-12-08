package AuthController

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

type RegisterResponse struct {
	Token      string      `json:"token"`
	ExpireTime time.Time   `json:"expire_time"`
	User       Models.User `json:"user"`
}

// Register user registration
func Register(c *fiber.Ctx) error {
	fmt.Printf("=== REGISTER START ===\n")

	// Start database transaction
	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Database connection error"})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Use FormValue for all fields to handle UTF-8 characters properly
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	phone := c.FormValue("phone")
	address := c.FormValue("address")
	nationalID := c.FormValue("national_id")
	password := c.FormValue("password")

	// Hash password
	hashedPass, err := Utils.GenerateHashPassword(password)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"message": "Password encryption error"})
	}

	// Create user model
	now := time.Now()
	user := Models.User{
		FirstName:  firstName,
		LastName:   lastName,
		Phone:      phone,
		Address:    address,
		NationalID: nationalID,
		Password:   hashedPass,
		Role:       "user",
		IsVerified: false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// Create user with error handling
	if err := tx.Create(&user).Error; err != nil {
		fmt.Printf("Database Error: %v\n", err)
		tx.Rollback()

		// Detect duplicate errors
		if isDuplicateError(err) {
			if contains(err.Error(), "national_id") {
				return c.Status(400).JSON(fiber.Map{"message": "National id already registered"})
			}
			if contains(err.Error(), "phone") {
				return c.Status(400).JSON(fiber.Map{"message": "Phone number already registered"})
			}
			return c.Status(400).JSON(fiber.Map{"message": "User already exists"})
		}

		return c.Status(500).JSON(fiber.Map{"message": "User registration error: " + err.Error()})
	}

	// Generate JWT token
	token, expireTime, err := Utils.CreateToken(
		user.ID,
		user.Role,
		user.FirstName+" "+user.LastName,
		user.Phone,
		false,
	)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"message": "Token generation error"})
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"message": "Transaction commit error"})
	}

	fmt.Printf("=== USER CREATED SUCCESSFULLY ===\n")
	fmt.Printf("User ID: %d\n", user.ID)

	// Success response
	return c.Status(200).JSON(RegisterResponse{
		Token:      token,
		ExpireTime: expireTime,
		User:       user,
	})
}

// Helper function to detect duplicate errors
func isDuplicateError(err error) bool {
	return err != nil && (contains(err.Error(), "Duplicate") ||
		contains(err.Error(), "unique") ||
		contains(err.Error(), "1062"))
}

// Helper function to check substring
func contains(s, substr string) bool {
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
