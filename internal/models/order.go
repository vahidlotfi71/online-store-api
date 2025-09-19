package models

import "time"

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID         uint        `gorm:"PrimaryKey" json:"id"`
	UserID     uint        `gorm:"not null" json:"user-id"`
	Status     OrderStatus `gorm:"default:pending" json:"status"`
	TotalPrice int64       `gorm:"not null" json:"total_price"`     //تومان
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"` //صریحاً مشخص می‌کند که فیلد OrderID داخل OrderItem کلید خارجی مربوط به این سفارش است.
	CreateAt   time.Time
	UpdateAt   time.Time
	//Relation
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"` //مشخص می‌کند فیلد UserID داخل همین جدول (orders) کلید خارجی است که به جدول کاربران اشاره می‌کند.
}

type OrderItem struct {
	ID        uint  `gorm:"primaryKey" json:"id"`
	OrderID   uint  `gorm:"not null" json:"order_id"`
	ProductID uint  `gorm:"not null" json:"product_id"`
	Quantity  int   `gorm:"not null" json:"quantity"`
	Price     int64 `gorm:"not null" json:"price"` // قیمت واحد در زمان خرید
	//Relations
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (Order) TableName() string     { return "orders" }
func (OrderItem) TableName() string { return "order_items" }
