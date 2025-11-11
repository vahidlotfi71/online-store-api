package Utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/internal/Models"
)

type JWTClaims struct {
	UserID          uint      `json:"user_id"`
	Phone           string    `json:"phone"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	IsVerified      bool      `json:"is_verified"`
	Expiration_date time.Time `json:"expiration_date"`
	jwt.RegisteredClaims
}

type User = Models.User

// CreateToken creates a new JWT token for a user
func CreateToken(user User, remember_me bool) (string, time.Time, error) {
	months := 1
	if remember_me {
		months = 6
	}

	expire_time := time.Now().AddDate(0, months, 0)

	claims := JWTClaims{
		UserID:          user.ID,
		Phone:           user.Phone,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		IsVerified:      user.IsVerified,
		Expiration_date: expire_time,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire_time),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app-name",
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Config.JWT_SECRET))

	if err != nil {
		return "", time.Time{}, err
	}
	return string(tokenString), expire_time, nil
}

// VerifyToken verifies and parses a JWT token
func VerifyToken(tokenString string) (int64, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Config.JWT_SECRET), nil
	})

	if err != nil {
		return 0, "", fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}

	// Check expiration using RegisteredClaims (standard way)
	if claims.Expiration_date.Unix() < time.Now().Unix() {
		return 0, "", fmt.Errorf("expired token")
	}

	return int64(claims.UserID), claims.Phone, nil

}
