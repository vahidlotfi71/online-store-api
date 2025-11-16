// file: Resources/ProductResource/ProductResource.go
package ProductResource

import (
	"time"

	"github.com/vahidlotfi71/online-store-api.git/internal/Models"
)

type ProductDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
		CreatedAt:   p.CreateAt,
		UpdatedAt:   p.UpdateAt,
	}
}

func Collection(products []Models.Product) []ProductDTO {
	out := make([]ProductDTO, 0, len(products))
	for _, p := range products {
		out = append(out, Single(p))
	}
	return out
}
