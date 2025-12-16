package tests

import (
	"testing"
	"time"

	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

func init() {
	Config.JWT_SECRET = "test-secret-key"
}

func TestCreateToken(t *testing.T) {
	token, expireTime, err := Utils.CreateToken(1, "user", "Test User", "09123456789", false)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Error("Expected token, got empty string")
	}

	if expireTime.Before(time.Now()) {
		t.Error("Token should not be expired")
	}
}

func TestVerifyToken_Valid(t *testing.T) {
	token, _, _ := Utils.CreateToken(1, "admin", "Admin", "09123456789", false)

	claims, err := Utils.VerifyToken(token)

	if err != nil {
		t.Errorf("Expected valid token, got error: %v", err)
	}

	if claims.ID != 1 {
		t.Errorf("Expected ID=1, got %d", claims.ID)
	}

	if claims.Role != "admin" {
		t.Errorf("Expected Role=admin, got %s", claims.Role)
	}
}

func TestVerifyToken_Invalid(t *testing.T) {
	_, err := Utils.VerifyToken("invalid-token")

	if err == nil {
		t.Error("Expected error for invalid token")
	}
}

func TestRememberMe_LongerExpiry(t *testing.T) {
	_, expire1, _ := Utils.CreateToken(1, "user", "Test", "09123456789", false)
	_, expire2, _ := Utils.CreateToken(1, "user", "Test", "09123456789", true)

	if !expire2.After(expire1) {
		t.Error("RememberMe token should have longer expiry")
	}
}
