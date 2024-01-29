// API requests related to watchlists, their watchlist_items, and item_notes

import axios from 'axios';
import { getJwtTokenFromCookies } from '../utils/authTokenUtils'


// Returns watchlist name, description and its watchlist items
// router.GET("/api/watchlist-items-with-movies, controllers.GetAllWatchlistItemsWithMoviesByWatchListID) 
export function fetchWatchlistAndItemsAPI (watchlistID) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    
    return axios.get('http://localhost:8080/api/watchlist-items-with-movies', {
        headers, 
        params: {
            watchlistID: watchlistID,
        },
      })
        .then(response => {
            return response.data; // Returning watchlist item's movie data + watchlist name & description
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

// router.DELETE("/api/watchlist-item?id={watchlistItemID}, controllers.DeleteWatchlistItemByID(id)
// Deletes watchlist_item via its id
export function deleteWatchlistItemAPI (watchlistItemID) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    
    return axios.delete('http://localhost:8080/api/watchlist-item', {
        headers, 
        params: {
            id: watchlistItemID,
        },
        })
        .then(response => {
            return response.data; // Returning success message "mesage: Watchlist item deleted successfully"
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

// Returns a list of watchlists belonging to a user; via user cookie
// router.GET("/api/watchlists-by-user-id?id={watchlistID", controllers.GetWatchlistsByUserID)
export function fetchWatchlistsAPI () {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    
    return axios.get('http://localhost:8080/api/watchlists-by-user-id', { headers })
        .then(response => {
            return response; // Returning entire response so that component can check status code
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

// 	router.Get("/api/watchlists/movie/{movieID}", controllers.GetWatchlistsByUserIDWithMovieIDCheck)
// GET all watchlists belong to user + watchlist_item count for each + boolean if queried movieID is in the watchlist
// Used to render watchlist belonging to a user on the Movies page.
export function fetchWatchlistsByUserIDWithMovieIDCheck (movieID) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        // Alerts frontend that the user is not logged in. Returns null for component to handle
        return null; 
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    
    const url = `http://localhost:8080/api/watchlists/movie/${movieID}`
    return axios.get(url, { headers })
        .then(response => {
            if (response.status === 204) {
                throw new Error('You haven\'t created any watchlists yet.');
                }
            return response.data; // Returning list of watchlists
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }
    

// 	router.DELETE("/api/watchlist?id={watchlistID}", controllers.DeleteWatchlistByID)
// DELETE watchlist via its id
export function deleteWatchlistAPI (watchlistID) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    
    return axios.delete('http://localhost:8080/api/watchlist', {
        headers, 
        params: {
            id: watchlistID,
        },
        })
        .then(response => {
            return response.data; // Returning message of { 'message': 'Watchlist deleted successfully'}
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

// router.Patch("/api/watchlist-name", controllers.UpdateWatchlistNameByID) // PATCH a watchlist name
// expects "?id={watchlistID}" query param + new name in the json body
export function editWatchlistNameAPI (watchlistID, newWatchlistName) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    const url = 'http://localhost:8080/api/watchlist-name';
    const params = {
        id: watchlistID,
    };
    const data = {
        'name': newWatchlistName,
    };

    return axios.patch(url, data, {headers, params})
        .then(response => {
            return response.data; // Returning { 'name': 'new name here' } to component
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }


// router.Patch("/api/watchlist-description", controllers.UpdateWatchlistDescriptionByID) // PATCH watchlist description
// expects "?id={watchlistID}" query param + new description in the json body"
export function editWatchlistDescriptionAPI (watchlistID, newWatchlistDescription) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    const url = 'http://localhost:8080/api/watchlist-description';
    const params = {
        id: watchlistID,
    };
    const data = {
        'description': newWatchlistDescription,
    };

    return axios.patch(url, data, {headers, params})
        .then(response => {
            return response.data; // Returning { 'description': 'new description here' } to component
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

// router.Post("/api/watchlists", controllers.CreateWatchlist)    // POST a watchlist; user_id retrieved from JWT token
// Watchlist data passed in json body. Returns {message: "Watchlist created successfully!"}.
export function createWatchlistAPI(newWatchlistName, newWatchlistDescription) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };

    const url = 'http://localhost:8080/api/watchlist';

    const data = {
        'name': newWatchlistName,
        'description': newWatchlistDescription,
    };

    return axios.post(url, data, {headers})
        .then(response => {
            return response.data; // Details of created watchlist so component can redirect via id
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

// router.Post("/api/watchlist-item", controllers.CreateWatchlistItemByWatchlistID) // POST create a watchlist item for a specific watchlist
export function addWatchlistItemAPI(watchlistID, movieID) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };

    const url = 'http://localhost:8080/api/watchlist-item';

    const parsedMovieID = parseInt(movieID, 10); // Convert string movieID into an int for backend 
    const data = {
        'watchlist_id': watchlistID,
        'movie_id': parsedMovieID,
        'checkmarked': false, // default mark item as unwatched
    };

    return axios.post(url, data, {headers})
        .then(response => {
            return response; // Holds watchlist_item details
            // Just need to check for successful 200 status & alert user in component
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

// router.Post("/api/watchlist-item-note", controllers.CreateWatchlistItemNote) // POST create watchlist item note for a specific watchlist item
export function createWatchlistItemNoteAPI (watchlistItemId, newItemNote) {
    // Fetch the user's stored JWT token from cookies
    console.log('post api hit')
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    const url = 'http://localhost:8080/api/watchlist-item-note';
    const data = {
        'watchlist_item_id': watchlistItemId,
        'item_notes': newItemNote,
    };

    return axios.post(url, data, {headers})
        .then(response => {
            console.log(response)
            return response; // Returns entire response with headers to front. Data contains user's updated data
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

// router.Patch("/api/watchlist-item-note", controllers.UpdateWatchlistItemNote) // PATCH watchlist_item_note 'item notes'. Passed through JSON body
export function editWatchlistItemNoteAPI (watchlistItemId, editedItemNote) {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
        }

    const headers = {
        Authorization: `Bearer ${token}`,
    };
    const url = 'http://localhost:8080/api/watchlist-item-note';
    const data = {
        'watchlist_item_id': watchlistItemId,
        'item_notes': editedItemNote,
    };

    return axios.patch(url, data, {headers})
        .then(response => {
            console.log(response)
            return response; // Returns entire response with headers to front. Data contains user's updated data
        })
        .catch(error => { // Will catch any error thrown by extractToken
            if (error.response) {
                const errorMessage = error.response.data.message;
                console.error('Error:', errorMessage);
                throw new Error(errorMessage);
            } else { 
                console.error('Network or other error:', error); 
                throw error;
            }
        });
    }

