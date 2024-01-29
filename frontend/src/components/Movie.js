import React, { useState, useEffect} from 'react';
import { Container, Paper, Typography, Button, InputLabel, Link, TextField } from '@mui/material';
import LoadingButton from '@mui/lab/LoadingButton';
import InputAdornment from '@mui/material/InputAdornment';
import AddIcon from '@mui/icons-material/Add';
import NavBar from './NavBar';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/Movie.css';
import { getMovieDataTMDB } from '../api/movieDataTMDB';
import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../utils/formatUtils';
import { useParams } from 'react-router-dom';
import { fetchWatchlistsByUserIDWithMovieIDCheckAPI, addWatchlistItemAPI } from '../api/watchlistAPI'
import { useNavigate } from 'react-router-dom';
import { createWatchlistAPI } from '../api/watchlistAPI.js'
import * as errorConstants from '../api/errorConstants';

import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import Alert from '@mui/material/Alert';
import Snackbar from '@mui/material/Snackbar';

import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';

// TODO
// 1. Change AddIcon button and icon to "Added" with CheckMark icon when successfully added to someone's watchlist
// popup menu of which watchlist to add to -- or dropdown to let them pick which watchlist they want to add to.

// 2. Make revenue and budget side by side with nicer icons.

// 3. Move the moive poster image more to the right along with the movie content accordingly-; padding / flex adjustments
// 4. Movie details currently shifts all the way to the right if the content isn't long enough to fill the 2/3 space. Format so that 
// movie details always begins aligned left next to the movie poster regardless of overall content lenght. Can see this difference based on the movie data

// 5. TEST 'setLongestWatchlistNameLength(fetchedWatchlists['watchlists']... ' line in useEffect
// when a new user who has no watchlists on file and see if component handles it correctly

