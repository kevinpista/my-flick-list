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
        return response.data;
      })
      .catch(error => {
        console.error(error);
        throw error;
      });
  }


    