package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Models/User"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Helper function to create in-memory test database
func setupUserTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Auto migrate models
	err = db.AutoMigrate(&Models.User{})
	assert.NoError(t, err)

	return db
}

func TestUserCreateDTO(t *testing.T) {
	dto := User.UserCreateDTO{
		FirstName:  "John",
		LastName:   "Doe",
		Phone:      "09123456789",
		Address:    "Test Address",
		NationalID: "1234567890",
		Password:   "hashedpassword",
	}

	assert.Equal(t, "John", dto.FirstName)
	assert.Equal(t, "Doe", dto.LastName)
	assert.Equal(t, "09123456789", dto.Phone)
}

func TestCreate(t *testing.T) {
	db := setupUserTestDB(t)

	dto := User.UserCreateDTO{
		FirstName:  "John",
		LastName:   "Doe",
		Phone:      "09123456789",
		Address:    "123 Test Street",
		NationalID: "1234567890",
		Password:   "hashedpassword123",
	}

	user, err := User.Create(db, dto)

	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "09123456789", user.Phone)
	assert.Equal(t, "user", user.Role)
	assert.False(t, user.IsVerified)
}

func TestFindByID(t *testing.T) {
	db := setupUserTestDB(t)

	// Create a user first
	dto := User.UserCreateDTO{
		FirstName:  "Jane",
		LastName:   "Smith",
		Phone:      "09123456788",
		Address:    "456 Test Avenue",
		NationalID: "0987654321",
		Password:   "hashedpassword456",
	}

	createdUser, err := User.Create(db, dto)
	assert.NoError(t, err)

	// Find the user
	foundUser, err := User.FindByID(db, createdUser.ID)

	assert.NoError(t, err)
	assert.Equal(t, createdUser.ID, foundUser.ID)
	assert.Equal(t, "Jane", foundUser.FirstName)
	assert.Equal(t, "Smith", foundUser.LastName)
}

func TestFindUserByID_NotFound(t *testing.T) {
	db := setupUserTestDB(t)

	_, err := User.FindByID(db, 9999)

	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUpdate(t *testing.T) {
	db := setupUserTestDB(t)

	// Create a user
	createDTO := User.UserCreateDTO{
		FirstName:  "Old",
		LastName:   "Name",
		Phone:      "09111111111",
		Address:    "Old Address",
		NationalID: "1111111111",
		Password:   "oldpassword",
	}

	user, err := User.Create(db, createDTO)
	assert.NoError(t, err)

	// Update the user
	updateDTO := User.UserUpdateDTO{
		FirstName:  "New",
		LastName:   "Name",
		Phone:      "09222222222",
		Address:    "New Address",
		NationalID: "2222222222",
		Password:   "newpassword",
	}

	err = User.Update(db, user.ID, updateDTO)

	assert.NoError(t, err)

	// Verify update
	updatedUser, _ := User.FindByID(db, user.ID)
	assert.Equal(t, "New", updatedUser.FirstName)
	assert.Equal(t, "09222222222", updatedUser.Phone)
}

func TestUserUpdate_NotFound(t *testing.T) {
	db := setupUserTestDB(t)

	updateDTO := User.UserUpdateDTO{
		FirstName: "Test",
		LastName:  "User",
	}

	err := User.Update(db, 9999, updateDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestSoftDelete(t *testing.T) {
	db := setupUserTestDB(t)

	// Create a user
	dto := User.UserCreateDTO{
		FirstName:  "Delete",
		LastName:   "Me",
		Phone:      "09333333333",
		Address:    "Delete Address",
		NationalID: "3333333333",
		Password:   "deletepass",
	}

	user, err := User.Create(db, dto)
	assert.NoError(t, err)

	// Soft delete
	err = User.SoftDelete(db, user.ID)

	assert.NoError(t, err)

	// Try to find (should not be found with default scope)
	_, err = User.FindByID(db, user.ID)
	assert.Error(t, err)

	// Verify with Unscoped
	var deletedUser Models.User
	err = db.Unscoped().First(&deletedUser, user.ID).Error
	assert.NoError(t, err)
	assert.False(t, deletedUser.DeletedAt.Time.IsZero())
}

func TestUserSoftDelete_NotFound(t *testing.T) {
	db := setupUserTestDB(t)

	err := User.SoftDelete(db, 9999)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestUpdate_PasswordOptional(t *testing.T) {
	db := setupUserTestDB(t)

	// Create user
	createDTO := User.UserCreateDTO{
		FirstName:  "Test",
		LastName:   "User",
		Phone:      "09444444444",
		Address:    "Test Address",
		NationalID: "4444444444",
		Password:   "originalpassword",
	}

	user, _ := User.Create(db, createDTO)
	originalPassword := user.Password

	// Update without password
	updateDTO := User.UserUpdateDTO{
		FirstName:  "Updated",
		LastName:   "User",
		Phone:      "09555555555",
		Address:    "Updated Address",
		NationalID: "5555555555",
		Password:   "", // Empty password
	}

	err := User.Update(db, user.ID, updateDTO)
	assert.NoError(t, err)

	// Verify password unchanged
	updatedUser, _ := User.FindByID(db, user.ID)
	assert.Equal(t, originalPassword, updatedUser.Password)
	assert.Equal(t, "Updated", updatedUser.FirstName)
}

func TestCreate_DuplicatePhone(t *testing.T) {
	db := setupUserTestDB(t)

	dto := User.UserCreateDTO{
		FirstName:  "First",
		LastName:   "User",
		Phone:      "09666666666",
		Address:    "Address 1",
		NationalID: "6666666666",
		Password:   "password1",
	}

	_, err := User.Create(db, dto)
	assert.NoError(t, err)

	// Try to create another user with same phone
	dto2 := User.UserCreateDTO{
		FirstName:  "Second",
		LastName:   "User",
		Phone:      "09666666666", // Same phone
		Address:    "Address 2",
		NationalID: "7777777777",
		Password:   "password2",
	}

	_, err = User.Create(db, dto2)
	assert.Error(t, err, "Should fail on duplicate phone")
}

func TestCreate_DuplicateNationalID(t *testing.T) {
	db := setupUserTestDB(t)

	dto := User.UserCreateDTO{
		FirstName:  "First",
		LastName:   "User",
		Phone:      "09777777777",
		Address:    "Address 1",
		NationalID: "8888888888",
		Password:   "password1",
	}

	_, err := User.Create(db, dto)
	assert.NoError(t, err)

	// Try to create another user with same national ID
	dto2 := User.UserCreateDTO{
		FirstName:  "Second",
		LastName:   "User",
		Phone:      "09888888888",
		Address:    "Address 2",
		NationalID: "8888888888", // Same national ID
		Password:   "password2",
	}

	_, err = User.Create(db, dto2)
	assert.Error(t, err, "Should fail on duplicate national ID")
}

func TestUserModel_TableName(t *testing.T) {
	user := Models.User{}
	assert.Equal(t, "users", user.TableName())
}

func TestUserModel_DefaultValues(t *testing.T) {
	db := setupUserTestDB(t)

	dto := User.UserCreateDTO{
		FirstName:  "Test",
		LastName:   "User",
		Phone:      "09999999999",
		Address:    "Test",
		NationalID: "9999999999",
		Password:   "test",
	}

	user, err := User.Create(db, dto)

	assert.NoError(t, err)
	assert.Equal(t, "user", user.Role)
	assert.False(t, user.IsVerified)
	assert.NotZero(t, user.CreatedAt)
	assert.NotZero(t, user.UpdatedAt)
}
