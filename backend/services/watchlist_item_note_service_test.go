package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/kevinpista/my-flick-list/backend/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateWatchlistItemNote(t *testing.T) {
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

	// Create test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}

	// Create test watchlist item to have a note added to
	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	watchlistItemNoteService := WatchlistItemNoteService{}

	testWatchlistItemNote := models.WatchlistItemNote{
		WatchlistItemID: createWatchlistItemResponse.ID,
		ItemNotes:     util.RandomParagraph(),
	}

	createWatchlistItemNoteResponse, err := watchlistItemNoteService.CreateWatchlistItemNote(testWatchlistItemNote)
	assert.NoError(t, err, "CreateWatchlistItemNote should not return an error")
	assert.Nil(t, err, "CreateWatchlistItemNote err should return nil")
	assert.NotNil(t, createWatchlistItemNoteResponse, "CreateWatchlistItemNote response should not be nil")
	assert.IsType(t, createWatchlistItemNoteResponse, &models.WatchlistItemNote{}, "CreateWatchlistItemNote should be a type: WatchlistItemNote")
	assert.Equal(t, createWatchlistItemNoteResponse.WatchlistItemID, testWatchlistItemNote.WatchlistItemID, "CreateWatchlistItemNote response WatchlistItemID field should equal input WatchlistItemID")
	assert.Equal(t, createWatchlistItemNoteResponse.ItemNotes, testWatchlistItemNote.ItemNotes, "CreateWatchlistItemNote response ItemNotes field should equal input ItemNotes")
}

func TestUpdateWatchlistItemNote(t *testing.T) {
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

	// Create test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}

	// Create test watchlist item to have a note added to
	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	watchlistItemNoteService := WatchlistItemNoteService{}

	// Create test item note first before updating
	testWatchlistItemNote := models.WatchlistItemNote{
		WatchlistItemID: createWatchlistItemResponse.ID,
		ItemNotes:     util.RandomParagraph(),
	}

	createWatchlistItemNoteResponse, err := watchlistItemNoteService.CreateWatchlistItemNote(testWatchlistItemNote)
	require.NoError(t, err, "CreateWatchlistItemNote should not return an error")
	require.Nil(t, err, "CreateWatchlistItemNote err should return nil")
	require.NotNil(t, createWatchlistItemNoteResponse, "CreateWatchlistItemNote response should not be nil")
	require.IsType(t, createWatchlistItemNoteResponse, &models.WatchlistItemNote{}, "CreateWatchlistItemNote should be a type: WatchlistItemNote")
	require.Equal(t, createWatchlistItemNoteResponse.WatchlistItemID, testWatchlistItemNote.WatchlistItemID, "CreateWatchlistItemNote response WatchlistItemID field should equal input WatchlistItemID")
	require.Equal(t, createWatchlistItemNoteResponse.ItemNotes, testWatchlistItemNote.ItemNotes, "CreateWatchlistItemNote response ItemNotes field should equal input ItemNotes")


	testWatchlistItemNoteUpdate := models.WatchlistItemNote{
		WatchlistItemID: testWatchlistItemNote.WatchlistItemID,
		ItemNotes:     util.RandomParagraph(),
	}

	updateWatchlistItemNoteResponse, err := watchlistItemNoteService.UpdateWatchlistItemNote(testWatchlistItemNoteUpdate)
	assert.NoError(t, err, "UpdateWatchlistItemNote should not return an error")
	assert.Nil(t, err, "UpdateWatchlistItemNote err should return nil")
	assert.NotNil(t, updateWatchlistItemNoteResponse, "UpdateWatchlistItemNote response should not be nil")
	assert.IsType(t, updateWatchlistItemNoteResponse, &models.WatchlistItemNote{}, "UpdateWatchlistItemNote should be a type: WatchlistItemNote")
	assert.Equal(t, updateWatchlistItemNoteResponse.WatchlistItemID, testWatchlistItemNoteUpdate.WatchlistItemID, "UpdateWatchlistItemNote response WatchlistItemID field should equal input WatchlistItemID")
	assert.Equal(t, updateWatchlistItemNoteResponse.ItemNotes, testWatchlistItemNoteUpdate.ItemNotes, "UpdateWatchlistItemNote response ItemNotes field should equal input ItemNotes")
}

