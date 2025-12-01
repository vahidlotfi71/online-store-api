// file: Resources/UserResource/user_resource.go
package UserResource

import (
	"time"

	"github.com/vahidlotfi71/online-store-api/Models"
)

type User struct {
	ID         uint      `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	NationalID string    `json:"national_ID"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func Single(u Models.User) User {
	return User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Phone:      u.Phone,
		Address:    u.Address,
		NationalID: u.NationalID,
		IsVerified: u.IsVerified,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func Collection(users []Models.User) []User {
	out := make([]User, 0, len(users))
	for _, u := range users {
		out = append(out, Single(u))
	}
	return out
}
