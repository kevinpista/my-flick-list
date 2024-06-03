package services

import (
	"testing"
	// "database/sql"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/kevinpista/my-flick-list/backend/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateWatchlistItemByWatchlistIDSuccess(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	assert.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	assert.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	assert.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	assert.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	assert.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")
}

func TestCreateWatchlistItemByWatchlistIDInvalidMovieID(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     9992999,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	assert.Error(t, err, "CreateWatchlistItemByWatchlistID should return an error due to Movie ID not existing in TMDB")
	assert.Nil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should return nil due to Movie ID not existing in TMDB")
}

func TestGetAllWatchlistItemsByWatchlistIDSuccess(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	getAllWatchlistItemResponse, err := watchlistItemService.GetAllWatchlistItemsByWatchlistID(createWatchlistResponse.ID)
	assert.NoError(t, err, "GetAllWatchlistItemsByWatchlistID should return no error")
	assert.NotNil(t, getAllWatchlistItemResponse, "GetAllWatchlistItemsByWatchlistID response should not be nil")
	assert.Len(t, getAllWatchlistItemResponse, 1, "GetAllWatchlistItemsByWatchlistID response should be length 1 due to 1 movie being added")
}

func TestGetWatchlistWithWatchlistItemsByWatchlistIDSuccess(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	getWatchlistItemsMoviesResponse, getWatchlistItemsNameResponse, getWatchlistItemsDescriptionResponse, err := watchlistItemService.GetWatchlistWithWatchlistItemsByWatchlistID(createWatchlistResponse.ID)
	assert.NoError(t, err, "GetWatchlistWithWatchlistItemsByWatchlistID should return no error")
	assert.NotNil(t, getWatchlistItemsMoviesResponse, "GetWatchlistWithWatchlistItemsByWatchlistID movies response should not be nil")
	assert.NotNil(t, getWatchlistItemsNameResponse, "GetWatchlistWithWatchlistItemsByWatchlistID name response should not be nil")
	assert.NotNil(t, getWatchlistItemsDescriptionResponse, "GetWatchlistWithWatchlistItemsByWatchlistID description response should not be nil")
	assert.Equal(t, getWatchlistItemsNameResponse, createWatchlistResponse.Name, "GetWatchlistWithWatchlistItemsByWatchlistID name response should equal input")
	assert.Equal(t, getWatchlistItemsDescriptionResponse, createWatchlistResponse.Description, "GetWatchlistWithWatchlistItemsByWatchlistID description response should equal input")
}

func TestDeleteWatchlistItemByIDSuccess(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	deleteWatchlistItemResponse := watchlistItemService.DeleteWatchlistItemByID(createWatchlistItemResponse.ID, createWatchlistResponse.ID)
	assert.NoError(t, deleteWatchlistItemResponse, "DeleteWatchlistItemByID should return no error")
	assert.Nil(t, deleteWatchlistItemResponse, "DeleteWatchlistItemByID response should not be nil")
}

func TestUpdateCheckmarkedBooleanByWatchlistItemByIDSuccess(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	testWatchlistItemBooleanUpdate := models.WatchlistItem{
		Checkmarked: false,
	}

	updateWatchlistItemCheckmarkResponse := watchlistItemService.UpdateCheckmarkedBooleanByWatchlistItemByID(createWatchlistItemResponse.ID, testWatchlistItemBooleanUpdate, createWatchlistResponse.ID)
	assert.NoError(t, updateWatchlistItemCheckmarkResponse, "UpdateCheckmarkedBooleanByWatchlistItemByID should return no error")
	assert.Nil(t, updateWatchlistItemCheckmarkResponse, "UpdateCheckmarkedBooleanByWatchlistItemByID response should not be nil")
}

func TestCheckIfMovieInWatchlistExistsTrue(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	checkIfMovieInWatchlistExistsResponse, err := watchlistItemService.CheckIfMovieInWatchlistExists(createWatchlistResponse.ID, createWatchlistItemResponse.MovieID)
	assert.NoError(t, err, "CheckIfMovieInWatchlistExists should return no error")
	assert.Nil(t, err, "CheckIfMovieInWatchlistExists err should be nil")
	assert.Equal(t, checkIfMovieInWatchlistExistsResponse, true, "CheckIfMovieInWatchlistExists boolean response should be true")
}

func TestCheckIfMovieInWatchlistExistsFalse(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	checkIfMovieInWatchlistExistsResponse, err := watchlistItemService.CheckIfMovieInWatchlistExists(createWatchlistResponse.ID, 9999992)
	assert.NoError(t, err, "CheckIfMovieInWatchlistExists should return no error")
	assert.Nil(t, err, "CheckIfMovieInWatchlistExists err should be nil")
	assert.Equal(t, checkIfMovieInWatchlistExistsResponse, false, "CheckIfMovieInWatchlistExists boolean response should be false")
}

