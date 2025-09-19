package models

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Brand       string `gorm:"not null" json:"brand"`
	Price       int64  `gorm:"not null" json:"price"` //تومان
	Description string `json:"description"`
	Stock       int    `gorm:"not null;default:0" json:"stock"` //موجودی
	IsActive    bool   `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func (Product) TableName() string { return "products" }
