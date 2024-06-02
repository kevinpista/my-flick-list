package services

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/kevinpista/my-flick-list/backend/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateWatchlistSuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first to fetch a valid UUID
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	userRegistrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, userRegistrationResponse, "The registration response should not be nil")
	require.NotEmpty(t, userRegistrationResponse.ID, "The registration response should include a non-empty user UUID")
	require.Equal(t, testUserRegister.Email, userRegistrationResponse.Email, "The registered user's email should match the input")

	watchlistService := WatchlistService{}

	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	watchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	assert.NoError(t, err, "CreateWatchList should return no error")
	assert.NotNil(t, watchlistResponse, "CreateWatchlist response should not be nil")

	assert.NotNil(t, watchlistResponse.ID, "CreateWatchlist response for watchlist ID should not be nil")
	assert.NotNil(t, watchlistResponse.UserID, "CreateWatchlist response for UserID should not be nil")
	assert.Equal(t, watchlistResponse.UserID, userRegistrationResponse.ID, "CreateWatchlist response for watchlist UserID should equal input UserID")
	assert.Equal(t, watchlistResponse.Name, testWatchlist.Name, "CreateWatchlist response for watchlist Name should equal input Name")
	assert.Equal(t, watchlistResponse.Description, testWatchlist.Description, "CreateWatchlist response for watchlist Description should equal input Description")
	assert.NotNil(t, watchlistResponse.CreatedAt, "CreateWatchlist response for CreatedAt should not be nil")
	assert.NotNil(t, watchlistResponse.UpdatedAt, "CreateWatchlist response for UpdatedAt should not be nil")

}

func TestCreateWatchlistInvalidUserID(t *testing.T) {
	// Test with a fake UUID
	invalidUUID := "511ca1a9-3ccf-4fe8-92da-111111111111"
	watchlistService := WatchlistService{}

	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	watchlistResponse, err := watchlistService.CreateWatchlist(uuid.MustParse(invalidUUID), testWatchlist)
	assert.Error(t, err, "CreateWatchList should return an error due to invalid UUID")
	assert.Nil(t, watchlistResponse, "CreateWatchlist response should be nil due to invalid UUID")
}

func TestGetAllWatchlists(t *testing.T) {
	watchlistService := WatchlistService{}

	watchlistResponse, err := watchlistService.GetAllWatchlists()
	assert.NoError(t, err, "GetAllWatchlists should not return an error")
	assert.NotNil(t, watchlistResponse, "GetAllWatchlists response should not be nil")
}

func TestGetAllWatchlistsByUserIDSuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first to fetch a valid UUID
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	userRegistrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, userRegistrationResponse, "The registration response should not be nil")
	require.NotEmpty(t, userRegistrationResponse.ID, "The registration response should include a non-empty user UUID")
	require.Equal(t, testUserRegister.Email, userRegistrationResponse.Email, "The registered user's email should match the input")

	watchlistService := WatchlistService{}

	// Insert test watchlist to be fetched
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")

	getWatchlistUserByIDResponse, err := watchlistService.GetAllWatchlistsByUserID(userRegistrationResponse.ID)
	assert.NoError(t, err, "GetAllWatchlistsByUserID should return no error")
	assert.NotNil(t, getWatchlistUserByIDResponse, "GetAllWatchlistsByUserID response should not be nil")
	assert.IsType(t, getWatchlistUserByIDResponse, []*models.Watchlist{}, "GetAllWatchlistsByUserID should be a type: slice of Watchlist items")
}

func TestGetAllWatchlistsByUserIDInvalidUserID(t *testing.T) {
	// Test with a fake UUID
	invalidUUID := "511ca1a9-3ccf-4fe8-92da-111111111111"
	watchlistService := WatchlistService{}

	watchlistResponse, err := watchlistService.GetAllWatchlistsByUserID(uuid.MustParse(invalidUUID))
	assert.Nil(t, err, "GetAllWatchlistsByUserID err should be nil due to nonexistent UUID")
	assert.IsType(t, watchlistResponse, []*models.Watchlist{}, "GetAllWatchlistsByUserID response should be an empty slice due to nonexistent UUID")
}

