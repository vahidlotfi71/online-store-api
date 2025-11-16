package ProductController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Resources/ProductResource"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models/Product"
)

type ProductCreateRequest struct {
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	IsActive    bool    `json:"is_active"`
}

func Store(c *fiber.Ctx) error {
	var req ProductCreateRequest
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

	dto := Product.ProductCreateDTO{
		Name:        req.Name,
		Brand:       req.Brand,
		Price:       req.Price,
		Description: req.Description,
		Stock:       req.Stock,
		IsActive:    req.IsActive,
	}

	product, err := Product.Create(tx, dto)
	if err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Commit failed"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"data":    ProductResource.Single(product),
	})
}