func TestGetWatchlistItemNoteByWatchlistItemID(t *testing.T) {
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

	// Create test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}

	// Create test watchlist item to have a note added to
	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	watchlistItemNoteService := WatchlistItemNoteService{}

	// Create test item note first before updating
	testWatchlistItemNote := models.WatchlistItemNote{
		WatchlistItemID: createWatchlistItemResponse.ID,
		ItemNotes:     util.RandomParagraph(),
	}

	createWatchlistItemNoteResponse, err := watchlistItemNoteService.CreateWatchlistItemNote(testWatchlistItemNote)
	require.NoError(t, err, "CreateWatchlistItemNote should not return an error")
	require.Nil(t, err, "CreateWatchlistItemNote err should return nil")
	require.NotNil(t, createWatchlistItemNoteResponse, "CreateWatchlistItemNote response should not be nil")
	require.IsType(t, createWatchlistItemNoteResponse, &models.WatchlistItemNote{}, "CreateWatchlistItemNote should be a type: WatchlistItemNote")
	require.Equal(t, createWatchlistItemNoteResponse.WatchlistItemID, testWatchlistItemNote.WatchlistItemID, "CreateWatchlistItemNote response WatchlistItemID field should equal input WatchlistItemID")
	require.Equal(t, createWatchlistItemNoteResponse.ItemNotes, testWatchlistItemNote.ItemNotes, "CreateWatchlistItemNote response ItemNotes field should equal input ItemNotes")


	testGetWatchlistItemNote := models.WatchlistItemNote{
		WatchlistItemID: testWatchlistItemNote.WatchlistItemID,
	}

	getWatchlistItemNoteResponse, err := watchlistItemNoteService.GetWatchlistItemNoteByWatchlistItemID(testGetWatchlistItemNote.WatchlistItemID)
	assert.NoError(t, err, "GetWatchlistItemNoteByWatchlistItemID should not return an error")
	assert.Nil(t, err, "GetWatchlistItemNoteByWatchlistItemID err should be nil")
	assert.NotNil(t, getWatchlistItemNoteResponse, "GetWatchlistItemNoteByWatchlistItemID response should not be nil")
	assert.IsType(t, getWatchlistItemNoteResponse, &models.WatchlistItemNote{}, "GetWatchlistItemNoteByWatchlistItemID should be a type: WatchlistItemNote")
	assert.Equal(t, getWatchlistItemNoteResponse.WatchlistItemID, testGetWatchlistItemNote.WatchlistItemID, "GetWatchlistItemNoteByWatchlistItemID response WatchlistItemID field should equal input WatchlistItemID")
	assert.Equal(t, getWatchlistItemNoteResponse.ItemNotes, testWatchlistItemNote.ItemNotes, "GetWatchlistItemNoteByWatchlistItemID response ItemNotes field should equal input ItemNotes")
}

func TestCheckIfWatchlistItemExistsTrue(t *testing.T) {
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

	// Create test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}

	// Create test watchlist item to have a note added to
	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	watchlistItemNoteService := WatchlistItemNoteService{}

	checkIfWatchlistItemExistsResponse, err := watchlistItemNoteService.CheckIfWatchlistItemExists(createWatchlistItemResponse.ID)
	assert.NoError(t, err, "CheckIfWatchlistItemExists should not return an error")
	assert.Nil(t, err, "CheckIfWatchlistItemExists err should return nil")
	assert.Equal(t, checkIfWatchlistItemExistsResponse, true, "CheckIfWatchlistItemExists response should equal true")
}

func TestCheckIfWatchlistItemExistsFalse(t *testing.T) {
	watchlistItemNoteService := WatchlistItemNoteService{}

	checkIfWatchlistItemExistsResponse, err := watchlistItemNoteService.CheckIfWatchlistItemExists(29992292)
	assert.NoError(t, err, "CheckIfWatchlistItemExists should not return an error")
	assert.Nil(t, err, "CheckIfWatchlistItemExists err should return nil")
	assert.Equal(t, checkIfWatchlistItemExistsResponse, false, "CheckIfWatchlistItemExists response should equal false")
}

func TestCheckIfWatchlistItemNoteExistsTrue(t *testing.T) {
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

	// Create test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}

	// Create test watchlist item to have a note added to
	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	watchlistItemNoteService := WatchlistItemNoteService{}

	// Create test item note first before updating
	testWatchlistItemNote := models.WatchlistItemNote{
		WatchlistItemID: createWatchlistItemResponse.ID,
		ItemNotes:     util.RandomParagraph(),
	}

	createWatchlistItemNoteResponse, err := watchlistItemNoteService.CreateWatchlistItemNote(testWatchlistItemNote)
	require.NoError(t, err, "CreateWatchlistItemNote should not return an error")
	require.Nil(t, err, "CreateWatchlistItemNote err should return nil")
	require.NotNil(t, createWatchlistItemNoteResponse, "CreateWatchlistItemNote response should not be nil")
	require.IsType(t, createWatchlistItemNoteResponse, &models.WatchlistItemNote{}, "CreateWatchlistItemNote should be a type: WatchlistItemNote")
	require.Equal(t, createWatchlistItemNoteResponse.WatchlistItemID, testWatchlistItemNote.WatchlistItemID, "CreateWatchlistItemNote response WatchlistItemID field should equal input WatchlistItemID")
	require.Equal(t, createWatchlistItemNoteResponse.ItemNotes, testWatchlistItemNote.ItemNotes, "CreateWatchlistItemNote response ItemNotes field should equal input ItemNotes")

	checkIfWatchlistItemNoteExistsResponse, err := watchlistItemNoteService.CheckIfWatchlistItemNoteExists(testWatchlistItemNote.WatchlistItemID)
	assert.NoError(t, err, "CheckIfWatchlistItemNoteExists should not return an error")
	assert.Nil(t, err, "CheckIfWatchlistItemNoteExists err should return nil")
	assert.Equal(t, checkIfWatchlistItemNoteExistsResponse, true, "CheckIfWatchlistItemNoteExists response should equal true")

}

