// file: internal/Models/User/user_repository.go
package User

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Models"
	"github.com/vahidlotfi71/online-store-api.git/Utils/Http"
	"gorm.io/gorm"
)

// فیلدهای قابل نوشتن از بیرون
var fillable = []string{
	"first_name", "last_name", "phone", "address", "national_id", "password",
}

/* ---------- DTOها (type-safe) ---------- */
type UserCreateDTO struct {
	FirstName  string
	LastName   string
	Phone      string
	Address    string
	NationalID string
	Password   string
}

type UserUpdateDTO struct {
	FirstName  string
	LastName   string
	Phone      string
	Address    string
	NationalID string
	Password   string // اگر خالی باشد آپدیت نمی‌شود
}

type PaginationMetadata struct {
	TotalRecords int
	TotalPages   int
	CurrentPage  int
	PerPage      int
}

/* ---------- صفحه‌بندی (با خواندن page & per_page از کوئری) ---------- */
func Paginate(tx *gorm.DB, c *fiber.Ctx) (users []Models.User, meta Http.PaginationMetadata, err error) {
	tx, meta = Http.Paginate(tx, c)
	err = tx.Find(&users).Error
	return
}

/* ---------- خواندن یک رکورد (بدون RAW SQL) ---------- */
func FindByID(tx *gorm.DB, id uint) (user Models.User, err error) {
	err = tx.Where("deleted_at IS NULL").First(&user, id).Error
	return
}

/* ---------- درج (با DTO و بدون RAW SQL) ---------- */
func Create(tx *gorm.DB, dto UserCreateDTO) (user Models.User, err error) {
	user = Models.User{
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		Phone:      dto.Phone,
		Address:    dto.Address,
		NationalID: dto.NationalID,
		Password:   dto.Password,
		Role:       "user",
		IsVerified: false,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}
	err = tx.Create(&user).Error
	return
}

/* ---------- به‌روزرسانی (با DTO و بدون RAW SQL) ---------- */
func Update(tx *gorm.DB, id uint, dto UserUpdateDTO) error {
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

	result := tx.Model(&Models.User{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}
	return nil
}

/* ---------- حذف منطقی (بدون RAW SQL) ---------- */
func SoftDelete(tx *gorm.DB, id uint) error {
	result := tx.Model(&Models.User{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", time.Now())

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}
	return nil
}
