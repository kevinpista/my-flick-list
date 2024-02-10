import React, { useState, useEffect} from 'react';
import { Paper, Typography, Button, InputLabel, Link, TextField } from '@mui/material';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import LoadingButton from '@mui/lab/LoadingButton';
import InputAdornment from '@mui/material/InputAdornment';
import AddIcon from '@mui/icons-material/Add';
import NavBar from './NavBar';
import '../css/Movie.css';
import { getMovieDataTMDB, getMovieTrailerTMDB } from '../api/movieDataTMDB.js';
import { formatReleaseDate, formatReleaseYear, formatRuntime, formatVoteCount, formatFinancialData } from '../utils/formatUtils';
import { useParams } from 'react-router-dom';
import { fetchWatchlistsByUserIDWithMovieIDCheckAPI, addWatchlistItemAPI } from '../api/watchlistAPI'
import YouTubeModal from './YouTubeModal.js'
import { useNavigate } from 'react-router-dom';
import { createWatchlistAPI } from '../api/watchlistAPI.js'
import * as errorConstants from '../api/errorConstants';
import useMediaQuery from '@mui/material/useMediaQuery';

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
import no_image_placeholder from '../static/no_image_placeholder.jpg';

// TODO

// 1. Make revenue and budget side by side with nicer icons.


