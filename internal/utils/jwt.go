package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vahidlotfi71/online-store-api.git/config"
)

func GenerateJWT(userID uint, phone, role string, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"phone":   phone,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //یک توکن جدید با الگوریتم HS256 ساخته می‌شه
	return token.SignedString([]byte(cfg.JWT.Secret))          //توکن ساخته شده با کلید مخفی (JWTSecret) امضا می‌شه.
}
