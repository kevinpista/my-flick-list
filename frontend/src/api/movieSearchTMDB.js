import axios from 'axios';

// Fetches the movie search results from TMDB API - must pass in query parameters

export function movieSearchTMDB(query) {
    // query argument will be normal spaced out words, encode query for spaces to become "%2B" 
    const encodedQuery = encodeURIComponent(query).replace(/%20/g, '%2B');
    
    return axios.get('http://localhost:8080/api/tmdb-search?query=' + encodedQuery)
    .then(response => {
        // TESTING
        console.log(response.data)
        if (response.status === 204) {
          throw new Error('No results found');
        }
        return response.data;
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

