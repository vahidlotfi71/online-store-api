package AdminController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models/Admin"
	"github.com/vahidlotfi71/online-store-api/Resources/AdminResource"
	"gorm.io/gorm"
)

func Show(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(400).JSON(fiber.Map{"message": "id param is required"})
	}
	num, _ := strconv.ParseUint(idStr, 10, 32)
	id := uint(num)

	admin, err := Admin.FindByID(Config.DB, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"message": "Admin not found"})
		}
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": AdminResource.Single(admin)})
}
