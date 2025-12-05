// file: Resources/OrderResource/OrderResource.go
package OrderResource

import (
	"time"

	"github.com/vahidlotfi71/online-store-api/Models"
)

type OrderItemDTO struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderDTO struct {
	ID         uint           `json:"id"`
	UserID     uint           `json:"user_id"`
	Status     string         `json:"status"`
	TotalPrice float64        `json:"total_price"`
	Items      []OrderItemDTO `json:"items"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func Single(o Models.Order) OrderDTO {
	items := make([]OrderItemDTO, len(o.Items))
	for i, item := range o.Items {
		items[i] = OrderItemDTO{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}
	return OrderDTO{
		ID:         o.ID,
		UserID:     o.UserID,
		Status:     string(o.Status),
		TotalPrice: o.TotalPrice,
		Items:      items,
		CreatedAt:  o.CreatedAt,
		UpdatedAt:  o.UpdatedAt,
	}
}

func Collection(orders []Models.Order) []OrderDTO {
	out := make([]OrderDTO, 0, len(orders))
	for _, o := range orders {
		out = append(out, Single(o))
	}
	return out
}
