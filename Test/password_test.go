package tests

import (
	"testing"

	"github.com/vahidlotfi71/online-store-api/Utils"
)

func TestGenerateHashPassword(t *testing.T) {
	password := "mypassword123"

	hash, err := Utils.GenerateHashPassword(password)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if hash == "" {
		t.Error("Expected hash, got empty string")
	}

	if hash == password {
		t.Error("Hash should not equal plain password")
	}
}

func TestVerifyPassword_Correct(t *testing.T) {
	password := "testpassword"
	hash, _ := Utils.GenerateHashPassword(password)

	err := Utils.VerifyPassword(password, hash)

	if err != nil {
		t.Errorf("Expected password to match, got error: %v", err)
	}
}

func TestVerifyPassword_Wrong(t *testing.T) {
	hash, _ := Utils.GenerateHashPassword("correct")

	err := Utils.VerifyPassword("wrong", hash)

	if err == nil {
		t.Error("Expected error for wrong password")
	}
}
