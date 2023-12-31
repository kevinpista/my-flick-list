// API requests related to watchlists, their watchlist_items, and item_notes

import axios from 'axios';
import { getJwtTokenFromCookies } from '../utils/authTokenUtils'


// Returns watchlist name, description and its watchlist items
// router.GET("/api/watchlist-items-with-movies, controllers.GetAllWatchlistItemsWithMoviesByWatchListID) 
export function fetchWatchlistAndItems (watchlistID) {
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
export function deleteWatchlistItem (watchlistItemID) {
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


// 	router.DELETE("/api/watchlist?id={watchlistID}", controllers.DeleteWatchlistByID) // DELETE watchlist via its id
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
export function editWatchlistName (watchlistID, newWatchlistName) {
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
export function editWatchlistDescription (watchlistID, newWatchlistDescription) {
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
            return response; // Returning to component so it can access status code 200
            // contains a response.data.message that says "Watchlist created successfully!"
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
