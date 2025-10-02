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
	TotalPrice float64     `gorm:"not null" json:"total_price"` //تومان
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreateAt   time.Time
	UpdateAt   time.Time
	//Relation
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"` // قیمت واحد در زمان خرید
	//Relations
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (Order) TableName() string     { return "orders" }
func (OrderItem) TableName() string { return "order_items" }
