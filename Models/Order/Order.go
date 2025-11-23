package Order

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Utils/Http"
	"gorm.io/gorm"
)

func Paginate(tx *gorm.DB, c *fiber.Ctx) (orders []Models.Order, meta Http.PaginationMetadata, err error) {
	tx, meta = Http.Paginate(tx, c)
	err = tx.Preload("Items.Product").Find(&orders).Error
	return
}

func FindByID(tx *gorm.DB, id uint) (order Models.Order, err error) {
	err = tx.Where("deleted_at IS NULL").Preload("Items.Product").First(&order, id).Error
	return
}

func UpdateStatus(tx *gorm.DB, id uint, status Models.OrderStatus) error {
	result := tx.Model(&Models.Order{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found or not active")
	}
	return nil
}

func SoftDelete(tx *gorm.DB, id uint) error {
	result := tx.Model(&Models.Order{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", time.Now())

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found or already deleted")
	}
	return nil
}

func CreateOrder(tx *gorm.DB, userID uint, items []OrderItemCreateDTO) (order Models.Order, err error) {
	total := 0.0
	var orderItems []Models.OrderItem

	for _, item := range items {
		var product Models.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			return order, err
		}

		itemTotal := product.Price * float64(item.Quantity)
		total += itemTotal

		orderItems = append(orderItems, Models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}

	order = Models.Order{
		UserID:     userID,
		Status:     Models.StatusPending,
		TotalPrice: total,
		Items:      orderItems,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}

	err = tx.Create(&order).Error
	return order, err
}

type OrderItemCreateDTO struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
