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
            if (response.status === 204) {
                throw new Error('You haven\'t added any movies to this watchlist yet.');
              }
            return response.data; // Returning watchlist item's movie data
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
export function fetchWatchlists () {
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
export function deleteWatchlist (watchlistID) {
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
            id: watchlistID,
        },
        })
        .then(response => {
            return response.data; // Returning success message "mesage: Watchlist deleted successfully"
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