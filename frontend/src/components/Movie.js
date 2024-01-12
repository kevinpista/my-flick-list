import React, { useState, useEffect} from 'react';
import { Container, Paper, Typography, Button } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import NavBar from './NavBar';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/Movie.css';
import { getMovieDataTMDB } from '../api/movieDataTMDB';
import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../utils/formatUtils';
import { useParams } from 'react-router-dom';
import { fetchWatchlistsAPI, addWatchlistItemAPI } from '../api/watchlistAPI'

import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import Alert from '@mui/material/Alert';
import Snackbar from '@mui/material/Snackbar';

// TODO
// 1. Change AddIcon button and icon to "Added" with CheckMark icon when successfully added to someone's watchlist
// popup menu of which watchlist to add to -- or dropdown to let them pick which watchlist they want to add to.

// 2. Make revenue and budget side by side with nicer icons.

// 3. Move the moive poster image more to the right along with the movie content accordingly-; padding / flex adjustments
// 4. Movie details currently shifts all the way to the right if the content isn't long enough to fill the 2/3 space. Format so that 
// movie details always begins aligned left next to the movie poster regardless of overall content lenght. Can see this difference based on the movie data

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
    const [validMovie, setValidMovie] = useState(null);

    const [error, setError] = useState(null);
    const { movieID } = useParams(); // Extract movieID from the URL params

    const [openDialog, setOpenDialog] = useState(false);
    const [selectedWatchlistID, setSelectedWatchlistID] = useState('');
    const [userWatchlists, setUserWatchlists] = useState(null);

    const [successAlertOpen, setSuccessAlertOpen] = useState(false);
    const [errorAlertOpen, setErrorAlertOpen] = useState(false);
    const [alertMessage, setAlertMessage] = useState('');


    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await getMovieDataTMDB(movieID)
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
                setValidMovie(true);
                
                // Fetch user's watchlist on mount
                const fetchedWatchlists = await fetchWatchlistsAPI()
                setUserWatchlists(fetchedWatchlists)

            } catch (error) {
                setError(error);
            }
        };

        fetchData();
        }, [movieID]);
      
    const moviePosterBaseUrl = "https://image.tmdb.org/t/p/w300_and_h450_bestv2";

    // Handles functions related to when user clicks "Add To Watchlist". 
    const handleOpenDialog = () => {
        setOpenDialog(true);
    };
    const handleCloseDialog = () => {
        setOpenDialog(false);
    };
    const handleWatchlistChange = (event) => {
        setSelectedWatchlistID(event.target.value);
    };
    const handleConfirm = async () => {
    // Send axios request with selectedWatchlistID and movieID
    console.log('user picked:', selectedWatchlistID); // test
    try {
        const response = await addWatchlistItemAPI(selectedWatchlistID, movieID)
        if (response.status === 200) {
            console.log("Movie added successfully to watchlist")
            handleSuccessAlertOpen();
        } else {
            console.error('Request failed with status:', response.status);
            handleErrorAlertOpen(`Request failed with status: ${response.status}`);
        }
    } catch (error) {
        console.log(error)
        handleErrorAlertOpen(`Error: ${error.message}`);
    } finally { // Closes dialog regardless of a successful or failed API request
        handleCloseDialog();
    }
    };

    // Handles alert messages related to when a user submits the watchlist + movie they want to add
    const handleSuccessAlertOpen = () => {
        setAlertMessage('Watchlist Added Successfully!');
        setSuccessAlertOpen(true);
    };
    
    const handleErrorAlertOpen = (errorMessage) => {
        setAlertMessage(errorMessage);
        setErrorAlertOpen(true);
    };
    
    const handleAlertClose = () => {
        setSuccessAlertOpen(false);
        setErrorAlertOpen(false);
        setAlertMessage('');
    };

    // RENDER COMPONENT
    return (
        <React.Fragment>
        <NavBar />
        <Container maxWidth="fluid">
            <Snackbar
                anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
                style={{ top: '50px' }} // Adjust the top position as needed
                open={successAlertOpen || errorAlertOpen}
                autoHideDuration={5000}
                onClose={handleAlertClose}
            >
                <Alert
                    onClose={handleAlertClose}
                    severity={successAlertOpen ? 'success' : 'error'}
                >
                    {alertMessage}
                </Alert>
            </Snackbar>
            {error ? (
                <h1 className='error'><u>Error loading movie:</u> {error.message}</h1>
            ) : ( 
            validMovie && (
            <Paper elevation={3} className="movie-paper">
                <div className="movie-poster">
                    <img className="poster-small" src={`${moviePosterBaseUrl}${moviePosterPath}`} alt="Movie Poster" />
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
                        onClick={handleOpenDialog} 
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
            )
        )}

            <Dialog
                open={openDialog}
                onClose={handleCloseDialog}
            >
                <DialogTitle>Select a Watchlist</DialogTitle>
                <DialogContent>
                    <Select
                    value={selectedWatchlistID}
                    onChange={handleWatchlistChange}
                    fullWidth
                    >
                    {/* Map through user's watchlists and populate the dropdown */}
                        {userWatchlists === null  ? (
                        <MenuItem disabled>You haven't created a Watchlist yet!</MenuItem>
                        ) : (
                        userWatchlists['watchlists'].map((watchlist) => (
                            <MenuItem key={watchlist.id} value={watchlist.id}>
                                {watchlist.name}
                            </MenuItem>
                        ))
                        )}
                    </Select>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseDialog} color="primary">
                    Cancel
                    </Button>
                    <Button onClick={handleConfirm} color="primary" disabled={!selectedWatchlistID}>
                    Confirm
                    </Button>
                </DialogActions>
            </Dialog>
        </Container>
        </React.Fragment>
  );
};
export default MoviePage;
