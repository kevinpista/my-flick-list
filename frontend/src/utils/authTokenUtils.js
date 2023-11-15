// Helper functions to help manage user JWT tokens
import Cookies from 'js-cookie';

// Extract user's new JWT token from backend header
export function extractToken(response) {
    if (response.headers && response.headers.authorization) {
      return response.headers.authorization.split(' ')[1]; // Split because token set like 'Bearer <tokenString>'
    }
    throw new Error('Token not found in response headers.'); // Error won't be caught here. Will be caught in the
    // api.js function that calls on this function and will handle the error
  }
  
// Set the JWT token in user's cookies
const COOKIE_NAME = 'jwtToken'; // Key name
export function setTokenInCookie(token) {
  Cookies.set(COOKIE_NAME, token);
}