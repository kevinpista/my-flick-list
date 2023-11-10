// API requiest to register the user

import axios from 'axios';
export function registerUser(formData) {

return axios.post('http://localhost:8080/api/user-registration', formData)
    .then(response => { // this block of code gets executed of status code is a 2xx code
        return response.data; // the .data is the actual JSON body of keys:values data created by my backend Model structs
    })
    .catch(error => { // this block of code gets executed if status code is a 4xx or 5xx code
        if (error.response) { // See brief-notes-to-self below to refresh what this part is
            const errorMessage = error.response.data.message;
            console.error('Registration error:', errorMessage);
            throw new Error(errorMessage); // throwing this 'new' error object and passing in our errorMessage string
            // this is a new error object which the errorMessage. to access the message in the component, do "error.message" 
            // and this is assuming the catch lists its variable name as "error" like in UserRegistration.js
        } else { // error detected, but no response body received so likely a network error
            console.error('Network or other error:', error); 
            throw error;
        }
    });
}


/*
Note the json body of an error sent by the backend looks like

{
    "error": true,
    "message": "name cannot be empty or contain only whitespace",
    "data.omitresponse": null
}

To access the message from the JSON body, you would need to catch the error (so a 4xx or 5xx code), 
then access it with "error.response.data.message"
The .response has all info such as statude code, headers, and data.
The .data has the actual JSON body above. Then you can access the body with .message or .error etc.

*/
