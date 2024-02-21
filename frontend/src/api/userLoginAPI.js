// API request to log in the user

import axios from 'axios';
import { extractTokenFromHeader, setTokenInCookie } from '../utils/authTokenUtils'

const apiUrl = process.env.REACT_APP_API_URL; // Backend API URL loaded via environment variable

export function loginUser(formData) {
return axios.post(`${apiUrl}/api/user-login`, formData)
    .then(response => {
        // Set JWT token in user's cookies
        const token = extractTokenFromHeader(response);
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