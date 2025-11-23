package AuthController

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

/* ----------  ساختار پاسخ  ---------- */
type loginResp struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}

/* ----------  ساختار ورودی  ---------- */
type loginReq struct {
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

/* ----------  تابع لاگین  ---------- */
func Login(c *fiber.Ctx) error {
	var body loginReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	if body.Phone == "" {
		return c.Status(422).JSON(fiber.Map{"message": "Phone is required"})
	}

	db := Config.DB
	var user Models.User

	// فقط بر اساس شماره تلفن جستجو می‌کنیم
	db.Where("deleted_at IS NULL").Where("phone = ?", body.Phone).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	if err := Utils.VerifyPassword(body.Password, user.Password); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Wrong credentials"})
	}

	/* ----  تولید JWT ---- */
	token, expire, err := Utils.CreateToken(
		user.ID,
		user.Role,
		user.FirstName+" "+user.LastName,
		user.Phone,
		body.RememberMe,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to generate token"})
	}

	/* ----  پاسخ موفق  ---- */
	return c.Status(200).JSON(loginResp{
		Token:      token,
		ExpireTime: expire,
	})
}
