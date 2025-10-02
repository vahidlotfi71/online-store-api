package services

import (
	"fmt"

	"github.com/vahidlotfi71/online-store-api.git/internal/models"
	"gorm.io/gorm"
)

// یعنی این سرویس برای کار با جدول سفارش‌ها از دیتابیس (gorm.DB) استفاده می‌کنه.
type OrderService struct {
	DB *gorm.DB
}

// با این تابع می‌تونیم یه سرویس سفارش جدید بسازیم و بهش دیتابیس بدیم
func NewOrderService(db *gorm.DB) *OrderService { return &OrderService{DB: db} }

// CreateOrder → ایجاد سفارش
func (s *OrderService) CreateOrder(userID uint, items []models.OrderItem) (*models.Order, error) {
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var total float64
	for i, item := range items {
		var product models.Product
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if product.Stock < item.Quantity {
			tx.Rollback()
			return nil, fmt.Errorf("محصول %s موجودی کافی ندارد", product.Name)
		}

		itemTotal := float64(item.Quantity) * (product.Price)
		total += itemTotal
		items[i].Price = product.Price // ذخیره قیمت واحد محصول در زمان سفارش

		// کاهش موجودی
		product.Stock -= item.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// ساخت رکورد سفارش
	// 	یک سفارش جدید ساخته می‌شه برای کاربر.
	// وضعیتش هم اول روی Pending (در انتظار) هست.
	order := &models.Order{
		UserID:     userID,
		Status:     models.StatusPending,
		TotalPrice: total,
	}
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// ثبت آیتم‌های سفارش
	for _, item := range items {
		// هر محصولی که کاربر خریده به جدول آیتم‌های سفارش (OrderItems) اضافه می‌شه.
		item.OrderID = order.ID
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// اگر همه مراحل موفق بودن، تراکنش commit می‌شه و تغییرات توی دیتابیس ذخیره می‌شن.
	tx.Commit()
	return order, nil
}

func (s *OrderService) GetUserOrders(userID uint) ([]*models.Order, error) {
	var list []*models.Order
	err := s.DB.Preload("Items.Product").Where("user_id = ?", userID).Find(&list).Error
	// Preload("Items.Product") باعث می‌شه محصولات داخل سفارش هم همراهش لود بشن (lazy loading نیست، eager loading هست).
	return list, err
}

// سفارش مشخصی رو با شناسه id میاره، همراه با محصولاتش.
func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	var o models.Order
	err := s.DB.Preload("Items.Product").First(&o, id).Error
	return &o, err
}
