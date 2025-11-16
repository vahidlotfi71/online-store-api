package AuthController

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models"
	"github.com/vahidlotfi71/online-store-api.git/internal/Utils"
)

type RegisterResponse struct {
	Token      string      `json:"token"`
	ExpireTime time.Time   `json:"expire_time"`
	User       Models.User `json:"user"`
}

// Register ثبت‌نام کاربر جدید
func Register(c *fiber.Ctx) error {
	// شروع یک تراکنش دیتابیس؛ در صورت خطا ۵۰۰ برمی‌گردد
	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Database connection error"})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 2) تعریف ساختار ورودی و پارس کردن بادی درخواست
	var input struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Phone      string `json:"phone"`
		Address    string `json:"address"`
		NationalID string `json:"national_ID"`
		Password   string `json:"password"`
		RememberMe bool   `json:"remember_me"`
	}
	if err := c.BodyParser(&input); err != nil {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input data"})
	}

	// هش کردن رمز عبور
	hashedPass, err := Utils.GenerateHashPassword(input.Password)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"message": "خطا در رمزنگاری رمز عبور"})
	}

	//  ساخت مدل کاربر و درج در جدول users
	now := time.Now()
	user := Models.User{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Phone:      input.Phone,
		Address:    input.Address,
		NationalID: input.NationalID,
		Password:   hashedPass,
		Role:       "user",
		IsVerified: false,
		CreateAt:   now,
		UpdateAt:   now,
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"message": "خطا در ثبت کاربر"})
	}

	// تولید JWT با اطلاعات کاربر
	token, expireTime, err := Utils.CreateToken(
		user.ID,
		user.Role,
		user.FirstName+" "+user.LastName,
		user.Phone,
		input.RememberMe,
	)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"message": "خطا در ایجاد توکن"})
	}

	//ثبت نهایی تراکنش
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"message": "خطا در ذخیره تراکنش"})
	}

	// ارسال پاسخ موفقیت
	return c.Status(200).JSON(RegisterResponse{
		Token:      token,
		ExpireTime: expireTime,
		User:       user,
	})
}
