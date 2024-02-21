import axios from 'axios';

// Fetches the movie search results from TMDB API - must pass in query parameters

const apiUrl = process.env.REACT_APP_API_URL; // Backend API URL loaded via environment variable

export function movieSearchTMDBAPI(query, page = 1) {
    // query argument will be normal spaced out words, encode query for spaces to become "%2B" 
    const encodedQuery = encodeURIComponent(query).replace(/%20/g, '%2B');
    return axios.get(`${apiUrl}/api/tmdb-search?query=${encodedQuery}&page=${page}`)
    .then(response => {
        return response; // Returning entire response so that component can check status code
      })
      .catch(error => {
        if (error.response) {
          const errorMessage = error.response.data.message;
          console.error('Thrown Error:', errorMessage);
          throw new Error(errorMessage);
        } else {
          console.error('Network or other error:', error);
          throw error;
        }
      });
    }

