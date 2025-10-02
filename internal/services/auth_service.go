package services

import (
	"github.com/vahidlotfi71/online-store-api.git/config"
	"github.com/vahidlotfi71/online-store-api.git/internal/models"
	"gorm.io/gorm"
)

type AuthService struct {
	DB  *gorm.DB       //DB → اتصال به دیتابیس (gorm).
	CFG *config.Config //CFG → تنظیمات پروژه.
}

// وقتی می‌خوای از AuthService استفاده کنی، باید اول یه نمونه بسازی.
// این تابع سازنده هست و یه اشاره‌گر (*AuthService) برمی‌گردونه.
func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{DB: db, CFG: cfg}
}

// . متدهای CRUD روی User
func (s *AuthService) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}

func (s *AuthService) GetUserByPhone(phone string) (*models.User, error) {
	var u models.User
	err := s.DB.Where("phone = ?", phone).First(&u).Error
	return &u, err
}

func (s *AuthService) GetUserByID(id uint) (*models.User, error) {
	var u models.User
	err := s.DB.First(&u, id).Error
	return &u, err //خروجی error هست → اگر موفق بشه nil، اگر خطا بشه خطا رو برمی‌گردونه.
}

func (s *AuthService) UpdateUser(user *models.User) error {
	return s.DB.Save(user).Error //gorm.Save → اگر رکورد وجود داشته باشه آپدیت می‌کنه، اگر نه می‌سازه.
}
