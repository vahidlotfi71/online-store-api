package AdminController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models/Admin"
	"github.com/vahidlotfi71/online-store-api.git/Resources/AdminResource"
)

func Trash(c *fiber.Ctx) error {
	admins, meta, err := Admin.Paginate(
		Config.DB.Where("deleted_at IS NOT NULL").Order("deleted_at ASC"),
		c,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"data":     AdminResource.Collection(admins),
		"metadata": meta,
	})
}
