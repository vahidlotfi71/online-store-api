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
	NationalID string         `json:"national_id" gorm:"unique;not null"`
	Password   string         `json:"-" gorm:"not null"`
	Role       string         `json:"-" gorm:"not null;default:user"`
	IsVerified bool           `json:"is_verified" gorm:"default:false"`
	VerifyCode string         `json:"-" gorm:"size:6"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
