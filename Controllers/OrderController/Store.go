package OrderController

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Resources/OrderResource"
)

type OrderCreateRequest struct {
	Items []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func Store(c *fiber.Ctx) error {
	//خواندن items از form-data
	itemsJSON := c.FormValue("items")
	if itemsJSON == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "items field is required",
		})
	}

	//  تبدیل JSON string به struct
	var req OrderCreateRequest
	if err := json.Unmarshal([]byte(itemsJSON), &req.Items); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid items format. Expected JSON array",
		})
	}

	// بررسی وجود آیتم‌ها
	if len(req.Items) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Order must have at least one item",
		})
	}

	user := c.Locals("user").(Models.User)

	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "DB connection error",
		})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// محاسبه مجموع و بررسی موجودی
	totalPrice := 0.0
	var orderItems []Models.OrderItem

	for _, item := range req.Items {
		var product Models.Product
		if err := tx.Where("deleted_at IS NULL AND is_active = ?", true).First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
			})
		}

		if item.Quantity <= 0 {
			tx.Rollback()
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Quantity must be greater than zero",
			})
		}

		if product.Stock < item.Quantity {
			tx.Rollback()
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Insufficient stock for product: " + product.Name,
			})
		}

		itemTotal := product.Price * float64(item.Quantity)
		totalPrice += itemTotal

		orderItems = append(orderItems, Models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		// کاهش موجودی
		product.Stock -= item.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update product stock",
			})
		}
	}

	// ایجاد سفارش
	order := Models.Order{
		UserID:     user.ID,
		Status:     Models.StatusPending,
		TotalPrice: totalPrice,
		Items:      orderItems,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create order",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Commit failed",
		})
	}

	// بارگذاری مجدد برای شامل شدن روابط
	var createdOrder Models.Order
	if err := Config.DB.Preload("Items.Product").First(&createdOrder, order.ID).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to load order details",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Order created successfully",
		"data":    OrderResource.Single(createdOrder),
	})
}
