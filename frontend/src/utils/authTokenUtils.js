// Helper functions to help manage user JWT tokens
import Cookies from 'js-cookie';

// Extract user's new JWT token from backend header
export function extractTokenFromHeader(response) {
    if (response.headers && response.headers.authorization) {
      return response.headers.authorization.split(' ')[1]; // Split because token set like 'Bearer <tokenString>'
    }
    throw new Error('Token not found in response headers.'); // Error won't be caught here. Will be caught in the
    // api.js function that calls on this function and will handle the error
  }
  
// Set the JWT token in user's cookies
const COOKIE_NAME = 'jwtToken'; // Key name

export function setTokenInCookie(token) {
  const expires = new Date(); // Initialize the expires variable
  expires.setDate(expires.getDate() + 14);
  Cookies.set(COOKIE_NAME, token, { expires, secure: true});
}

// Removes JWT Token from cookies
export function removeTokenFromCookie() {
  Cookies.remove(COOKIE_NAME);
}

// Example getStoredToken function
export function getJwtTokenFromCookies() {
  return Cookies.get('jwtToken');
}

// Redirects user to the login page if backend verifies their JWT token is expired
export const handleTokenExpiration = () => {
  console.log('Token expired. Redirecting to login page.');

  // TO-DO double check redirect is correct
  window.location.href = '/user-login';
};

// Parses a user's JWT cookie token. Extracts and returns the user_id from the token
export const fetchUserIdFromToken = (token) => {
  try {
    const [, tokenPayload] = token.split('.'); // Putting 1 ',' means we are ignoring 1st split element and 
    // accessing the 2nd element in the split and storing it in variable 'tokenPayload'
    const decodedPayload = JSON.parse(atob(tokenPayload)); // Decode and access the user_id

    if (!decodedPayload || !decodedPayload.user_id) {
      throw new Error('Invalid token: Unable to access user_id');
    }
    return decodedPayload.user_id;

  } catch (error) { // catches new Error created above to propogated to function call
    throw new Error('Error decoding token: ' + error.message);
  }
    /*
    The JWT token string itself is separated into 3 parts by 2 '.'
    const parts = jwtToken.split('.');
    console.log(parts[0]); // Header
    console.log(parts[1]); // Payload -- contains UserId: user_id as assigned in backend token_service.go claims struct
    console.log(parts[2]); // Signature
  */

}