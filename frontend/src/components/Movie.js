import React, { useState, useEffect} from 'react';
import { Container, Paper, Typography, Button } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import NavBar from './NavBar';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/Movie.css';
import { getMovieDataTMDB } from '../api/movieDataTMDB';
import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../utils/formatUtils';


// TODO
// 1. Change AddIcon button and icon to "Added" with CheckMark icon when successfully added to someone's watchlist
// popup menu of which watchlist to add to -- or dropdown to let them pick which watchlist they want to add to.

// 2. Make revenue and budget side by side with nicer icons.

// 3. Move the moive poster image more to the right along with the movie content accordingly-; padding / flex adjustments
// 4. Movie details currently shifts all the way to the right if the content isn't long enough to fill the 2/3 space. Format so that 
// movie details always begins aligned left next to the movie poster regardless of overall content lenght. Can see this difference based on the movie data

// 5. const movieID needs to be dynamically updated based on what movieID user is sending a request to view for at the certain webpage; 
// likely make it a query url for our frontend

const MoviePage = () => {
    const [moviePosterPath, setMoviePosterPath] = useState('');
    const [movieTitle, setMovieTitle] = useState('');
    const [movieReleaseDate, setMovieReleaseDate] = useState('');
    const [movieGenres, setMovieGenres] = useState([]); // Possibly more than 1 genre
    const [movieRuntime, setMovieRuntime] = useState('');
    const [movieTagline, setMovieTagline] = useState('');
    const [movieOverview, setMovieOverview] = useState('');
    const [movieRevenue, setMovieRevenue] = useState('');
    const [movieBudget, setMovieBudget] = useState('');
    const [movieVoteAverage, setMovieVoteAverage] = useState(0);
    const [movieVoteCount, setMovieVoteCount] = useState(0);

    useEffect(() => {
      // Hardcoding movieID for now
      const movieID = '12445';
  
      // Axios API call
      getMovieDataTMDB(movieID)
        .then(data => {
          // Extract the movie data from the single JSON response
          const moviePosterPathFromTMDBAPI = data.movie.poster_path;
          const movieTitleFromTMDBAPI = data.movie.original_title;
          // Format the release_date data as it is provided as "YYYY-MM-DD"
          const movieReleaseDateFromTMDBAPI = data.movie.release_date;
          const formattedReleaseDate = formatReleaseDate(movieReleaseDateFromTMDBAPI);

          // Format the runtime data as it's provided as "minutes"
          const movieRuntimeFromTMDBAPI = data.movie.runtime;
          const formattedRuntime = formatRuntime(movieRuntimeFromTMDBAPI);

          // Format the vote_average data as it's provided in a long decimal number
          const movieVoteAverageFromTMDBAPI = data.movie.vote_average;
          const formattedVoteAverage = Math.round(movieVoteAverageFromTMDBAPI * 10) / 10; // Round to one decimal place.

          // Format the vote_count data as it's provided as a long integer number
          const movieVoteCountFromTMDBAPI = data.movie.vote_count;
          const formattedVoteCount = formatVoteCount(movieVoteCountFromTMDBAPI);

          const movieGenresFromTMDBAPI = data.movie.genres.map(genre => genre.name); // Possibly more than 1 genre
          const movieTaglineFromTMDBAPI = data.movie.tagline;
          const movieOverviewFromTMDBAPI = data.movie.overview;

          // Format revenue and budget data
          const movieRevenueFromTMDBIAPI = data.movie.revenue;
          const formattedRevenue = formatFinancialData(movieRevenueFromTMDBIAPI);
          const movieBudgetFromTMDBAPI = data.movie.budget;
          const formattedBudget = formatFinancialData(movieBudgetFromTMDBAPI);
          
          setMoviePosterPath(moviePosterPathFromTMDBAPI);
          setMovieTitle(movieTitleFromTMDBAPI);
          setMovieReleaseDate(formattedReleaseDate);
          setMovieRuntime(formattedRuntime);
          setMovieVoteCount(formattedVoteCount);
          setMovieVoteAverage(formattedVoteAverage);
          setMovieGenres(movieGenresFromTMDBAPI);
          setMovieTagline(movieTaglineFromTMDBAPI);
          setMovieOverview(movieOverviewFromTMDBAPI);
          setMovieRevenue(formattedRevenue);
          setMovieBudget(formattedBudget);
        })
        .catch(error => {
          console.log(error);
        });
    }, []);
      
    const moviePosterBaseUrl = "https://image.tmdb.org/t/p/w300_and_h450_bestv2";

    // RENDER COMPONENT
    return (
        <React.Fragment>
        <NavBar />
        <Container maxWidth="fluid">
            <Paper elevation={3} className="movie-paper">

                <div className="movie-poster">
                    <img class="poster-small" src={`${moviePosterBaseUrl}${moviePosterPath}`} alt="Movie Poster" />
                </div>

                <div className="movie-details">
                    <Typography variant="h3" className="movie-title" fontWeight="bold">
                        {movieTitle}
                    </Typography>
                    <div className="movie-description">
                        <Typography variant="body4">
                            {movieReleaseDate} | {movieGenres.join(', ')} | {movieRuntime}              
                        </Typography>
                    </div>

                    <div className="movie-ratings">
                        <Typography variant="body3" >
                            Ratings: {movieVoteAverage} out of 10 | ({movieVoteCount})
                        </Typography>
                    </div>

                    <Typography variant="body4" gutterBottom className="movie-tagline">
                        {movieTagline}
                    </Typography>

                    <Typography variant="h5" className="movie-description">
                        Overview
                    </Typography>
                    
                    <Typography variant="body1" paragraph>
                        {movieOverview}
                    </Typography>
                    
                    <div className="movie-financials">
                        <Typography variant="body1" gutterBottom>
                            Revenue: {movieRevenue} || Budget: {movieBudget}
                        </Typography>
                    </div>
                    
                    <Button 
                        variant="contained"
                        color="primary"
                        size="large" 
                        className="add-to-watchlist-btn"
                        endIcon={<AddIcon />}
                        >
                        ADD TO WATCHLIST
                    </Button>
                </div>

            </Paper>
        </Container>
        </React.Fragment>
  );
};
export default MoviePage;
