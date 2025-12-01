package Models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	FirstName  string         `json:"first_name" gorm:"not null"`
	LastName   string         `json:"last_name" gorm:"not null"`
	Phone      string         `json:"phone" gorm:"unique;not null"`
	Address    string         `json:"address" gorm:"not null"`
	NationalID string         `json:"national_ID" gorm:"unique;not null"`
	Password   string         `json:"-" gorm:"not null"`
	Role       string         `json:"-" gorm:"not null;default:user"`
	IsVerified bool           `json:"is_verified" gorm:"default:false"`
	VerifyCode string         `json:"-" gorm:"size:6"`
	CreatedAt  time.Time      `json:"created_at"`     // ✅ استاندارد GORM
	UpdatedAt  time.Time      `json:"updated_at"`     // ✅ استاندارد GORM
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"` // ✅ استاندارد GORM برای soft delete
}

func (User) TableName() string {
	return "users"
}
