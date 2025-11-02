package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/internal/models"
	"github.com/vahidlotfi71/online-store-api.git/internal/services"
	"github.com/vahidlotfi71/online-store-api.git/internal/utils"
)

// created a struct that only has a product service inside it.
type ProductController struct {
	ProductService *services.ProductService
}

// wrote a constructor.
func NewProductController(ps *services.ProductService) *ProductController {
	return &ProductController{ProductService: ps}
}

func (pc *ProductController) CreateProduct(c *fiber.Ctx) error {
	var input struct {
		Name        string  `json:"name"`
		Brand       string  `json:"brand"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
		Stock       int     `json:"stock"`
	}
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "The data format is incorrect.")
	}

	product := &models.Product{
		Name:        input.Name,
		Brand:       input.Brand,
		Price:       input.Price,
		Description: input.Description,
		Stock:       input.Stock,
		IsActive:    true,
	}

	if err := pc.ProductService.CreateProduct(product); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error saving data")
	}

	return utils.SuccessResponse(c, product)
}

func (pc *ProductController) GetProducts(c *fiber.Ctx) error {
	page, limit := 1, 10
	if p := c.Query("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if l := c.Query("limit"); l != "" {
		limit, _ = strconv.Atoi(l)
	}

	filter := map[string]interface{}{}
	if name := c.Query("name"); name != "" {
		filter["name"] = name
	}
	if brand := c.Query("brand"); brand != "" {
		filter["brand"] = brand
	}

	list, total, err := pc.ProductService.GetProducts(filter, page, limit)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error fetching data")
	}

	return c.JSON(fiber.Map{"success": true, "message": "List of products", "data": fiber.Map{"products": list, "pagination": fiber.Map{"page": page, "limit": limit, "total": total}}})
}

func (pc *ProductController) GetProductByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	product, err := pc.ProductService.GetProductByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Product not found")
	}

	return utils.SuccessResponse(c, product)
}

func (pc *ProductController) UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	product, err := pc.ProductService.GetProductByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Product not found")
	}

	// So instead of always having to delete the product, we can control its status with this IsActive field.
	var input struct {
		Name        string  `json:"name"`
		Brand       string  `json:"brand"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
		Stock       int     `json:"stock"`
		IsActive    bool    `json:"is_active"`
	}
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "The data format is incorrect.")
	}

	product.Name = input.Name
	product.Brand = input.Brand
	product.Price = input.Price
	product.Description = input.Description
	product.Stock = input.Stock
	product.IsActive = input.IsActive

	if err := pc.ProductService.UpdateProduct(product); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error updating data")
	}

	return c.JSON(fiber.Map{"success": true, "message": "Product updated successfully", "data": product})
}

func (pc *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	if err := pc.ProductService.DeleteProduct(uint(id)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error deleting product")
	}

	return c.JSON(fiber.Map{"success": true, "message": "Product deleted successfully"})
}