func TestCheckIfWatchlistExistsTrue(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}
	checkIfWatchlistExistsResponse, err := watchlistItemService.CheckIfWatchlistExists(createWatchlistResponse.ID)
	assert.NoError(t, err, "CheckIfWatchlistExists should return no error")
	assert.Nil(t, err, "CheckIfWatchlistExists err should be nil")
	assert.Equal(t, checkIfWatchlistExistsResponse, true, "CheckIfWatchlistExists boolean response should be true")
}

func TestCheckIfWatchlistExistsFalse(t *testing.T) {
	watchlistItemService := WatchlistItemService{}
	checkIfWatchlistExistsResponse, err := watchlistItemService.CheckIfWatchlistExists(9999992)
	assert.NoError(t, err, "CheckIfWatchlistExists should return no error")
	assert.Nil(t, err, "CheckIfWatchlistExists err should be nil")
	assert.Equal(t, checkIfWatchlistExistsResponse, false, "CheckIfWatchlistExists boolean response should be false")
}

func TestGetWatchlistItemOwnerUserIDSuccess(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}
	checkWatchlistOwnerIDResponse, err := watchlistItemService.GetWatchlistOwnerUserID(createWatchlistResponse.ID)
	assert.NoError(t, err, "GetWatchlistOwnerUserID should return no error")
	assert.Nil(t, err, "GetWatchlistOwnerUserID err should be nil")
	assert.Equal(t, checkWatchlistOwnerIDResponse, userRegistrationResponse.ID, "GetWatchlistOwnerUserID response should equal userRegistrationResponse.ID")
}

func TestGetWatchlistItemOwnerUserIDInvalidID(t *testing.T) {
	watchlistItemService := WatchlistItemService{}
	checkWatchlistOwnerIDResponse, err := watchlistItemService.GetWatchlistOwnerUserID(2147483646)
	assert.Error(t, err, "GetWatchlistOwnerUserID should return error due to nonexistent watchlist")
	assert.NotNil(t, err, "GetWatchlistOwnerUserID err should not be nil due to nonexistent watchlist")
	assert.Equal(t, checkWatchlistOwnerIDResponse, uuid.Nil, "GetWatchlistOwnerUserID response should be UUID.nil due to nonexistent watchlist")
}

func TestGetWatchlistItemWatchlistIdSuccess(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	getWatchlistItemWatchlistIdResponse, err := watchlistItemService.GetWatchlistItemWatchlistId(createWatchlistItemResponse.ID)
	assert.NoError(t, err, "GetWatchlistItemWatchlistId should return no error")
	assert.Nil(t, err, "GetWatchlistItemWatchlistId err should be nil")
	assert.NotNil(t, getWatchlistItemWatchlistIdResponse, "GetWatchlistItemWatchlistId response should not be nil")
	assert.Equal(t, getWatchlistItemWatchlistIdResponse, createWatchlistResponse.ID, "GetWatchlistItemWatchlistId response should equal createWatchlistResponse.ID")
}

func TestGetWatchlistItemWatchlistIdInvalidID(t *testing.T) {
	watchlistItemService := WatchlistItemService{}
	getWatchlistItemWatchlistIDResponse, err := watchlistItemService.GetWatchlistItemWatchlistId(2147483646)
	assert.Error(t, err, "GetWatchlistItemWatchlistId should return error due to nonexistent watchlist item")
	assert.NotNil(t, err, "GetWatchlistItemWatchlistId err should not be nil due to nonexistent watchlist item")
	assert.Equal(t, getWatchlistItemWatchlistIDResponse, 0, "GetWatchlistItemWatchlistId response should be 0 due to nonexistent watchlist")
}

func TestCheckIfMovieExistsTrue(t *testing.T) {
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

	// Insert test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.Name, "CreateWatchlist.Name should not be nil")
	require.NotNil(t, createWatchlistResponse.Description, "CreateWatchlist.Description should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	watchlistItemService := WatchlistItemService{}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	checkIfMovieExistsResponse, err := watchlistItemService.CheckIfMovieExists(createWatchlistItemResponse.MovieID)
	assert.NoError(t, err, "CheckIfMovieExists should return no error")
	assert.Nil(t, err, "CheckIfMovieExists err should be nil")
	assert.Equal(t, checkIfMovieExistsResponse, true, "CheckIfMovieExists response should equal true")
}

func TestCheckIfMovieExistsFalse(t *testing.T) {
	watchlistItemService := WatchlistItemService{}
	checkIfMovieExistsResponse, err := watchlistItemService.CheckIfMovieExists(9992992)
	assert.NoError(t, err, "CheckIfMovieExists should return no error")
	assert.Nil(t, err, "CheckIfMovieExists err should be nil")
	assert.Equal(t, checkIfMovieExistsResponse, false, "CheckIfMovieExists response should equal false")
}
