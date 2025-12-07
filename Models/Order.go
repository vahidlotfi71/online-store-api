package Models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	Status     OrderStatus    `gorm:"default:pending" json:"status"`
	TotalPrice float64        `gorm:"not null" json:"total_price"`
	Items      []OrderItem    `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	//Relation
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type OrderItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `gorm:"not null" json:"order_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"not null" json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	//Relations
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (Order) TableName() string     { return "orders" }
func (OrderItem) TableName() string { return "order_items" }
