// file: Controllers/UserController.go
package UserController

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models"
	"github.com/vahidlotfi71/online-store-api.git/Models/User"
	"github.com/vahidlotfi71/online-store-api.git/Resources/UserResource"
	"github.com/vahidlotfi71/online-store-api.git/Utils"
	"gorm.io/gorm"
)

/* ---------- DTO ---------- */
type UserUpdateRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	NationalID string `json:"national_id"`
	Password   string `json:"password"` // اگر خالی باشد آپدیت نمی‌شود
}

/* ---------- ویرایش کاربر (بدون آپلود فایل) ---------- */
func Update(c *fiber.Ctx) error {
	// ۱) شناسه از مسیر
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "id param is required"})
	}
	num, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "id must be a positive integer"})
	}
	id := uint(num)

	// ۲) Parsing JSON
	var req UserUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	// ۳) شروع تراکنش
	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "DB connection error"})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ۴) چک وجود کاربر (حذف‌نشده)
	var user Models.User
	if err := tx.Where("deleted_at IS NULL").First(&user, id).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// ۵) هش پسورد (در صورت ارسال)
	if req.Password != "" {
		req.Password, err = Utils.GenerateHashPassword(req.Password)
		if err != nil {
			tx.Rollback()
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
		}
	}

	// ۶) ساخت DTO برای Repository
	dto := User.UserUpdateDTO{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Address:    req.Address,
		NationalID: req.NationalID,
		Password:   req.Password, // تصویر نداریم
	}

	// ۷) به‌روزرسانی در DB
	if err := User.Update(tx, id, dto); err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// ۸) کامیت موفق
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Commit failed"})
	}

	// ۹) پاسخ استاندارد (با مدل به‌روزرسانی‌شده)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    UserResource.Single(user),
	})
}
