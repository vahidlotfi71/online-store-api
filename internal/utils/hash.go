package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword → پسورد رو هش می‌کنه (برای ذخیره‌سازی امن).
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //پسورد رو به بایت تبدیل می‌کنه (چون bcrypt با بایت کار می‌کنه).
	//bcrypt.DefaultCost یعنی سطح سختی پیش‌فرض الگوریتم
	return string(bytes), err
}

// CheckPassword → بررسی می‌کنه آیا پسورد وارد شده با هش ذخیره‌شده مطابقت داره یا نه.
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