func TestCheckIfWatchlistItemNoteExistsFalse(t *testing.T) {
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

	// Create test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}

	// Create test watchlist item
	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	watchlistItemNoteService := WatchlistItemNoteService{}

	checkIfWatchlistItemNoteExistsResponse, err := watchlistItemNoteService.CheckIfWatchlistItemNoteExists(testWatchlistItem.WatchlistID)
	assert.NoError(t, err, "CheckIfWatchlistItemNoteExists should not return an error")
	assert.Nil(t, err, "CheckIfWatchlistItemNoteExists err should return nil")
	assert.Equal(t, checkIfWatchlistItemNoteExistsResponse, false, "CheckIfWatchlistItemNoteExists response should equal false")
}

func TestCheckIfUserOwnsWatchlistItemTrue(t *testing.T) {
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

	// Create test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}

	// Create test watchlist item to have a note added to
	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	watchlistItemNoteService := WatchlistItemNoteService{}

	checkIfUserOwnsWatchlistItemResponse, err := watchlistItemNoteService.CheckIfUserOwnsWatchlistItem(userRegistrationResponse.ID, createWatchlistItemResponse.ID)
	assert.NoError(t, err, "CheckIfUserOwnsWatchlistItem should not return an error")
	assert.Nil(t, err, "CheckIfUserOwnsWatchlistItem err should return nil")
	assert.Equal(t, checkIfUserOwnsWatchlistItemResponse, true, "CheckIfUserOwnsWatchlistItem response should equal true")
}

func TestCheckIfUserOwnsWatchlistItemFalse(t *testing.T) {
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

	// Create test watchlist to have an item added to
	testWatchlist := models.Watchlist{
		Name:        util.RandomName() + " " + util.RandomName(),
		Description: util.RandomParagraph(),
	}

	createWatchlistResponse, err := watchlistService.CreateWatchlist(userRegistrationResponse.ID, testWatchlist)
	require.NoError(t, err, "CreateWatchList should return no error")
	require.NotNil(t, createWatchlistResponse, "CreateWatchlist response should not be nil")
	require.NotNil(t, createWatchlistResponse.ID, "CreateWatchlist.ID should not be nil")

	watchlistItemService := WatchlistItemService{}

	// Create test watchlist item to have a note added to
	testWatchlistItem := models.WatchlistItem{
		WatchlistID: createWatchlistResponse.ID,
		MovieID:     1895,
		Checkmarked: true,
	}

	createWatchlistItemResponse, err := watchlistItemService.CreateWatchlistItemByWatchlistID(testWatchlistItem)
	require.NoError(t, err, "CreateWatchlistItemByWatchlistID should return no error")
	require.NotNil(t, createWatchlistItemResponse, "CreateWatchlistItemByWatchlistID response should not be nil")
	require.NotNil(t, createWatchlistItemResponse.ID, "CreateWatchlistItemByWatchlistID response ID field should not be nil")
	require.Equal(t, createWatchlistItemResponse.WatchlistID, testWatchlistItem.WatchlistID, "CreateWatchlistItemByWatchlistID response WatchlistID field should equal input WatchlistID")
	require.Equal(t, createWatchlistItemResponse.MovieID, testWatchlistItem.MovieID, "CreateWatchlistItemByWatchlistID response MovieID field should equal input MovieID")

	watchlistItemNoteService := WatchlistItemNoteService{}
	
	invalidUUID := "511ca1a9-3ccf-4fe8-92da-111111111111"

	checkIfUserOwnsWatchlistItemResponse, err := watchlistItemNoteService.CheckIfUserOwnsWatchlistItem(uuid.MustParse(invalidUUID), createWatchlistItemResponse.ID)
	assert.NoError(t, err, "CheckIfUserOwnsWatchlistItem should not return an error")
	assert.Nil(t, err, "CheckIfUserOwnsWatchlistItem err should return nil")
	assert.Equal(t, checkIfUserOwnsWatchlistItemResponse, false, "CheckIfUserOwnsWatchlistItem response should equal false")
}
