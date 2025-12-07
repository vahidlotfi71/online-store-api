package ProductResource

import (
	"time"

	"github.com/vahidlotfi71/online-store-api/Models"
	"gorm.io/gorm"
)

type ProductDTO struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Brand       string         `json:"brand"`
	Price       float64        `json:"price"`
	Description string         `json:"description"`
	Stock       int            `json:"stock"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"` // ✅ تغییر
	UpdatedAt   time.Time      `json:"updated_at"` // ✅ تغییر
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func Single(p Models.Product) ProductDTO {
	return ProductDTO{
		ID:          p.ID,
		Name:        p.Name,
		Brand:       p.Brand,
		Price:       p.Price,
		Description: p.Description,
		Stock:       p.Stock,
		IsActive:    p.IsActive,
		CreatedAt:   p.CreatedAt, // ✅ تغییر از CreateAt
		UpdatedAt:   p.UpdatedAt, // ✅ تغییر از UpdateAt
		DeletedAt:   p.DeletedAt,
	}
}

func Collection(products []Models.Product) []ProductDTO {
	out := make([]ProductDTO, 0, len(products))
	for _, p := range products {
		out = append(out, Single(p))
	}
	return out
}
