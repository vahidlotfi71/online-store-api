// file: internal/Controllers/Admin/AdminController/Update.go
package AdminController

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models/Admin"
	"github.com/vahidlotfi71/online-store-api/Resources/AdminResource"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

type AdminUpdateRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	NationalID string `json:"national_id"`
	Password   string `json:"password"` // optional
}

func Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(400).JSON(fiber.Map{"message": "id param is required"})
	}
	num, _ := strconv.ParseUint(idStr, 10, 32)
	id := uint(num)

	var req AdminUpdateRequest
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

	if req.Password != "" {
		hashedPass, err := Utils.GenerateHashPassword(req.Password)
		if err != nil {
			tx.Rollback()
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
		}
		req.Password = hashedPass
	}

	dto := Admin.AdminUpdateDTO{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Address:    req.Address,
		NationalID: req.NationalID,
		Password:   req.Password,
	}

	if err := Admin.Update(tx, id, dto); err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// خواندن مجدد برای داشتن struct به‌روزرسانی‌شده
	updatedAdmin, err := Admin.FindByID(tx, id)
	if err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Commit failed"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Admin updated successfully",
		"data":    AdminResource.Single(updatedAdmin),
	})
}
