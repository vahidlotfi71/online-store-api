package Models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Brand       string         `gorm:"not null" json:"brand"`
	Price       float64        `gorm:"not null" json:"price"`
	Description string         `json:"description"`
	Stock       int            `gorm:"not null;default:0" json:"stock"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Product) TableName() string { return "products" }
