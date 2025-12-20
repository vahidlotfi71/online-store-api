package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Models/Product"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupProductTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&Models.Product{})
	assert.NoError(t, err)

	return db
}

func TestProductCreate_Success(t *testing.T) {
	db := setupProductTestDB(t)

	dto := Product.ProductCreateDTO{
		Name:        "Test Product",
		Brand:       "Test Brand",
		Price:       99.99,
		Description: "Test Description",
		Stock:       10,
		IsActive:    true,
	}

	product, err := Product.Create(db, dto)

	assert.NoError(t, err)
	assert.NotZero(t, product.ID)
	assert.Equal(t, "Test Product", product.Name)
	assert.Equal(t, "Test Brand", product.Brand)
	assert.Equal(t, 99.99, product.Price)
	assert.Equal(t, 10, product.Stock)
	assert.True(t, product.IsActive)
}

func TestFindByID_Success(t *testing.T) {
	db := setupProductTestDB(t)

	dto := Product.ProductCreateDTO{
		Name:        "Find Me",
		Brand:       "Brand X",
		Price:       49.99,
		Description: "Description",
		Stock:       5,
		IsActive:    true,
	}

	created, _ := Product.Create(db, dto)

	found, err := Product.FindByID(db, created.ID)

	assert.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
	assert.Equal(t, "Find Me", found.Name)
}

