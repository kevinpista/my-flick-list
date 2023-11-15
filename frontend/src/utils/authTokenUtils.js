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
  expires.setDate(expires.getDate() + 1); // Expiration of 1 for now
  Cookies.set(COOKIE_NAME, token, { expires, secure: true});
}

// Redirects user to the login page if backend verifies their JWT token is expired
export const handleTokenExpiration = () => {
  console.log('Token expired. Redirecting to login page.');

  // TO-DO double check redirect is correct
  window.location.href = '/user-login';
};