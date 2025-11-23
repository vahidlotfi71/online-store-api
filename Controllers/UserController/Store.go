// file: Controllers/UserController.go
package UserController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models/User"
	"github.com/vahidlotfi71/online-store-api/Resources/UserResource"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

/* ---------- DTO (فقط ساختار،) ---------- */
type UserCreateRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	NationalID string `json:"national_id"`
	Password   string `json:"password"`
}

// ایجاد کابر
func Store(c *fiber.Ctx) error {
	// ۱) Parsing JSON
	var req UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	// ۲) شروع تراکنش
	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "DB connection error"})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ۳) هش پسورد
	hashedPass, err := Utils.GenerateHashPassword(req.Password)
	if err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
	}

	// ۴) ساخت DTO برای Repository
	dto := User.UserCreateDTO{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Address:    req.Address,
		NationalID: req.NationalID,
		Password:   hashedPass,
	}

	// ۵) درج در DB
	user, err := User.Create(tx, dto)
	if err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// ۶) کامیت موفق
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Commit failed"})
	}

	// ۷) پاسخ استاندارد
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    UserResource.Single(user),
	})
}