const MoviePage = () => {
    // Movie data variables
    const [moviePosterPath, setMoviePosterPath] = useState('');
    const [movieBackdropPath, setMovieBackdropPath] = useState('');
    const [movieTitle, setMovieTitle] = useState('');
    const [movieReleaseDate, setMovieReleaseDate] = useState('');
    const [movieReleaseYear, setMovieReleaseYear] = useState(null);
    const [movieGenres, setMovieGenres] = useState([]); // Possibly more than 1 genre
    const [movieRuntime, setMovieRuntime] = useState('');
    const [movieTagline, setMovieTagline] = useState('');
    const [movieOverview, setMovieOverview] = useState('');
    const [movieRevenue, setMovieRevenue] = useState('');
    const [movieBudget, setMovieBudget] = useState('');
    const [movieVoteAverage, setMovieVoteAverage] = useState(0);
    const [progressVoteAverage, setProgressVoteAverage] = useState(null);
    const [movieVoteCount, setMovieVoteCount] = useState(0);
    const [validMovie, setValidMovie] = useState(null);
    const [movieError, setMovieError] = useState(null);
    const { movieID } = useParams(); // Extract movieID from the URL params
    const [movieTrailerYouTubeID, setMovieTrailerYouTubeID] = useState(null);
    const [isYouTubeModalOpen, setIsYouTubeModalOpen] = useState(false);
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
                console.log(data)
                // Extract the movie data from the single JSON response
                const moviePosterPathFromTMDBAPI = data.movie.poster_path;
                const movieBackdropPathFromTMDBAPI = data.movie.backdrop_path;
                const movieTitleFromTMDBAPI = data.movie.original_title;
                // Format the release_date data as it is provided as "YYYY-MM-DD"
                const movieReleaseDateFromTMDBAPI = data.movie.release_date;
                const formattedReleaseDate = formatReleaseDate(movieReleaseDateFromTMDBAPI);
                const formattedReleaseYear = formatReleaseYear(movieReleaseDateFromTMDBAPI);

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
                setMovieBackdropPath(movieBackdropPathFromTMDBAPI);
                setMovieTitle(movieTitleFromTMDBAPI);
                setMovieReleaseDate(formattedReleaseDate);
                setMovieReleaseYear(formattedReleaseYear);
                setMovieRuntime(formattedRuntime);
                setMovieVoteCount(formattedVoteCount);
                setMovieVoteAverage(formattedVoteAverage);
                setProgressVoteAverage(Math.round(formattedVoteAverage * 10));
                setMovieGenres(movieGenresFromTMDBAPI);
                setMovieTagline(movieTaglineFromTMDBAPI);
                setMovieOverview(movieOverviewFromTMDBAPI);
                setMovieRevenue(formattedRevenue);
                setMovieBudget(formattedBudget);
                setValidMovie(true);
                
                // Fetch movie's trailer video
                const trailerData = await getMovieTrailerTMDB(movieID);
                setMovieTrailerYouTubeID(trailerData.youtube_video_id)
                console.log(trailerData)

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

    function getProgressColor(value){
        if (value <= 40) return 'error.main';
        if (value <= 60) return 'warning.main';
        if (value <= 78) return 'yellow';
        if (value <= 100) return 'success.main';
        return 'info.main'; // Fallback to primary color.
    }

    function CircularProgressWithLabel({ progressVoteAverage, movieVoteAverage }) {
        const color = getProgressColor(progressVoteAverage);
        const largeScreen = useMediaQuery('(min-width:1280px)');
        const size = largeScreen ? 70 : 50; // Size based on screen width
        return (
            <Box sx={{ 
                position: 'relative', 
                display: 'inline-flex', 
                borderRadius: '50%', 
                backgroundColor: '#081c22', 
                padding: '4px',
                width: `${size + 8}px`, // Need this exact padding for size to have background centered with bar
                height: `${size + 8}px`
            }}>

              <CircularProgress variant="determinate" value={progressVoteAverage} size={size} sx={{ color }} />
              <Box
                sx={{
                  top: 0,
                  left: 0,
                  bottom: 0,
                  right: 0,
                  position: 'absolute',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  textAlign: 'center',
                }}
              >
                <div className="progress-label">
                    {`${(progressVoteAverage)}%`}
                </div>
              </Box>
            </Box>
          );
    }

    const moviePosterBaseUrl = "https://image.tmdb.org/t/p/w600_and_h900_bestv2";
    const movieBackdropBaseUrl = "https://image.tmdb.org/t/p/w1920_and_h800_multi_faces";
    const finalMoviePosterUrl = moviePosterPath === "" ? no_image_placeholder : moviePosterBaseUrl + "" + moviePosterPath;


    const handleOnClickPlayTrailer = () => setIsYouTubeModalOpen(true);

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
                <section className='background-header'
                style={{
                    backgroundImage: `url(${movieBackdropBaseUrl}${movieBackdropPath})`, 
                    borderBottom: '1px solid var(--primaryColor)',
                    backgroundPosition: 'left calc((50vw - 170px) - 340px) top',
                    backgroundSize: 'cover',
                    backgroundRepeat: 'no-repeat'
                }}
                >
             <section className='background-overlay'
                  style={{backgroundImage: `linear-gradient(to right, rgba(10.5, 31.5, 31.5, 1) calc((50vw - 170px) - 340px), rgba(10.5, 31.5, 31.5, 0.84) 50%, rgba(10.5, 31.5, 31.5, 0.84) 100%)`}}
             >
                <div className="movie-content">

                    <div className="movie-poster">
                        <img className="poster-small" src={finalMoviePosterUrl} alt="Movie Poster" />
                    </div>

                    <div className="movie-details" >
                        <h2 className="movie-title">
                            {movieTitle} 
                            {movieReleaseYear && <span className="release-year">({movieReleaseYear})</span>}
                        </h2>

                        <div className="movie-description">
                        <Typography variant="body4">
                            <span className="release">{movieReleaseDate}</span>
                            <span className="genres">{movieGenres.join(', ')}</span> 
                            <span className="runtime">{movieRuntime}</span>            
                        </Typography>
                        </div>

                        <div className="movie-ratings">
                            <Typography variant="body3" >
                                Ratings: {movieVoteAverage} out of 10 | ({movieVoteCount})
                            </Typography>
                        </div>

                        <div className="movie-rating-progress-circle">
                            <CircularProgressWithLabel 
                            progressVoteAverage={progressVoteAverage} 
                            movieVoteAverage={movieVoteAverage}                      
                            />
                        </div>

                        { movieTrailerYouTubeID && (
                            <button onClick={handleOnClickPlayTrailer}>
                                Play Trailer
                            </button>
                        )}
                        { isYouTubeModalOpen && <YouTubeModal isOpen={isYouTubeModalOpen} setIsOpen={setIsYouTubeModalOpen} videoId={movieTrailerYouTubeID} /> }

                        <Typography variant="body4" gutterBottom className="movie-tagline">
                            {movieTagline}
                        </Typography>

                        <h5 className="overview">
                            Overview
                        </h5>
                        
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
                            sx={{
                                boxShadow: `0px 0px 2px rgba(255, 255, 255, 0.2),
                                0px 0px 6px rgba(255, 255, 255, 0.1),
                                0px 0px 10px rgba(255, 255, 255, 0.05)`,   

                                border: `1px solid rgba(255, 255, 255, 0.25)`,
                                // boxShadow: `-2px 4px 6px rgba(255, 255, 255, 0.25)`, alternative white drop shadow
                                }}
                            >
                            ADD TO WATCHLIST
                        </Button>
                    </div>
                </div>
            </section>
            </section>
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
        </React.Fragment>
        </ThemeProvider>
  );
};
export default MoviePage;
