package AdminResource

import (
	"time"

	"github.com/vahidlotfi71/online-store-api/Models"
)

type AdminDTO struct {
	ID         uint      `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	NationalID string    `json:"national_id"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func Single(a Models.Admin) AdminDTO {
	return AdminDTO{
		ID:         a.ID,
		FirstName:  a.FirstName,
		LastName:   a.LastName,
		Phone:      a.Phone,
		Address:    a.Address,
		NationalID: a.NationalID,
		IsVerified: a.IsVerified,
		CreatedAt:  a.CreateAt,
		UpdatedAt:  a.UpdateAt,
	}
}

func Collection(admins []Models.Admin) []AdminDTO {
	out := make([]AdminDTO, 0, len(admins))
	for _, a := range admins {
		out = append(out, Single(a))
	}
	return out
}
