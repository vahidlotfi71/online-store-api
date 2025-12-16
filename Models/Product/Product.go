package Product

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Utils/Http"
	"gorm.io/gorm"
)

type ProductCreateDTO struct {
	Name        string
	Brand       string
	Price       float64
	Description string
	Stock       int
	IsActive    bool
}

type ProductUpdateDTO struct {
	Name        string
	Brand       string
	Price       float64
	Description string
	Stock       int
	IsActive    bool
}

type PaginationMetadata struct {
	TotalRecords int
	TotalPages   int
	CurrentPage  int
	PerPage      int
}

func Paginate(tx *gorm.DB, c *fiber.Ctx) (products []Models.Product, meta Http.PaginationMetadata, err error) {
	tx, meta = Http.Paginate(tx, c)
	err = tx.Find(&products).Error
	return
}

func FindByID(tx *gorm.DB, id uint) (product Models.Product, err error) {
	err = tx.Where("deleted_at IS NULL").First(&product, id).Error
	return
}

func Create(tx *gorm.DB, dto ProductCreateDTO) (product Models.Product, err error) {
	product = Models.Product{
		Name:        dto.Name,
		Brand:       dto.Brand,
		Price:       dto.Price,
		Description: dto.Description,
		Stock:       dto.Stock,
		IsActive:    dto.IsActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = tx.Create(&product).Error
	return
}

func Update(tx *gorm.DB, id uint, dto ProductUpdateDTO) error {
	updates := map[string]interface{}{
		"name":        dto.Name,
		"brand":       dto.Brand,
		"price":       dto.Price,
		"description": dto.Description,
		"stock":       dto.Stock,
		"is_active":   dto.IsActive,
		"updated_at":  time.Now(),
	}

	result := tx.Model(&Models.Product{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found or already deleted")
	}
	return nil
}

func SoftDelete(tx *gorm.DB, id uint) error {
	result := tx.Model(&Models.Product{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", time.Now())

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found or already deleted")
	}
	return nil
}

func DecreaseStock(tx *gorm.DB, productID uint, quantity int) error {
	var product Models.Product
	if err := tx.First(&product, productID).Error; err != nil {
		return err
	}

	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	product.Stock -= quantity
	return tx.Save(&product).Error
}
