package models

import "time"

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	FirstName  string `json:"first_name" gorm:"not null"`
	LastName   string `json:"last_name" gorm:"not null"`
	Phone      string `json:"Phone" gorm:"unique;not null"`
	Address    string `json:"address" gorm:"not null"`
	NationalID string `json:"national_ID" gorm:"unique;not null"`
	Password   string `json:"-" gorm:"not null"`
	Role       string `json:"role" gorm:"default:user"` //user or admin
	IsVerified bool   `json:"is_verified" gorm:"default:false"`
	VerifyCode string `json:"-"  gorm:"size:6"`
	CreateAt   time.Time
	UpdateAt   time.Time
}

func (User) TableName() string { return "users" }
