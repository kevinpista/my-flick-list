// API requests related to watchlists, their watchlist_items, and item_notes

import axios from 'axios';
import { getJwtTokenFromCookies } from '../utils/authTokenUtils'


// Returns all watch items and its full movie data
// router.GET("/api/watchlist-items-with-movies, controllers.GetAllWatchlistItemsWithMoviesByWatchListID) 
export function fetchWatchlistItems (watchlistID) {
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

/*
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

*/