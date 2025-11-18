package AuthController

import "github.com/gofiber/fiber/v2"

func AdminLogout(c *fiber.Ctx) error {
	// اگر سشن یا توکن کاربر را ذخیره کرده‌ ایم، آن را حذف میکنیم
	// به عنوان مثال:
	// c.ClearCookie("admin_session")
	// یا اگر از JWT استفاده میکنیم توکن را در سمت کلاینت حذف میکنیم

	return c.Status(200).JSON(fiber.Map{
		"message": "You have been logged out",
	})
}
