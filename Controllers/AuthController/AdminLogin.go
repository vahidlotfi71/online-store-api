package AuthController

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models"
	"github.com/vahidlotfi71/online-store-api.git/Utils"
)

type adminLoginResp struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}

type adminLoginReq struct {
	Email      string `json:"email"` // or Phone
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

func AdminLogin(c *fiber.Ctx) error {
	var body adminLoginReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	db := Config.DB
	var admin Models.Admin

	switch {
	case body.Email != "":
		db.Where("deleted_at IS NULL").Where("email = ?", body.Email).First(&admin)
	case body.Phone != "":
		db.Where("deleted_at IS NULL").Where("phone = ?", body.Phone).First(&admin)
	default:
		return c.Status(422).JSON(fiber.Map{"message": "Email or phone is required"})
	}

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
