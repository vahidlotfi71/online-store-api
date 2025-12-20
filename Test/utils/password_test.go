package Test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vahidlotfi71/online-store-api/Utils"
)

func TestGenerateHashPassword(t *testing.T) {
	password := "mypassword123"

	hash, err := Utils.GenerateHashPassword(password)

	assert.NoError(t, err, "Hash generation should not return error")
	assert.NotEmpty(t, hash, "Hash should not be empty")
	assert.NotEqual(t, password, hash, "Hash should be different from password")
}

func TestVerifyPassword_Success(t *testing.T) {
	password := "mypassword123"
	hash, _ := Utils.GenerateHashPassword(password)

	err := Utils.VerifyPassword(password, hash)

	assert.NoError(t, err, "Password verification should succeed")
}

func TestVerifyPassword_Failure(t *testing.T) {
	password := "mypassword123"
	wrongPassword := "wrongpassword"
	hash, _ := Utils.GenerateHashPassword(password)

	err := Utils.VerifyPassword(wrongPassword, hash)

	assert.Error(t, err, "Password verification should fail with wrong password")
}

func TestGenerateHashPassword_EmptyString(t *testing.T) {
	password := ""

	hash, err := Utils.GenerateHashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}
