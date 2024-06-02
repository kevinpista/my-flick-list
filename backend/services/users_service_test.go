package services

import (
	"testing"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/kevinpista/my-flick-list/backend/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Initialize the UserService instance; implicitly uses the global 'db' to access test database configured in main_test.go.
// Seeded random data
func TestRegisterUserSuccess(t *testing.T) {
	userService := UserService{}

	testUser := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword", // This should be plain text as hashing is handled within the service
	}

	registrationResponse, err := userService.RegisterUser(testUser)

	// Check errors
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, registrationResponse, "The registration response should not be nil")

	// Verify data returned
	assert.NotEmpty(t, registrationResponse.ID, "The registration response should include a non-empty UUID")
	assert.Equal(t, testUser.Email, registrationResponse.Email, "The registered user's email should match the input")
}

func TestHandleLoginSuccess(t *testing.T) {
	userService := UserService{}

	testUser := models.User{
		Email:    "test@example.com",
		Password: "testpassword",
	}

	loginResponse, err := userService.HandleLogin(testUser)

	require.NoError(t, err, "HandleLogin should not return an error")
	require.NotNil(t, loginResponse, "The login response should not be nil")

	assert.NotEmpty(t, loginResponse.ID, "The logged in user's ID response should include a non-empty ID")
	assert.Equal(t, testUser.Email, loginResponse.Email, "Logged in user email should match")
}

func TestHandleLoginUserNotFound(t *testing.T) {
	userService := UserService{}

	testUser := models.User{
		Email:    util.RandomName() + "@example.com",
		Password: "arbitrary_password",
	}

	loginResponse, err := userService.HandleLogin(testUser)

	assert.Error(t, err, "HandleLogin should return an error due to user not being found")
	assert.Nil(t, loginResponse, "The login response should be nil")
}

func TestHandleLoginIncorrectPassword(t *testing.T) {
	userService := UserService{}
	// Register test user first
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	registrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, registrationResponse, "The registration response should not be nil")

	require.NotEmpty(t, registrationResponse.ID, "The registration response should include a non-empty UUID")
	require.Equal(t, testUserRegister.Email, registrationResponse.Email, "The registered user's email should match the input")

	// Test password of newly registered user is valid
	testUserLogin := models.User{
		Email:    testUserRegister.Email,
		Password: "testpassword",
	}

	loginResponse, err := userService.HandleLogin(testUserLogin)
	require.NoError(t, err, "HandleLogin should not return an error after testing newly registered user")
	require.NotNil(t, loginResponse, "The login response should not be nil after testing newly registered user")

	// Test with incorrect password
	testUserLoginWithIncorrectPassword := models.User{
		Email:    testUserRegister.Email,
		Password: "wrongpassword",
	}

	loginResponse, err = userService.HandleLogin(testUserLoginWithIncorrectPassword)
	assert.Error(t, err, "HandleLogin should return an error due to incorrect password")
	assert.Nil(t, loginResponse, "The login response should be nil due to incorrect password")

}

func TestGetUserByIDSuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first
	testUser := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	registrationResponse, err := userService.RegisterUser(testUser)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, registrationResponse, "The registration response should not be nil")

	require.NotEmpty(t, registrationResponse.ID, "The registration response should include a non-empty UUID")
	require.Equal(t, testUser.Email, registrationResponse.Email, "The registered user's email should match the input")

	// Test GetUserByID function and returned data against testUserRegister data
	getUserResponse, err := userService.GetUserByID(registrationResponse.ID)
	require.NoError(t, err, "GetUserByID should not return any error")
	require.NotNil(t, getUserResponse, "GetUserByID response object should not be nil")

	assert.Equal(t, testUser.Name, getUserResponse.Name, "The GetUserByID name should match the input")
	assert.Equal(t, testUser.Email, getUserResponse.Email, "The GetUserByID email should match the input")
	assert.Equal(t, registrationResponse.ID, getUserResponse.ID, "The GetUserByID UUID should match the input")
	assert.NotNil(t, getUserResponse.CreatedAt, "GetUserByID CreatedAt field should not be nil")
	assert.NotNil(t, getUserResponse.UpdatedAt, "GetUserByID UpdatedAt field should not be nil")
}

func TestGetUserByIDFailure(t *testing.T) {
	userService := UserService{}

	invalidUUID := "511ca1a9-3ccf-4fe8-92da-111111111111"
	getUserResponse, err := userService.GetUserByID(uuid.MustParse(invalidUUID))

	assert.Error(t, err, "GetUserByID should return an error due to invalid UUID parameter")
	assert.Equal(t, err, sql.ErrNoRows, "GetUserByID error should return sql.ErrNoRows message due to invalid UUID parameter")
	assert.Nil(t, getUserResponse, "GetUserByID response object should be nil due to invlaid UUID parameter")
}