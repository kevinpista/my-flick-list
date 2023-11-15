// API request to log in the user

import axios from 'axios';
import { extractToken, setTokenInCookie } from '../utils/authTokenUtils'

export function loginUser(formData) {
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

Sucessful response returns to the UserLogin.js component the below.

You individually access the token like "reponse.token"
Or you can destructure the key:val with "const { token, userData } = response;"
Then the variable "token" will be the string and "data.id" will be the id 

{
  token: 'your_jwt_token',
  data: {
    "id": "1e1d19a4-7b2a-4ffc-80c3-63b5b17f0ec3",
    "name": "Heyo There",
    "email": "heyo@yahoo.com",
    "created_at": "2023-11-09T15:29:47.918007-08:00",
    "updated_at": "2023-11-09T15:29:47.918007-08:00"
  }
}


*/