func TestGetWatchlistsByUserIDWithMovieCountInvalidUserISuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first to fetch a valid UUID
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	userRegistrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, userRegistrationResponse, "The registration response should not be nil")
	require.NotEmpty(t, userRegistrationResponse.ID, "The registration response should include a non-empty user UUID")
	require.Equal(t, testUserRegister.Email, userRegistrationResponse.Email, "The registered user's email should match the input")

	watchlistService := WatchlistService{}

	// Insert test watchlist to be fetched with count
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")

	getWatchlistUserByIDResponse, err := watchlistService.GetWatchlistsByUserIDWithMovieCount(userRegistrationResponse.ID)
	assert.NoError(t, err, "GetWatchlistsByUserIDWithMovieCount should return no error")
	assert.NotNil(t, getWatchlistUserByIDResponse, "GetWatchlistsByUserIDWithMovieCount response should not be nil")
	assert.IsType(t, getWatchlistUserByIDResponse, []*models.WatchlistWithItemCount{}, "GetAllWatchlistsByUserID should be a type: slice of WatchlistWithItemCount items")
}

func TestGetWatchlistsByUserIDWithMovieCountInvalidUserID(t *testing.T) {
	// Test with a fake UUID
	invalidUUID := "511ca1a9-3ccf-4fe8-92da-111111111111"
	watchlistService := WatchlistService{}

	watchlistResponse, err := watchlistService.GetWatchlistsByUserIDWithMovieCount(uuid.MustParse(invalidUUID))
	assert.Nil(t, err, "GetWatchlistsByUserIDWithMovieCount err should be nil due to nonexistent UUID")
	assert.IsType(t, watchlistResponse, []*models.WatchlistWithItemCount{}, "GetWatchlistsByUserIDWithMovieCount response should be an empty slice due to nonexistent UUID")
}

func TestGetWatchlistByIDSuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first to fetch a valid UUID
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	userRegistrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, userRegistrationResponse, "The registration response should not be nil")
	require.NotEmpty(t, userRegistrationResponse.ID, "The registration response should include a non-empty user UUID")
	require.Equal(t, testUserRegister.Email, userRegistrationResponse.Email, "The registered user's email should match the input")

	watchlistService := WatchlistService{}

	// Insert test watchlist to be fetched
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	getWatchlistByIDResponse, err := watchlistService.GetWatchlistByID(createWatchlistResponse.ID)
	assert.NoError(t, err, "GetWatchlistByID should return no error")
	assert.NotNil(t, getWatchlistByIDResponse, "GetWatchlistByID response should not be nil")
	assert.IsType(t, getWatchlistByIDResponse, &models.Watchlist{}, "GetWatchlistByID should be a type: Watchlist")
}

func TestGetWatchlistByIDInvalidID(t *testing.T) {
	watchlistService := WatchlistService{}
	getWatchlistUserByIDResponse, err := watchlistService.GetWatchlistByID(2147483646)
	assert.Error(t, err, "GetWatchlistByID should return an error due to nonexistent watchlist")
	assert.Error(t, err, "GetWatchlistByID should return an error due to nonexistent watchlist")
	assert.Equal(t, err, sql.ErrNoRows, "GetWatchlistByID error should return sql.ErrNoRows message due nonexistent watchlist")
	assert.Nil(t, getWatchlistUserByIDResponse, "GetWatchlistByID response should be nil due to nonexistent watchlist")
}

func TestDeleteWatchlistByIDSuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first to fetch a valid UUID
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	userRegistrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, userRegistrationResponse, "The registration response should not be nil")
	require.NotEmpty(t, userRegistrationResponse.ID, "The registration response should include a non-empty user UUID")
	require.Equal(t, testUserRegister.Email, userRegistrationResponse.Email, "The registered user's email should match the input")

	watchlistService := WatchlistService{}

	// Insert test watchlist to be deleted
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	deleteWatchlistByIDResponse := watchlistService.DeleteWatchlistByID(createWatchlistResponse.ID)
	assert.Nil(t, deleteWatchlistByIDResponse, "DeleteWatchlistByID should return nil due to successful deletion")
}

func TestGetWatchlistOwnerUserIDSuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first to fetch a valid UUID
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	userRegistrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, userRegistrationResponse, "The registration response should not be nil")
	require.NotEmpty(t, userRegistrationResponse.ID, "The registration response should include a non-empty user UUID")
	require.Equal(t, testUserRegister.Email, userRegistrationResponse.Email, "The registered user's email should match the input")

	watchlistService := WatchlistService{}

	// Insert test watchlist
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	getWatchlistOwnerByUserIDResponse, err := watchlistService.GetWatchlistOwnerUserID(createWatchlistResponse.ID)
	assert.NoError(t, err, "GetWatchlistOwnerUserID should return no error")
	assert.NotNil(t, getWatchlistOwnerByUserIDResponse, "GetWatchlistOwnerUserID response should not be nil")
	assert.Equal(t, createWatchlistResponse.UserID, getWatchlistOwnerByUserIDResponse, "GetWatchlistOwnerUserID should equal createWatchlistResponse.UserID")
}

