package AuthController

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

type adminLoginResp struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}

type adminLoginReq struct {
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

func AdminLogin(c *fiber.Ctx) error {
	var body adminLoginReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	if body.Phone == "" {
		return c.Status(422).JSON(fiber.Map{"message": "Phone is required"})
	}

	db := Config.DB
	var admin Models.Admin

	// فقط بر اساس شماره تلفن جستجو می‌کنیم
	db.Where("deleted_at IS NULL").Where("phone = ?", body.Phone).First(&admin)

	if admin.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Admin not found"})
	}

	if err := Utils.VerifyPassword(body.Password, admin.Password); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Wrong credentials"})
	}

	token, expire, err := Utils.CreateToken(
		admin.ID,
		admin.Role,
		admin.FirstName+" "+admin.LastName,
		admin.Phone,
		body.RememberMe,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to generate token"})
	}

	return c.Status(200).JSON(adminLoginResp{
		Token:      token,
		ExpireTime: expire,
	})
}
