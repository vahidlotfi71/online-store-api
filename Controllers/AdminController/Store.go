// file: internal/Controllers/Admin/AdminController/Store.go
package AdminController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models/Admin"
	"github.com/vahidlotfi71/online-store-api.git/Resources/AdminResource"
	"github.com/vahidlotfi71/online-store-api.git/Utils"
)

type AdminCreateRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	NationalID string `json:"national_id"`
	Password   string `json:"password"`
}

func Store(c *fiber.Ctx) error {
	var req AdminCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "DB connection error"})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	hashedPass, err := Utils.GenerateHashPassword(req.Password)
	if err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
	}

	dto := Admin.AdminCreateDTO{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Address:    req.Address,
		NationalID: req.NationalID,
		Password:   hashedPass,
	}

	admin, err := Admin.Create(tx, dto)
	if err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Commit failed"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Admin created successfully",
		"data":    AdminResource.Single(admin),
	})
}
