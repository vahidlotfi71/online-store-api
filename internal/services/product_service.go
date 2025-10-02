package services

import (
	"github.com/vahidlotfi71/online-store-api.git/internal/models"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

func (s *ProductService) CreateProduct(p *models.Product) error {
	return s.DB.Create(p).Error
}

func (s *ProductService) GetProducts(filter map[string]interface{}, page, limit int) ([]*models.Product, int64, error) {
	//filter → یه map از فیلترها (مثلاً {"name": "iPhone", "brand": "Apple"}).
	//page, limit → برای pagination (صفحه‌بندی)
	var list []*models.Product
	var total int64
	q := s.DB.Model(&models.Product{}).Where("is_active = ?", true)
	if v, ok := filter["name"]; ok {
		q = q.Where("name LIKE ?", "%"+v.(string)+"%")
	}
	if v, ok := filter["brand"]; ok {
		q = q.Where("brand LIKE ?", "%"+v.(string)+"%")
	}
	q.Count(&total)
	offset := (page - 1) * limit
	err := q.Limit(limit).Offset(offset).Find(&list).Error
	// Limit(limit): محدود کردن تعداد رکوردهای بازگشتی
	// Offset(offset): مشخص کردن نقطه شروع برای دریافت رکوردها
	// Find(&list): اجرای کوئری و ذخیره نتایج در اسلایس list
	// .Error: دریافت خطا در صورت وجود
	// page = 2
	// limit = 10
	// offset = (2 - 1) * 10 = 10
	// 	SELECT * FROM products
	// WHERE is_active = true
	// LIMIT 10 OFFSET 10;
	return list, total, err
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	var p models.Product
	err := s.DB.First(&p, id).Error
	return &p, err
}

func (s *ProductService) UpdateProduct(p *models.Product) error {
	return s.DB.Save(p).Error
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.DB.Delete(&models.Product{}, id).Error
}
