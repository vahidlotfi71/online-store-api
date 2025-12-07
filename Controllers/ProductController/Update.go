package ProductController

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Models/Product"
	"github.com/vahidlotfi71/online-store-api/Resources/ProductResource"
	"gorm.io/gorm"
)

func Update(c *fiber.Ctx) error {
	// ۱) دریافت شناسه محصول
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "id param is required"})
	}

	num, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID format"})
	}
	id := uint(num)

	// ۲) شروع تراکنش
	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "DB connection error"})
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ۳) یافتن محصول فعلی
	var product Models.Product
	if err := tx.Where("deleted_at IS NULL").First(&product, id).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Product not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// ۴) خواندن داده‌ها
	var dto Product.ProductUpdateDTO

	// Name
	if nameVal := strings.TrimSpace(c.FormValue("name")); nameVal != "" {
		dto.Name = nameVal
	} else {
		dto.Name = product.Name
	}

	// Brand
	if brandVal := strings.TrimSpace(c.FormValue("brand")); brandVal != "" {
		dto.Brand = brandVal
	} else {
		dto.Brand = product.Brand
	}

	// Price
	priceVal := strings.TrimSpace(c.FormValue("price"))
	if priceVal != "" {
		if price, err := strconv.ParseFloat(priceVal, 64); err == nil {
			dto.Price = price
		} else {
			dto.Price = product.Price
		}
	} else {
		dto.Price = product.Price
	}

	// Description
	if descVal := strings.TrimSpace(c.FormValue("description")); descVal != "" {
		dto.Description = descVal
	} else {
		dto.Description = product.Description
	}

	// Stock
	stockVal := strings.TrimSpace(c.FormValue("stock"))
	if stockVal != "" {
		if stock, err := strconv.Atoi(stockVal); err == nil {
			dto.Stock = stock
		} else {
			dto.Stock = product.Stock
		}
	} else {
		dto.Stock = product.Stock
	}

	// IsActive - فقط true/false
	isActiveVal := strings.TrimSpace(c.FormValue("is_active"))
	if isActiveVal != "" {
		// از آنجایی که validation قبل از این مرحله انجام شده
		// فقط true یا false وارد می‌شود
		dto.IsActive = (strings.ToLower(isActiveVal) == "true")
	} else {
		// اگر خالی بود، مقدار قبلی را نگه دار
		dto.IsActive = product.IsActive
	}

	// ۵) فراخوانی تابع Update
	if err := Product.Update(tx, id, dto); err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// ۶) کامیت تراکنش
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Commit failed"})
	}

	// ۷) بارگذاری مجدد
	var freshProduct Models.Product
	if err := Config.DB.Where("deleted_at IS NULL").First(&freshProduct, id).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to reload product"})
	}

	// ۸) پاسخ
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"data":    ProductResource.Single(freshProduct),
	})
}
