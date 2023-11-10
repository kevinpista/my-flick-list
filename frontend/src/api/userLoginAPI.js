// API request to log in the user

import axios from 'axios';
export function loginUser(formData) {

return axios.post('http://localhost:8080/api/user-login', formData)
    .then(response => {
        return response.data; 
    })
    .catch(error => {
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

