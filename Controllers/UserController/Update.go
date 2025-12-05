// FILE: Controllers/UserController/Update.go
package UserController

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Models/User"
	"github.com/vahidlotfi71/online-store-api/Resources/UserResource"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

/* ---------- DTO ---------- */
type UserUpdateRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	NationalID string `json:"national_id"`
	Password   string `json:"password"` // optional
}

/* ---------- ویرایش کاربر (ادمین یا خود کاربر) ---------- */
func Update(c *fiber.Ctx) error {
	fmt.Printf(">>> UserController.Update: auth=%s\n", c.Get("Authorization"))

	// ۱) تشخیص نقش از طریق Locals
	var (
		role   string
		userID uint
	)
	if admin, ok := c.Locals("admin").(Models.Admin); ok {
		role = "admin"
		userID = admin.ID
	} else if user, ok := c.Locals("user").(Models.User); ok {
		role = "user"
		userID = user.ID
	} else {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	// ۲) تعیین شناسه کاربر هدف
	var targetID uint
	if role == "user" {
		// برای کاربر عادی، فقط پروفایل خودش قابل ویرایش است
		targetID = userID
	} else {
		// برای ادمین، شناسه از مسیر خوانده می‌شود
		idStr := c.Params("id")
		if idStr == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "ID parameter is required"})
		}
		num, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
		}
		targetID = uint(num)
	}

	// ۳) اگر کاربر بود، فقط اجازه ویرایش خودش را دارد
	if role == "user" && targetID != userID {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"message": "You can only update your own profile"})
	}

	// ۴) خواندن داده‌های ورودی
	var req UserUpdateRequest
	req.FirstName = strings.TrimSpace(c.FormValue("first_name"))
	req.LastName = strings.TrimSpace(c.FormValue("last_name"))
	req.Phone = strings.TrimSpace(c.FormValue("phone"))
	req.Address = strings.TrimSpace(c.FormValue("address"))
	req.NationalID = strings.TrimSpace(c.FormValue("national_id"))
	req.Password = c.FormValue("password")

	// ۵) شروع تراکنش
	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "DB connection error"})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ۶) چک وجود کاربر هدف (فقط حذف‌نشده‌ها)
	var user Models.User
	if err := tx.Where("deleted_at IS NULL").First(&user, targetID).Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	// ۷) اگر فیلدی خالی بود، از مدل فعلی بخوان
	if req.FirstName == "" {
		req.FirstName = user.FirstName
	}
	if req.LastName == "" {
		req.LastName = user.LastName
	}
	if req.Phone == "" {
		req.Phone = user.Phone
	}
	if req.Address == "" {
		req.Address = user.Address
	}
	if req.NationalID == "" {
		req.NationalID = user.NationalID
	}

	// ۸) هش پسورد در صورت ارسال
	if req.Password != "" {
		hash, err := Utils.GenerateHashPassword(req.Password)
		if err != nil {
			tx.Rollback()
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
		}
		req.Password = hash
	}

	// ۹) ساخت DTO
	dto := User.UserUpdateDTO{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Address:    req.Address,
		NationalID: req.NationalID,
		Password:   req.Password,
	}

	// ۱۰) به‌روزرسانی
	if err := User.Update(tx, targetID, dto); err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// ۱۱) کامیت
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Commit failed"})
	}

	// ۱۲) بارگذاری مجدد برای داشتن داده‌های تازه
	var freshUser Models.User
	if err := Config.DB.Where("deleted_at IS NULL").First(&freshUser, targetID).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to reload user"})
	}

	// ۱۳) پاسخ
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    UserResource.Single(freshUser),
	})
}
