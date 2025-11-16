package Admin

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models"
	"github.com/vahidlotfi71/online-store-api.git/internal/Utils/Http"
	"gorm.io/gorm"
)

var fillable = []string{
	"first_name", "last_name", "phone", "address", "national_id", "password",
}

type AdminCreateDTO struct {
	FirstName  string
	LastName   string
	Phone      string
	Address    string
	NationalID string
	Password   string
}

type AdminUpdateDTO struct {
	FirstName  string
	LastName   string
	Phone      string
	Address    string
	NationalID string
	Password   string // اگر خالی باشد آپدیت نمی‌شود
}

func Paginate(tx *gorm.DB, c *fiber.Ctx) (admins []Models.Admin, meta Http.PaginationMetadata, err error) {
	tx, meta = Http.Paginate(tx, c)
	err = tx.Find(&admins).Error
	return
}

func FindByID(tx *gorm.DB, id uint) (admin Models.Admin, err error) {
	err = tx.Where("deleted_at IS NULL").First(&admin, id).Error
	return
}

func Create(tx *gorm.DB, dto AdminCreateDTO) (admin Models.Admin, err error) {
	admin = Models.Admin{
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		Phone:      dto.Phone,
		Address:    dto.Address,
		NationalID: dto.NationalID,
		Password:   dto.Password,
		Role:       "admin",
		IsVerified: false,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}
	err = tx.Create(&admin).Error
	return
}

func Update(tx *gorm.DB, id uint, dto AdminUpdateDTO) error {
	updates := map[string]interface{}{
		"first_name":  dto.FirstName,
		"last_name":   dto.LastName,
		"phone":       dto.Phone,
		"address":     dto.Address,
		"national_id": dto.NationalID,
		"updated_at":  time.Now(),
	}
	if dto.Password != "" {
		updates["password"] = dto.Password
	}

	result := tx.Model(&Models.Admin{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("admin not found or already deleted")
	}
	return nil
}

func SoftDelete(tx *gorm.DB, id uint) error {
	result := tx.Model(&Models.Admin{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", time.Now())

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("admin not found or already deleted")
	}
	return nil
}
