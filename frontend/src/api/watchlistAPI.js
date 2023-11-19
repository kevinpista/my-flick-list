// API requests related to watchlists, their watchlist_items, and item_notes

import axios from 'axios';
import { setTokenInCookie, fetchUserIdFromToken } from '../utils/authTokenUtils'

// Creates a watchlist for a user. User must be logged in
// 	router.Post("/api/watchlists", controllers.CreateWatchlist)    // POST a watchlist
export function createWatchlist (formData) {
return axios.post('http://localhost:8080/api/user-login', formData)
    .then(response => {
        // Set JWT token in user's cookies
        const token = extractToken(response);
        setTokenInCookie(token);
        return response.data; // Returning both the JWT and userData 
    })
    .catch(error => { // Will catch any error thrown by extractToken
        if (error.response) {
            const errorMessage = error.response.data.message;
            console.error('Login error:', errorMessage);
            throw new Error(errorMessage);
        } else { 
            console.error('Network or other error:', error); 
            throw error;
        }
    });
}
/*
import axios from 'axios';
import { getJwtTokenFromCookies } from '../utils/authTokenUtils';

export function getUsers() {
    // Fetch the user's stored JWT token from cookies
    const token = getJwtTokenFromCookies();
    if (!token) {
        console.error('Token not available or expired');
        // For now, will use a Promise.reject method instead of redirect
        return Promise.reject('Token not available or expired');
      }

    const headers= {
        Authorization: `Bearer ${token}`,
    };
    
    return axios.get('http://localhost:8080/api/users', {headers})
    .then(response => {
        const [, payloadBase64] = token.split('.');
        const decodedPayload = JSON.parse(atob(payloadBase64));
        console.log(decodedPayload)
        console.log(decodedPayload.user_id)
        return response.data;
      })
      .catch(error => {
        console.error(error);
        throw error;
      });
  }

*/