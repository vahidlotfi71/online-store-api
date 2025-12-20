package Test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

func init() {
	// Set a test JWT secret
	Config.JWT_SECRET = "test-secret-key-for-testing-only"
}

func TestCreateToken_UserRole(t *testing.T) {
	id := uint(1)
	role := "user"
	name := "Test User"
	phone := "09123456789"
	rememberMe := false

	token, expireTime, err := Utils.CreateToken(id, role, name, phone, rememberMe)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.True(t, expireTime.After(time.Now()), "Expire time should be in the future")
}

func TestCreateToken_AdminRole(t *testing.T) {
	id := uint(2)
	role := "admin"
	name := "Admin User"
	phone := "09123456788"
	rememberMe := false

	token, expireTime, err := Utils.CreateToken(id, role, name, phone, rememberMe)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.True(t, expireTime.After(time.Now()))
}

func TestCreateToken_RememberMe(t *testing.T) {
	id := uint(1)
	role := "user"
	name := "Test User"
	phone := "09123456789"

	// Without remember me (1 month)
	_, expireTime1, _ := Utils.CreateToken(id, role, name, phone, false)

	// With remember me (6 months)
	_, expireTime2, _ := Utils.CreateToken(id, role, name, phone, true)

	assert.True(t, expireTime2.After(expireTime1), "Remember me should extend expiration")
}

func TestVerifyToken_Success(t *testing.T) {
	id := uint(1)
	role := "user"
	name := "Test User"
	phone := "09123456789"

	token, _, err := Utils.CreateToken(id, role, name, phone, false)
	assert.NoError(t, err)

	claims, err := Utils.VerifyToken(token)

	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, id, claims.ID)
	assert.Equal(t, role, claims.Role)
	assert.Equal(t, name, claims.Name)
	assert.Equal(t, phone, claims.Phone)
}

func TestVerifyToken_InvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"

	claims, err := Utils.VerifyToken(invalidToken)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestVerifyToken_EmptyToken(t *testing.T) {
	claims, err := Utils.VerifyToken("")

	assert.Error(t, err)
	assert.Nil(t, claims)
}
