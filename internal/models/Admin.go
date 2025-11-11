package Models

import "time"

type Admin struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	FirstName  string     `json:"first_name" gorm:"not null"`
	LastName   string     `json:"last_name" gorm:"not null"`
	Phone      string     `json:"phone" gorm:"unique;not null"`
	Address    string     `json:"address" gorm:"not null"`
	NationalID string     `json:"national_ID" gorm:"unique;not null"`
	Password   string     `json:"-" gorm:"not null"`
	Role       string     `json:"-" gorm:"not null;default:admin"` // مقدار پیش‌فرض admin
	IsVerified bool       `json:"is_verified" gorm:"default:false"`
	VerifyCode string     `json:"-"  gorm:"size:6"`
	CreateAt   time.Time  `json:"created_at"`
	UpdateAt   time.Time  `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}

func (Admin) TableName() string { return "admin" }
