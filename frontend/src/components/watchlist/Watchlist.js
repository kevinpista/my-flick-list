import React, { useState, useEffect} from 'react';
import { useParams } from 'react-router-dom';
import { Container, Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, Typography } from '@mui/material';
import InputAdornment from '@mui/material/InputAdornment';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';
import { fetchWatchlistAndItems, editWatchlistName, editWatchlistDescription } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';
import axios from 'axios';
import { getJwtTokenFromCookies } from '../../utils/authTokenUtils'

import WatchlistItemsTable from './WatchlistItemsTable';

import { ThemeProvider } from '@mui/material/styles';
import {muiTheme} from '../../css/MuiThemeProvider.js';

// TODO - if "error" is thrown relating to no movies in watchlist yet, still display
// the title and description, have message displayed, but include movie search bar for the user

// Individual Watchlist that represents 1 single watchlist and holds up to 20 movies

const Watchlist = () => {
  const { watchlistID } = useParams(); // Extract watchlistID from the URL params
  
  const [watchlistName, setWatchlistName] = useState(null);
  const [watchlistDescription, setWatchlistDescription] = useState(null);
  const [watchlistItems, setWatchlistItems] = useState(null); // In JSON object format
  const [error, setError] = useState(null);

  // Dialog forms to edit name & description
  const [isEditNameDialogOpen, setEditNameDialogOpen] = useState(false);
  const [isEditDescriptionDialogOpen, setEditDescriptionDialogOpen] = useState(false);
  const [newWatchlistName, setNewWatchlistName] = useState('');
  const [newWatchlistDescription, setNewWatchlistDescription] = useState('');
  const [dialogErrorMessage, setDialogErrorMessage] = useState(''); // Use 1 for both Name & Description edits 


  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetchWatchlistAndItems(watchlistID);
        // Still want to render Name + Description if there are no items added yet
        if (response['watchlist-items'] === null) {
          setWatchlistName(response['name']);
          setWatchlistDescription(response['description']);
          setError(Error("You have not added any movies yet."));
        } else {
          setWatchlistItems(response);
          setWatchlistName(response['name']);
          setWatchlistDescription(response['description']);
        }
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
}, [watchlistID]);

const handleDeleteWatchlistItem = async (watchlistItemId) => {
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
    
    await axios.delete(`http://localhost:8080/api/watchlist-item?id=${watchlistItemId}`, { headers });
    // Update the watchlist items in the state after a deletion
    setWatchlistItems((prevItems) => {
      const currentItems = prevItems && prevItems['watchlist-items']; // Extract the array from the object
      const updatedItems = Array.isArray(currentItems)
        ? currentItems.filter((item) => item.id !== watchlistItemId) // Re-render all items not equal to the itemID that was deleted
        : [];
      return { 'watchlist-items': updatedItems }; // Maintain JSON object structure
    });
  } catch (error) {
    console.error('Error deleting item:', error);
  }
};

// Handle Edit Watchlist Name
const handleEditNameButtonClick = () => {
  setEditNameDialogOpen(true);
};

const handleEditNameDialogClose = () => {
  setEditNameDialogOpen(false);
  setDialogErrorMessage(''); // Clear error message when the dialog is closed
};

const handleEditNameDialogSubmit = async () => {
  try {
    const response = await editWatchlistName(watchlistID, newWatchlistName);
    if (response) {
      setWatchlistName(response.name);
      setEditNameDialogOpen(false);
    }
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

// Handle Edit Watchlist Description
const handleEditDescriptionButtonClick = () => {
  setEditDescriptionDialogOpen(true);
};

const handleEditDescriptionDialogClose = () => {
  setEditDescriptionDialogOpen(false);
  setDialogErrorMessage(''); // Clear error message when the dialog is closed
};

const handleEditDescriptionDialogSubmit = async () => {
  try {
    const response = await editWatchlistDescription(watchlistID, newWatchlistDescription);
    if (response) {
      setWatchlistDescription(response.description);
      setEditDescriptionDialogOpen(false);
    }
  } catch (error) {
    if (error.message === errorConstants.ERROR_INVALID_NAME) {
      setDialogErrorMessage('Error: Description cannot be empty.');
    } else if (error.message === errorConstants.ERROR_BAD_REQUEST) {
      setDialogErrorMessage('Error: Bad request. Please try again.');
    } else {
      setDialogErrorMessage(`Error updating watchlist description: ${error.message}`);
    }
  };
};

  return (
    <ThemeProvider theme={muiTheme}>
    <React.Fragment>
      <NavBar />
      <div className="watchlist-root">
      <Container maxWidth={"xl"} className="watchlist-item-grid-container">
        <div className="watchlist-name-div">
          <h1 className="watchlist-name">{watchlistName}</h1>
        </div>

        <p className="watchlist-description">{watchlistDescription}</p>
        {error ? (
          <h1 className='error'><u>Heads Up:</u> {error.message}</h1>
        ) : (
          watchlistItems && (
          <WatchlistItemsTable 
            watchlistItems={watchlistItems}
            onDeleteWatchlistItem={handleDeleteWatchlistItem} // onDeleteWatchlistItem function gets passed to component. When called, it invokes handleDeleteWatchlistItem
            setWatchlistItems={setWatchlistItems} // Also passed to component
          />
          )
        )}
      </Container>
      <div className="watchlist-buttons-table">
        <Button variant="contained" onClick={handleEditNameButtonClick}>
          Edit Watchlist Name
        </Button>
        <Button variant="contained" onClick={handleEditDescriptionButtonClick}>
          Edit Description
        </Button>
        </div>
      </div>

      {/* Modal for editing watchlist name */}
      <Dialog
        open={isEditNameDialogOpen}
        onClose={handleEditNameDialogClose}
        maxWidth="md"
        fullWidth={true}
      >
        <DialogTitle>Edit Watchlist Name</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            label="Enter new name..."
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
          {dialogErrorMessage && (
            <Typography color="error" variant="body2">
              {dialogErrorMessage}
            </Typography>
          )}
        </DialogContent>
        <DialogActions>
          <Button variant="contained" onClick={handleEditNameDialogClose}>Cancel</Button>
          <Button
          variant="contained"
          onClick={handleEditNameDialogSubmit}
          disabled={
            newWatchlistName.length > 60 // Character limit for watchlist name
          }
          >
            Submit
          </Button>
        </DialogActions>
      </Dialog>

      {/* Modal for editing watchlist description */}
      <Dialog 
        open={isEditDescriptionDialogOpen}
        onClose={handleEditDescriptionDialogClose}
        maxWidth="lg"
        fullWidth={true}
      >
        <DialogTitle>Edit Watchlist Description</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            label="Enter new description..."
            value={newWatchlistDescription}
            onChange={(e) => setNewWatchlistDescription(e.target.value)}
            multiline
            maxWidth="lg"
            fullWidth={true}
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
          <Button variant="contained" onClick={handleEditDescriptionDialogClose}>Cancel</Button>
          <Button
          variant="contained"
          onClick={handleEditDescriptionDialogSubmit}
          disabled={
            newWatchlistDescription.length > 500 // Character limit for watchlist description
          }
          >
            Submit
          </Button>
        </DialogActions>
      </Dialog>
    </React.Fragment>
    </ThemeProvider>
  );
};

export default Watchlist;
