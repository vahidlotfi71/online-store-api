package ProductController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Resources/ProductResource"
	"github.com/vahidlotfi71/online-store-api/Utils"
	"github.com/vahidlotfi71/online-store-api/Utils/Http"
)

func Index(c *fiber.Ctx) error {
	var products []Models.Product

	// همه محصولات حذف نشده
	tx := Config.DB.Table("products").Where("deleted_at IS NULL")

	isAdmin := Utils.IsAdmin(c)

	// بررسی آیا کاربر ادمین است؟
	if isAdmin {
		// برای ادمین: همه محصولات (فعال و غیرفعال)
		tx = tx.Order("id")
	} else {
		// برای کاربران عادی و مهمان: فقط محصولات فعال
		tx = tx.Where("is_active = ?", true).Order("id")
	}

	var metadata Http.PaginationMetadata
	tx, metadata = Http.Paginate(tx, c)
	tx.Find(&products)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":     ProductResource.Collection(products),
		"metadata": metadata,
	})
}
