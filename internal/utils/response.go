package utils

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	//c *fiber.Ctx → کانتکست درخواست است، یعنی همه اطلاعات مربوط به درخواست و پاسخ تو این شی هست.
	//data interface{} → داده‌ای که می‌خوایم برگردونیم (هر نوعی می‌تونه باشه).
	return c.JSON(fiber.Map{"success": true, "data": data})
}

func ErrorResponse(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(fiber.Map{"success": false, "message": msg})
}