func TestFindByID_NotFound(t *testing.T) {
	db := setupProductTestDB(t)

	_, err := Product.FindByID(db, 9999)

	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUpdate_Success(t *testing.T) {
	db := setupProductTestDB(t)

	// Create product
	createDTO := Product.ProductCreateDTO{
		Name:        "Old Name",
		Brand:       "Old Brand",
		Price:       10.00,
		Description: "Old Description",
		Stock:       1,
		IsActive:    false,
	}

	product, _ := Product.Create(db, createDTO)

	// Update product
	updateDTO := Product.ProductUpdateDTO{
		Name:        "New Name",
		Brand:       "New Brand",
		Price:       20.00,
		Description: "New Description",
		Stock:       2,
		IsActive:    true,
	}

	err := Product.Update(db, product.ID, updateDTO)

	assert.NoError(t, err)

	// Verify
	updated, _ := Product.FindByID(db, product.ID)
	assert.Equal(t, "New Name", updated.Name)
	assert.Equal(t, "New Brand", updated.Brand)
	assert.Equal(t, 20.00, updated.Price)
	assert.Equal(t, 2, updated.Stock)
	assert.True(t, updated.IsActive)
}

func TestUpdate_NotFound(t *testing.T) {
	db := setupProductTestDB(t)

	updateDTO := Product.ProductUpdateDTO{
		Name: "Test",
	}

	err := Product.Update(db, 9999, updateDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestSoftDelete_Success(t *testing.T) {
	db := setupProductTestDB(t)

	dto := Product.ProductCreateDTO{
		Name:        "Delete Me",
		Brand:       "Brand",
		Price:       10.00,
		Description: "Description",
		Stock:       5,
		IsActive:    true,
	}

	product, _ := Product.Create(db, dto)

	err := Product.SoftDelete(db, product.ID)

	assert.NoError(t, err)

	// Should not find with default scope
	_, err = Product.FindByID(db, product.ID)
	assert.Error(t, err)

	// Verify deletion with Unscoped
	var deleted Models.Product
	err = db.Unscoped().First(&deleted, product.ID).Error
	assert.NoError(t, err)
	assert.False(t, deleted.DeletedAt.Time.IsZero())
}

func TestSoftDelete_NotFound(t *testing.T) {
	db := setupProductTestDB(t)

	err := Product.SoftDelete(db, 9999)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestDecreaseStock_Success(t *testing.T) {
	db := setupProductTestDB(t)

	dto := Product.ProductCreateDTO{
		Name:        "Stock Test",
		Brand:       "Brand",
		Price:       10.00,
		Description: "Description",
		Stock:       10,
		IsActive:    true,
	}

	product, _ := Product.Create(db, dto)

	err := Product.DecreaseStock(db, product.ID, 3)

	assert.NoError(t, err)

	// Verify stock decreased
	updated, _ := Product.FindByID(db, product.ID)
	assert.Equal(t, 7, updated.Stock)
}

func TestDecreaseStock_InsufficientStock(t *testing.T) {
	db := setupProductTestDB(t)

	dto := Product.ProductCreateDTO{
		Name:        "Low Stock",
		Brand:       "Brand",
		Price:       10.00,
		Description: "Description",
		Stock:       5,
		IsActive:    true,
	}

	product, _ := Product.Create(db, dto)

	err := Product.DecreaseStock(db, product.ID, 10)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "insufficient stock")
}

func TestDecreaseStock_ProductNotFound(t *testing.T) {
	db := setupProductTestDB(t)

	err := Product.DecreaseStock(db, 9999, 1)

	assert.Error(t, err)
}

func TestProductModel_TableName(t *testing.T) {
	product := Models.Product{}
	assert.Equal(t, "products", product.TableName())
}

func TestProductCreateDTO_Structure(t *testing.T) {
	dto := Product.ProductCreateDTO{
		Name:        "Test",
		Brand:       "Brand",
		Price:       99.99,
		Description: "Desc",
		Stock:       10,
		IsActive:    true,
	}

	assert.Equal(t, "Test", dto.Name)
	assert.Equal(t, "Brand", dto.Brand)
	assert.Equal(t, 99.99, dto.Price)
	assert.Equal(t, 10, dto.Stock)
	assert.True(t, dto.IsActive)
}

func TestProductUpdateDTO_Structure(t *testing.T) {
	dto := Product.ProductUpdateDTO{
		Name:        "Updated",
		Brand:       "New Brand",
		Price:       149.99,
		Description: "New Desc",
		Stock:       20,
		IsActive:    false,
	}

	assert.Equal(t, "Updated", dto.Name)
	assert.Equal(t, "New Brand", dto.Brand)
	assert.Equal(t, 149.99, dto.Price)
	assert.Equal(t, 20, dto.Stock)
	assert.False(t, dto.IsActive)
}

func TestCreate_DefaultValues(t *testing.T) {
	db := setupProductTestDB(t)

	dto := Product.ProductCreateDTO{
		Name:        "Defaults Test",
		Brand:       "Brand",
		Price:       10.00,
		Description: "Description",
		Stock:       0,
		IsActive:    false,
	}

	product, err := Product.Create(db, dto)

	assert.NoError(t, err)
	assert.Equal(t, 0, product.Stock)
	assert.False(t, product.IsActive)
	assert.NotZero(t, product.CreatedAt)
	assert.NotZero(t, product.UpdatedAt)
}

func TestUpdate_PartialUpdate(t *testing.T) {
	db := setupProductTestDB(t)

	// Create
	createDTO := Product.ProductCreateDTO{
		Name:        "Original",
		Brand:       "Original Brand",
		Price:       100.00,
		Description: "Original Description",
		Stock:       10,
		IsActive:    true,
	}

	product, _ := Product.Create(db, createDTO)

	// Update only name and price
	updateDTO := Product.ProductUpdateDTO{
		Name:        "Updated Name",
		Brand:       product.Brand,
		Price:       150.00,
		Description: product.Description,
		Stock:       product.Stock,
		IsActive:    product.IsActive,
	}

	err := Product.Update(db, product.ID, updateDTO)

	assert.NoError(t, err)

	updated, _ := Product.FindByID(db, product.ID)
	assert.Equal(t, "Updated Name", updated.Name)
	assert.Equal(t, 150.00, updated.Price)
	assert.Equal(t, "Original Brand", updated.Brand) // Unchanged
}
