package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models/User"
	"github.com/vahidlotfi71/online-store-api/Resources/UserResource"
)

func Trash(c *fiber.Ctx) error {
	users, meta, err := User.Paginate(
		Config.DB.Unscoped().Where("deleted_at IS NOT NULL").Order("id"),
		c,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"data":     UserResource.Collection(users),
		"metadata": meta,
	})
}
