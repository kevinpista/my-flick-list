import React, { useState, useEffect} from 'react';
import { Container, Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, Typography } from '@mui/material';
import InputAdornment from '@mui/material/InputAdornment';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';
import { fetchWatchlistsAPI, createWatchlistAPI, deleteWatchlistAPI } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';

import ListOfWatchlistsTable from './ListOfWatchlistsTable.js';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../../css/MuiThemeProvider.js';

// List of Watchlists; states name of watchlist, its description, and how many movies inside
// Accessed via URL of /watchlists

const ListOfWatchlists = () => {
  const [watchlistData, setWatchlistData] = useState(null); // In JSON object format
  const [error, setError] = useState(null);

  const [isCreateWatchlistDialogOpen, setCreateWatchlistDialogOpen] = useState(false);
  const [newWatchlistName, setNewWatchlistName] = useState('');
  const [newWatchlistDescription, setNewWatchlistDescription] = useState('');

  const [dialogErrorMessage, setDialogErrorMessage] = useState('');

  const fetchData = async () => {
    try {
      const response = await fetchWatchlistsAPI();
      setWatchlistData(response);
    } catch (error) {
      setError(error);
      if (error.message === errorConstants.ERROR_BAD_REQUEST) {
        console.log('Bad request');
    } else {
      console.log('Unexpected error occured');
    }
  }
};

  useEffect(() => {
    // Call fetchData when component mounts
    fetchData();
}, []);

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
  try {
    const response = await createWatchlistAPI(newWatchlistName, newWatchlistDescription);
    if (response) {
      // Will consider redirecting user to newly created watchlist. 
      // For now, will simply refetch watchlist data for user

      // Clear the TextField inputs
      setNewWatchlistName('');
      setNewWatchlistDescription('');

      setCreateWatchlistDialogOpen(false);
      // Trigger fetchData after successful creation
      fetchData();
    }
  } catch (error) {
    if (error.message === errorConstants.ERROR_BAD_REQUEST) {
      setDialogErrorMessage('Error: Bad request. Please try again.');
    } else {
      setDialogErrorMessage(`Error: ${error.message}`);
    }
  };
};

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
      <DialogActions>
        <Button variant="contained" onClick={handleCreateWatchlistButtonClose}>
          Exit</Button>

        <Button 
        variant="contained"
        onClick={handleCreateWatchlistDialogSubmit}
        disabled={
          newWatchlistName.length > 60 || // Character limit for watchlist name
          newWatchlistDescription.length > 500 // Character limit for watchlist description
        }
        >
          Create
        </Button>

      </DialogActions>
    </Dialog>

    </React.Fragment>
    </ThemeProvider>
  );
};

export default ListOfWatchlists;