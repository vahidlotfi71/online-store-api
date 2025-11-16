// file: internal/Controllers/Admin/AdminController/ClearTrash.go
package AdminController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Models"
)

func ClearTrash(c *fiber.Ctx) error {
	limitStr := c.Query("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	result := Config.DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		Limit(limit).
		Delete(&Models.Admin{})

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":       "Trash cleared successfully",
		"cleared_count": result.RowsAffected,
	})
}
