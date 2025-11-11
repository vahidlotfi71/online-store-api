package Utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vahidlotfi71/online-store-api.git/Config"
)

// JWTClaims عمومی برای User و Admin
type JWTClaims struct {
	ID    uint   `json:"id"`   // ID کاربر یا ادمین
	Role  string `json:"role"` // "user" یا "admin"
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

// ایجاد توکن عمومی برای User یا Admin
func CreateToken(id uint, role, name, phone string, rememberMe bool) (string, time.Time, error) {
	months := 1
	if rememberMe {
		months = 6
	}

	expireTime := time.Now().AddDate(0, months, 0)

	claims := JWTClaims{
		ID:    id,
		Role:  role,
		Name:  name,
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "online-store-api",
			Subject:   strconv.FormatUint(uint64(id), 10),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Config.JWT_SECRET))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expireTime, nil
}

// بررسی و استخراج اطلاعات توکن
func VerifyToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Config.JWT_SECRET), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// بررسی زمان انقضا
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("expired token")
	}

	return claims, nil
}
