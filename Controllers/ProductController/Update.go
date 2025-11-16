package ProductController

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Resources/ProductResource"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models/Product"
	"gorm.io/gorm"
)

type ProductUpdateRequest struct {
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	IsActive    bool    `json:"is_active"`
}

func Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "id param is required"})
	}
	num, _ := strconv.ParseUint(idStr, 10, 32)
	id := uint(num)

	var req ProductUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "DB connection error"})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var product Models.Product
	if err := tx.Where("deleted_at IS NULL").First(&product, id).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Product not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	dto := Product.ProductUpdateDTO{
		Name:        req.Name,
		Brand:       req.Brand,
		Price:       req.Price,
		Description: req.Description,
		Stock:       req.Stock,
		IsActive:    req.IsActive,
	}

	if err := Product.Update(tx, id, dto); err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Commit failed"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"data":    ProductResource.Single(product),
	})
}
