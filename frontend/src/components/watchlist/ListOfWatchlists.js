import React, { useState, useEffect} from 'react';
import { useNavigate } from 'react-router-dom';
import { Container, Paper, Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, Typography, Link } from '@mui/material';
import LoadingButton from '@mui/lab/LoadingButton';
import InputAdornment from '@mui/material/InputAdornment';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';
import { fetchWatchlistsAPI, createWatchlistAPI, deleteWatchlistAPI } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';

import ListOfWatchlistsTable from './ListOfWatchlistsTable.js';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../../css/MuiThemeProvider.js';
import { getJwtTokenFromCookies } from '../../utils/authTokenUtils'


// List of Watchlists; states name of watchlist, its description, and how many movies inside
// Accessed via URL of /watchlists

const ListOfWatchlists = () => {
  const jwtToken = getJwtTokenFromCookies();
  const [noWatchlistsFound, setNoWatchlistsFound] = useState(false);
  const navigate = useNavigate();
  const [watchlistData, setWatchlistData] = useState(null); // In JSON object format
  const [error, setError] = useState(null);

  const [isCreateWatchlistDialogOpen, setCreateWatchlistDialogOpen] = useState(false);
  const [newWatchlistName, setNewWatchlistName] = useState('');
  const [newWatchlistDescription, setNewWatchlistDescription] = useState('');

  const [dialogErrorMessage, setDialogErrorMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const fetchData = async () => {
    try {
      const response = await fetchWatchlistsAPI(); // Entire response with headers
      if (response.status === 204) {
        setNoWatchlistsFound(true);
      } else {
        setWatchlistData(response.data);
      }
    } catch (error) {
      if (error.message === errorConstants.ERROR_BAD_REQUEST) {
        console.log('Bad request');
        setError(Error('Bad request. Try again'));
    }
    else if (error.message === errorConstants.TOKEN_EXPIRED) {
        console.error('Token expired!!!');
        setError(Error('Token expired or missing. Redirecting you to login ... '));
        setTimeout(() => {
          navigate('/user-login');
        }, 2500); // Redirect to login
    } else {
        console.error('An unexpected error occured');
        setError(error)
    }
    }
  };

  useEffect(() => {
    // Call fetchData if jwtToken is present
    if (jwtToken) {
      fetchData();
    }
  }, [jwtToken]);

// deleteWatchlistAPI Call
const handleDeleteWatchlist = async (watchlistID) => {
  try {
    const response = await deleteWatchlistAPI(watchlistID);
    if (response) {
      console.log(response.message);

      // Update list of watchlists table
      setWatchlistData((prevItems) => {
        const currentItems = prevItems && prevItems['watchlists']; // Extract the array from the object
        const updatedItems = Array.isArray(currentItems)
          ? currentItems.filter((watchlist) => watchlist.id !== watchlistID) // Re-render all watchlists not equal to the watchlistID that was deleted
          : [];
        return { 'watchlists': updatedItems }; // Maintain JSON object structure
      });
    }

  } catch (error) {
    if (error.message === errorConstants.ERROR_BAD_REQUEST) {
      console.error('Bad request:', error);
    } 
    else if (error.message === errorConstants.TOKEN_EXPIRED) {
      console.error('Token expired!!!');
      setError(Error('Token expired or missing. Redirecting you to login ... '));
      setTimeout(() => {
        navigate('/user-login');
      }, 2500); // Redirect to login
    } else {
      console.error('An unexpected error occured');
    }
  };
};

// Handle Create Watchlist
const handleCreateWatchlistButtonClick = () => {
  setCreateWatchlistDialogOpen(true);
};

const handleCreateWatchlistButtonClose = () => {
  setCreateWatchlistDialogOpen(false);
  setDialogErrorMessage(''); // Clear error message when the dialog is closed
};

// createWatchlistAPI Call
const handleCreateWatchlistDialogSubmit = async () => {
  setLoading(true)
  try {
    const response = await createWatchlistAPI(newWatchlistName, newWatchlistDescription);
    if (response) {

      setTimeout(() => {
        navigate(`/watchlist/${response.id}`)
      }, 2000); // Redirect to new watchlist 

    }
  } catch (error) {
    if (error.message === errorConstants.ERROR_BAD_REQUEST) {
      setLoading(false)
      setDialogErrorMessage('Error: Bad request. Please try again.');
    } else {
      setLoading(false)
      setDialogErrorMessage(`Error: ${error.message}`);
    } 
  }
};

// Renders "Login or Sign Up" pop up if a jwtToken is not found
  if (!jwtToken) {
    return (
      <ThemeProvider theme={muiTheme}>
        <NavBar />
        <Container maxWidth="sm" style={{ marginTop: '50px', textAlign: 'center' }}>
        <Paper elevation={6} style={{ padding: '25px' }}>
          <Typography variant="h6">
            Please sign up or log in to create a watchlist.
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
        </Container>
      </ThemeProvider>
    );
  };

// Renders "Create Watchlist" button pop up if the user does not have any watchlists yet
  if (noWatchlistsFound) {
    return (
      <ThemeProvider theme={muiTheme}>
        <NavBar />
        <Container maxWidth="sm" style={{ marginTop: '50px', textAlign: 'center' }}>
        <Paper elevation={6} style={{ padding: '25px' }}>
          <Typography variant="h6">
            Let's create your first watchlist!
          </Typography>
          <div style ={{ margin: '10px' }}>
            <Button variant="contained" color="primary" size="large" onClick={handleCreateWatchlistButtonClick} style={{ width: '200px', margin: '10px' }}>
              Create Watchlist
            </Button>
    
          </div>
          <Typography variant="h7" >
            You will need a watchlist to add movies.
          </Typography>
        </Paper>
        </Container>

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
          {dialogErrorMessage && (
            <Typography color="error" variant="body2">
              {dialogErrorMessage}
            </Typography>
          )}
        </DialogContent>
        <DialogActions style={{ paddingBottom: '20px', paddingRight: '18px' }}>
          <Button variant="contained" onClick={handleCreateWatchlistButtonClose}>
            Exit</Button>

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
      </ThemeProvider>
    );
  };

  // Main component render
  return (
    <ThemeProvider theme={muiTheme}>
    <React.Fragment>
      <NavBar />
    <Container maxWidth={"xl"} className="watchlist-item-grid-container">
      <Button variant="contained" onClick={handleCreateWatchlistButtonClick}>
        Create a Watchlist
      </Button>
      <h1 className="watchlist-name">Your Watchlists</h1>
      {error ? (
        <h1 className='error'><u>Error:</u> {error.message}</h1>
      ) : (
        watchlistData && (
        <ListOfWatchlistsTable 
          watchlistData={watchlistData}
          onDeleteWatchlist={handleDeleteWatchlist} // onDeleteWatchlist function gets passed to component. When called, it invokes handleDeleteWatchlist
        />
        )
      )}
    </Container>

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
        {dialogErrorMessage && (
          <Typography color="error" variant="body2">
            {dialogErrorMessage}
          </Typography>
        )}
      </DialogContent>
      <DialogActions style={{ paddingBottom: '20px', paddingRight: '18px' }}>
        <Button variant="contained" onClick={handleCreateWatchlistButtonClose}>
          Exit</Button>

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

    </React.Fragment>
    </ThemeProvider>
  );
};

export default ListOfWatchlists;