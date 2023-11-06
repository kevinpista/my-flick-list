   /*
   // Make an API request to your Go backend for user registration
    try {
        const response = await axios.post('/api/user-registration', formData);

        // Check if the registration was successful
        if (response.status === 201) {
        // Registration successful; you can redirect or show a success message.
        console.log('Registration successful');
        }
    } catch (error) {
        // Handle registration error, e.g., display an error message.
        console.error('Registration failed', error);
    }
    };

    */


import axios from 'axios';

export function registerUser(formData) {
    return axios.post('http://localhost:8080/api/user-registration', formData)
    .then(response => {
        return response.data;
    })
    .catch(error => {
        console.log(error);
        throw error; // Rethrow the error to handle in component
    });
}

