import axios from 'axios';

// Fetches the movie data from TMDB API - must pass in the movie ID which we will assume is provided by the search page
// or other pages which will have the movie ID if it ever is showing any sort of movie data

export function getMovieDataTMDB(movie_id) {
    const movieID = movie_id
    
    return axios.get('http://localhost:8080/api/tmdb-movie?query=' + movieID)
    .then(response => {
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

