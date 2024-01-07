import React, { useState, useEffect} from 'react';
import { Container, Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, Typography } from '@mui/material';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';
import { fetchWatchlists } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';
import axios from 'axios';
import { getJwtTokenFromCookies } from '../../utils/authTokenUtils'

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

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetchWatchlists();
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
  
  fetchData();
}, []);

const handleDeleteWatchlist = async (WatchlistId) => {
  try {
    console.log("delete button hit outer")
    const token = getJwtTokenFromCookies();
    if (!token) {
      console.error('Token not available or expired');
      // For now, will use a Promise.reject method instead of redirect
      return Promise.reject('Token not available or expired');
    }
    const headers = {
      Authorization: `Bearer ${token}`,
    };    

    await axios.delete(`http://localhost:8080/api/watchlist?id=${WatchlistId}`, { headers });
    // Update the watchlist in the state after a deletion
    setWatchlistData((prevItems) => {
      const currentItems = prevItems && prevItems['watchlists']; // Extract the array from the object
      const updatedItems = Array.isArray(currentItems)
        ? currentItems.filter((watchlist) => watchlist.id !== WatchlistId) // Re-render all watchlists not equal to the watchlistID that was deleted
        : [];
      return { 'watchlists': updatedItems }; // Maintain JSON object structure
    });
  } catch (error) {
    console.error('Error deleting watchlist:', error);
  }
 console.log("test")
};

const handleCreateWatchlistButtonClick = () => {
  setCreateWatchlistDialogOpen(true);
};

const handleCreateWatchlistButtonClose = () => {
  setCreateWatchlistDialogOpen(false);
  setDialogErrorMessage(''); // Clear error message when the dialog is closed
};

// API Call
const handleCreateWatchlistDialogSubmit = async () => {
  try {
    console.log("we in");
    /*
    const response = await editWatchlistName(watchlistID, newWatchlistName);
    if (response) {
      console.log("Successfully created")
    }
    */
  } catch (error) {
    if (error.message === errorConstants.ERROR_INVALID_NAME) {
      setDialogErrorMessage('Error: Name cannot be empty.');
    } else if (error.message === errorConstants.ERROR_BAD_REQUEST) {
      setDialogErrorMessage('Error: Bad request. Please try again.');
    } else {
      setDialogErrorMessage(`Error updating watchlist name: ${error.message}`);
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
      <DialogTitle>Create a New Watchlist</DialogTitle>
      <DialogContent>
        <TextField
          autoFocus
          id="watchlist-name"
          label="Enter Watchlist Name..."
          value={newWatchlistName}
          onChange={(e) => setNewWatchlistName(e.target.value)}
          multiline
          fullWidth
          margin="dense"
          variant="standard"
        />
        {dialogErrorMessage && (
          <Typography color="error" variant="body2">
            {dialogErrorMessage}
          </Typography>
        )}
      </DialogContent>
      <DialogActions>
        <Button variant="contained" onClick={handleCreateWatchlistButtonClose}>Cancel</Button>
        <Button variant="contained" onClick={handleCreateWatchlistDialogSubmit}>Submit</Button>
      </DialogActions>
    </Dialog>

    </React.Fragment>
    </ThemeProvider>
  );
};

export default ListOfWatchlists;