func TestGetWatchlistOwnerUserIDInvalidID(t *testing.T) {
	watchlistService := WatchlistService{}
	getWatchlistOwnerByUserIDResponse, err := watchlistService.GetWatchlistOwnerUserID(2147483646)
	assert.Error(t, err, "GetWatchlistOwnerUserID should return an error due to nonexistent watchlist")
	assert.Equal(t, getWatchlistOwnerByUserIDResponse, uuid.Nil, "GetWatchlistOwnerUserID response should be UUID.nil due to nonexistent watchlist")
}

func TestUpdateWatchlistNameSuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first to fetch a valid UUID
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	userRegistrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, userRegistrationResponse, "The registration response should not be nil")
	require.NotEmpty(t, userRegistrationResponse.ID, "The registration response should include a non-empty user UUID")
	require.Equal(t, testUserRegister.Email, userRegistrationResponse.Email, "The registered user's email should match the input")

	watchlistService := WatchlistService{}

	// Insert test watchlist to be updated
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testUpdatedWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
	}

	updateWatchlistNameResponse, err := watchlistService.UpdateWatchlistName(createWatchlistResponse.ID, testUpdatedWatchlist)
	assert.NoError(t, err, "UpdateWatchlistName should return no error")
	assert.NotNil(t, updateWatchlistNameResponse, "UpdateWatchlistName response should not be nil")
	assert.Equal(t, updateWatchlistNameResponse.Name, testUpdatedWatchlist.Name, "UpdateWatchlistName response Name should equal input name")
	assert.NotEqual(t, updateWatchlistNameResponse.UpdatedAt, createWatchlistResponse.UpdatedAt, "UpdateWatchlistName response UpdatedAt should not equal createWatchlistResponse.UpdatedAt")

}

func TestUpdateWatchlistNameInvalidID(t *testing.T) {
	watchlistService := WatchlistService{}

	testUpdatedWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}
	updateWatchlistNameResponse, err := watchlistService.UpdateWatchlistName(2147483646, testUpdatedWatchlist)
	assert.Error(t, err, "UpdateWatchlistName should return an error due to nonexistent watchlist")
	assert.Nil(t, updateWatchlistNameResponse, "UpdateWatchlistName response should be nil due to nonexistent watchlist")
}


func TestUpdateWatchlistDescriptionSuccess(t *testing.T) {
	userService := UserService{}
	// Register test user first to fetch a valid UUID
	testUserRegister := models.User{
		Name:     util.RandomName(),
		Email:    util.RandomName() + "@example.com",
		Password: "testpassword",
	}

	userRegistrationResponse, err := userService.RegisterUser(testUserRegister)
	require.NoError(t, err, "RegisterUser should not return an error")
	require.NotNil(t, userRegistrationResponse, "The registration response should not be nil")
	require.NotEmpty(t, userRegistrationResponse.ID, "The registration response should include a non-empty user UUID")
	require.Equal(t, testUserRegister.Email, userRegistrationResponse.Email, "The registered user's email should match the input")

	watchlistService := WatchlistService{}

	// Insert test watchlist to be updated
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testUpdatedWatchlist := models.Watchlist{
		Description: util.RandomParagraph(),
	}

	updateWatchlistDescriptionResponse, err := watchlistService.UpdateWatchlistDescription(createWatchlistResponse.ID, testUpdatedWatchlist)
	assert.NoError(t, err, "UpdateWatchlistDescription should return no error")
	assert.NotNil(t, updateWatchlistDescriptionResponse, "UpdateWatchlistDescription response should not be nil")
	assert.Equal(t, updateWatchlistDescriptionResponse.Description, testUpdatedWatchlist.Description, "UpdateWatchlistDescription response Description should equal input Description")
	assert.NotEqual(t, updateWatchlistDescriptionResponse.UpdatedAt, createWatchlistResponse.UpdatedAt, "UpdateWatchlistDescription response UpdatedAt should not equal createWatchlistResponse.UpdatedAt")
}

func TestUpdateWatchlistDescriptionInvalidID(t *testing.T) {
	watchlistService := WatchlistService{}

	testUpdatedWatchlist := models.Watchlist{
		Description: util.RandomParagraph(),
	}
	updateWatchlistDescriptionResponse, err := watchlistService.UpdateWatchlistDescription(2147483646, testUpdatedWatchlist)
	assert.Error(t, err, "UpdateWatchlistDescription should return an error due to nonexistent watchlist")
	assert.Nil(t, updateWatchlistDescriptionResponse, "UpdateWatchlistDescription response should be nil due to nonexistent watchlist")
}