const MoviePage = () => {
    // Movie data variables
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
    const [movieError, setMovieError] = useState(null);
    const { movieID } = useParams(); // Extract movieID from the URL params
    const navigate = useNavigate();

    // Add to Watchlist Dropdown List variables
    const [selectedWatchlistID, setSelectedWatchlistID] = useState('');
    const [selectedWatchlist, setSelectedWatchlist] = useState('placeholder');
    const [userWatchlists, setUserWatchlists] = useState(null);
    const [openWatchlistDropdownDialog, setOpenWatchlistDropdownDialog] = useState(false);
    const [longestWatchlistNameLength, setLongestWatchlistNameLength] = useState(0);

    // Alert + Snackbar variables for Add to Watchlist + Create Watchlist actions
    const [successAlertOpen, setSuccessAlertOpen] = useState(false);
    const [errorAlertOpen, setErrorAlertOpen] = useState(false);
    const [alertMessage, setAlertMessage] = useState('');

    // Create Watchlist variables
    const [isCreateWatchlistDialogOpen, setCreateWatchlistDialogOpen] = useState(false);
    const [newWatchlistName, setNewWatchlistName] = useState('');
    const [newWatchlistDescription, setNewWatchlistDescription] = useState('');
    const [createWatchlistDialogErrorMessage, setCreateWatchlistDialogErrorMessage] = useState('');
    const [loading, setLoading] = useState(false); // Loading state for Creatch Watchlist button

    // Handle Create Watchlist
    const handleCreateWatchlistButtonClick = () => {
        setCreateWatchlistDialogOpen(true);
    };
    const handleCreateWatchlistButtonClose = () => {
        setCreateWatchlistDialogOpen(false);
        setCreateWatchlistDialogErrorMessage(''); // Clear error message when the dialog is closed
    };

    // createWatchlistAPI Call
    const handleCreateWatchlistDialogSubmit = async () => {
        setLoading(true)
        try {
        const response = await createWatchlistAPI(newWatchlistName, newWatchlistDescription);
        if (response) {
            const fetchedWatchlists = await fetchWatchlistsByUserIDWithMovieIDCheckAPI(movieID);           setCreateWatchlistDialogOpen(false);
            setUserWatchlists(fetchedWatchlists);
            // Calculate the longest watchlist name for proper styling in watchlist dropdown menu dialog
            setLongestWatchlistNameLength(fetchedWatchlists['watchlists'].reduce((max, watchlist) => {
                return Math.max(max, watchlist.name.length, 30);
              }, 0));
            setCreateWatchlistDialogOpen(false);
            handleWatchlistCreateSuccessAlertOpen();
        }
        } catch (error) {
            if (error.message === errorConstants.ERROR_BAD_REQUEST) {
                setLoading(false)
                setCreateWatchlistDialogErrorMessage('Error: Bad request. Please try again.');
            } else {
                setLoading(false)
                setCreateWatchlistDialogErrorMessage(`Error: ${error.message}`);
            } 
        }
    };

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
                const fetchedWatchlists = await fetchWatchlistsByUserIDWithMovieIDCheckAPI(movieID)
                if (fetchedWatchlists === null) { // User is not logged so API call will return null
                    setUserWatchlists(null)
                } else if (fetchedWatchlists.status === 204) {
                    setUserWatchlists({"watchlists": []}) // Set empty to later conditionally render a "Create Watchlist" dialog
                } else {
                    setUserWatchlists(fetchedWatchlists)
                    // Calculate the longest watchlist name for proper styling in watchlist dropdown menu dialog
                    setLongestWatchlistNameLength(fetchedWatchlists['watchlists'].reduce((max, watchlist) => {
                        return Math.max(max, watchlist.name.length, 30);
                      }, 0));
                }

            } catch (error) {
                setMovieError(error); // Note, may be an error related to fetchWatchlistsByUserIDWithMovieIDCheckAPI()
                // But will display it with general error message of "Error loading movie"
            }
        };

        fetchData();
        }, [movieID]);
      
    const moviePosterBaseUrl = "https://image.tmdb.org/t/p/w300_and_h450_bestv2";

    // Handles functions related to when user clicks "Add To Watchlist". 
    const handleOpenWatchlistDropdownDialog = () => {
        setOpenWatchlistDropdownDialog(true);
    };
    const handleCloseWatchlistDropdownDialog = () => {
        setOpenWatchlistDropdownDialog(false);
        setSelectedWatchlistID('');
        setSelectedWatchlist('placeholder');
    };
    const handleWatchlistChange = (event) => {
        setSelectedWatchlist(event.target.value)
        setSelectedWatchlistID(event.target.value.id);
    };
    const handleAddMovieToWatchlist = async () => {
        try {
            const response = await addWatchlistItemAPI(selectedWatchlistID, movieID)
            if (response.status === 200) {
                const fetchedWatchlists = await fetchWatchlistsByUserIDWithMovieIDCheckAPI(movieID)
                setUserWatchlists(fetchedWatchlists)
                handleMovieAddSuccessAlertOpen();
            } else {
                console.error('Request failed with status:', response.status);
                handleMovieAddErrorAlertOpen(`Request failed with status: ${response.status}`);
            }
        } catch (error) {
            handleMovieAddErrorAlertOpen(`Error: ${error.message}`);
        } finally { // Closes dialog regardless of a successful or failed API request
            handleCloseWatchlistDropdownDialog();
        }
    };

    // Handles alert messages related to user adding a movie or creating a new watchlist
    const handleMovieAddSuccessAlertOpen = () => {
        setAlertMessage('Movie added to Watchlist successfully!');
        setSuccessAlertOpen(true);
    };
    const handleMovieAddErrorAlertOpen = (errorMessage) => {
        setAlertMessage(errorMessage);
        setErrorAlertOpen(true);
    };
    
    const handleAlertClose = () => {
        setSuccessAlertOpen(false);
        setErrorAlertOpen(false);
        setAlertMessage('');
    };
    // Message related to creating a watchlist - error message handled inside dialog form via 'createWatchlistDialogErrorMessage'
    const handleWatchlistCreateSuccessAlertOpen = () => {
        setAlertMessage('Watchlist created sucessfully!');
        setSuccessAlertOpen(true);
    };

    // RENDER COMPONENT
    return (
        <ThemeProvider theme={muiTheme}>
        <React.Fragment>
        <NavBar />
        <Container maxWidth="fluid">
            <Snackbar
                anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
                style={{ top: '50px' }}
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
            {movieError ? (
                <h1 className='error'><u>Error loading movie:</u> {movieError.message}</h1>
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
                        onClick={handleOpenWatchlistDropdownDialog} 
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
        {/* Watchlist dropdown menu */}
        {/* 2 Edge Cases - User Not Logged In (null) or Logged In But No Watchlists (.length === 0) */}
            {userWatchlists === null ? (
                <Dialog
                    open={openWatchlistDropdownDialog}
                    onClose={handleCloseWatchlistDropdownDialog}
                    style={{ textAlign: 'center' }}
                >
                    <Paper elevation={6} style={{ padding: '25px 70px' }}>
                    <Typography variant="h6">
                        Please sign up or log in to add to a watchlist.
                    </Typography>
                    <div style ={{ margin: '10px' }}>
                        <Button variant="contained" color="primary" onClick={() => navigate('/user-login')} style={{ margin: '10px' }}>
                            Log In
                        </Button>
                        <Button variant="outlined" color="secondary" onClick={() => navigate('/user-registration')} style={{ margin: '10px' }}>
                            Sign Up
                        </Button>
                    </div>
                    <Typography variant="h7">
                    Use a{' '}
                    <Link href="/user-login" underline="always">
                        demo account
                    </Link>
                    {' '} instead.
                    </Typography>
                    </Paper>
                </Dialog>
            ) : userWatchlists['watchlists'].length === 0 ? (
                <>
                <Dialog
                    open={openWatchlistDropdownDialog}
                    onClose={handleCloseWatchlistDropdownDialog}
                    style={{ textAlign: 'center' }}
                >
                    <Paper elevation={6} style={{ padding: '25px 70px' }}>
                    <Typography variant="h6">
                        Let's Create Your First Watchlist
                    </Typography>
                    <div style ={{ margin: '10px' }}>
                    <Button variant="contained" color="primary" size="large" onClick={handleCreateWatchlistButtonClick} style={{ width: '200px', margin: '10px' }}>
                        Create Watchlist
                    </Button>
                    </div>
                    <Typography variant="h7">
                        A watchlist is required to add a movie.
                    </Typography>
                    </Paper>
                </Dialog>
                {/* Modal creating a watchlist */}
                <Dialog
                    open={isCreateWatchlistDialogOpen}
                    onClose={handleCreateWatchlistButtonClose}
                    maxWidth="md"
                    fullWidth={true}
                >
                <DialogTitle><b>Create a New Watchlist</b></DialogTitle>
                <DialogContent>
                    <TextField
                        autoFocus
                        id="watchlist-name"
                        label="Watchlist Name"
                        value={newWatchlistName}
                        onChange={(e) => setNewWatchlistName(e.target.value)}
                        multiline
                        fullWidth
                        margin="dense"
                        variant="standard"
                        // Display character limit and changes text to red if user goes over limit
                        InputProps={{
                            endAdornment: (
                            <InputAdornment position="end">
                            <span style={{ color: newWatchlistName.length > 60 ? 'red' : 'inherit' }}>
                                {newWatchlistName.length}/{60}
                            </span>
                            </InputAdornment>
                            ),
                        }}
                    />
                    <TextField
                        autoFocus
                        id="watchlist-description"
                        label="Watchlist Description"
                        value={newWatchlistDescription}
                        onChange={(e) => setNewWatchlistDescription(e.target.value)}
                        multiline
                        fullWidth
                        margin="dense"
                        variant="standard"
                        // Display character limit and changes text to red if user goes over limit
                        InputProps={{
                            endAdornment: (
                            <InputAdornment position="end">
                            <span style={{ color: newWatchlistDescription.length > 500 ? 'red' : 'inherit' }}>
                                {newWatchlistDescription.length}/{500}
                            </span>
                            </InputAdornment>
                            ),
                        }}
                    />
                    {createWatchlistDialogErrorMessage && (
                    <Typography color="error" variant="body2">
                        {createWatchlistDialogErrorMessage}
                    </Typography>
                    )}
                </DialogContent>
                <DialogActions style={{ paddingBottom: '20px', paddingRight: '18px' }}>
                    <Button variant="contained" onClick={handleCreateWatchlistButtonClose}>
                        Exit
                    </Button>

                    <LoadingButton 
                        variant="contained"
                        loading={loading}
                        onClick={handleCreateWatchlistDialogSubmit}
                        disabled={
                        newWatchlistName.length > 60 || // Character limit for watchlist name
                        newWatchlistDescription.length > 500 // Character limit for watchlist description
                        }
                    >
                        Create
                    </LoadingButton>

                </DialogActions>
                </Dialog>
                </>
            ) : (
            <Dialog
                open={openWatchlistDropdownDialog}
                onClose={handleCloseWatchlistDropdownDialog}
            >
                <DialogTitle>Select a Watchlist</DialogTitle>
                <DialogContent>
                <InputLabel id="watchlist-placeholder">Watchlist</InputLabel>
                    {/* Need renderValue prop to correctly show selected item on one line due to using 2 divs for MenuItem */}
                    <Select
                        id="select-watchlist"
                        value={selectedWatchlist}
                        onChange={handleWatchlistChange}
                        fullWidth
                        renderValue={(selectedValue) => (
                            <div>
                            {selectedValue === 'placeholder' ? (
                                <InputLabel style={{ width: `${longestWatchlistNameLength + 10}ch` }}>
                                    Select a watchlist
                                </InputLabel>
                            ) : (
                                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                                    <div style={{ flex: '1', width: `${longestWatchlistNameLength + 2}ch` }}>
                                        {selectedValue.name}
                                    </div>
                                        <div style={{ textAlign: 'right', paddingLeft: '2px' }}>
                                        {selectedValue.contains_queried_movie ? '[Already in Watchlist]' : `[${selectedValue.watchlist_item_count} movies]`}
                                    </div>
                                </div>
                            )}
                            </div>
                        )}
                    >
                    {/* Map through user's watchlists and populate the dropdown */}
                        {userWatchlists === null  ? (
                        <MenuItem disabled>You haven't created a Watchlist yet!</MenuItem>
                        ) : (
                        userWatchlists['watchlists'].map((watchlist) => (
                            <MenuItem 
                                key={watchlist.id}
                                value={watchlist}
                                disabled={watchlist.contains_queried_movie}
                                style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}
                            >
                                <div style={{ flex: '1', width: `${longestWatchlistNameLength + 2}ch` }}>
                                    {watchlist.name}
                                </div>
                                <div style={{ textAlign: 'right', paddingLeft: '10px' }}> 
                                    {watchlist.contains_queried_movie ? '[Already in Watchlist]' : `[${watchlist.watchlist_item_count} movies]`}
                                </div>
                            </MenuItem>
                        ))
                        )}
                    </Select>
                </DialogContent>
                <DialogActions style={{ paddingBottom: '20px', paddingRight: '18px' }}>
                    <Button variant="contained" onClick={handleCloseWatchlistDropdownDialog} color="primary">
                        Cancel
                    </Button>
                    <Button variant="contained" onClick={handleAddMovieToWatchlist} color="primary" disabled={!selectedWatchlistID || selectedWatchlistID === 'placeholder'}>
                        Add Movie
                    </Button>
                </DialogActions>
            </Dialog>
            )}
        </Container>
        </React.Fragment>
        </ThemeProvider>
  );
};
export default MoviePage;
