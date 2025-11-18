package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models/User"
	"github.com/vahidlotfi71/online-store-api.git/Resources/UserResource"
)

func Trash(c *fiber.Ctx) error {
	users, meta, err := User.Paginate(
		Config.DB.Where("deleted_at IS NOT NULL").Order("id"),
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
