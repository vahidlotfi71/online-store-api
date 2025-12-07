package Models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	FirstName  string         `json:"first_name" gorm:"not null"`
	LastName   string         `json:"last_name" gorm:"not null"`
	Phone      string         `json:"phone" gorm:"unique;not null"`
	Address    string         `json:"address" gorm:"not null"`
	NationalID string         `json:"national_id" gorm:"unique;not null"`
	Password   string         `json:"-" gorm:"not null"`
	Role       string         `json:"-" gorm:"not null;default:admin"`
	IsVerified bool           `json:"is_verified" gorm:"default:false"`
	VerifyCode string         `json:"-" gorm:"size:6"`
	CreatedAt  time.Time      `json:"created_at"` // ✅ تغییر از CreateAt
	UpdatedAt  time.Time      `json:"updated_at"` // ✅ تغییر از UpdateAt
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Admin) TableName() string { return "admin" }
