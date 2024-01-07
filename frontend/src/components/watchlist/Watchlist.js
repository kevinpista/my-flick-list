import React, { useState, useEffect} from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Container, Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, Typography } from '@mui/material';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';
import { fetchWatchlistAndItems } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';
import axios from 'axios';
import { getJwtTokenFromCookies } from '../../utils/authTokenUtils'

import WatchlistItemsTable from './WatchlistItemsTable';

// TODO - if "error" is thrown relating to no movies in watchlist yet, still display
// the title and description, have message displayed, but include movie search bar for the user

// Individual Watchlist that represents 1 single watchlist and holds up to 20 movies

const Watchlist = () => {
  const { watchlistID } = useParams(); // Extract watchlistID from the URL params
  
  const [watchlistName, setWatchlistName] = useState(null);
  const [watchlistDescription, setWatchlistDescription] = useState(null);
  const [watchlistItems, setWatchlistItems] = useState(null); // In JSON object format
  const [error, setError] = useState(null);

  const [isEditNameDialogOpen, setEditNameDialogOpen] = useState(false);
  const [newWatchlistName, setNewWatchlistName] = useState('');
  const [newWatchlistDescription, setNewWatchlistDescription] = useState('');
  const [dialogErrorMessage, setDialogErrorMessage] = useState(''); // Use 1 for both Name & Description edits 


  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetchWatchlistAndItems(watchlistID);
        setWatchlistItems(response);
        setWatchlistName(response['name']);
        setWatchlistDescription(response['description']);
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

const handleEditNameButtonClick = () => {
  setEditNameDialogOpen(true);
};

const handleEditNameDialogClose = () => {
  setEditNameDialogOpen(false);
  setDialogErrorMessage(''); // Clear error message when the dialog is closed
};

const handleEditNameDialogSubmit = async () => {
  try {
    const token = getJwtTokenFromCookies();
    if (!token) {
      console.error('Token not available or expired');
      return Promise.reject('Token not available or expired');
    }

    const headers = {
      Authorization: `Bearer ${token}`,
    };

    const response = await axios.patch(
      `http://localhost:8080/api/watchlist-name?id=${watchlistID}`,
      { name: newWatchlistName },
      { headers }
    );

    setWatchlistName(response.data.name);
    setEditNameDialogOpen(false);

  } catch (error) {
    console.error('Error updating watchlist name:', error);
    setDialogErrorMessage(`Error updating watchlist name: ${error.message}`);
  }
};

  return (
    <React.Fragment>
      <NavBar />
      <Container maxWidth={"xl"} className="watchlist-item-grid-container">
        <div className="watchlist-name-div">
          <h1 className="watchlist-name">{watchlistName}</h1>
          <Button variant="outlined" onClick={handleEditNameButtonClick}>
            Edit Watchlist Name
          </Button>
        </div>
        <p className="watchlist-description">{watchlistDescription}</p>
        {error ? (
          <h1 className='error'><u>Error:</u> {error.message}</h1>
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
      {/* Modal for editing watchlist name */}
      <Dialog open={isEditNameDialogOpen} onClose={handleEditNameDialogClose}>
        <DialogTitle>Edit Watchlist Name</DialogTitle>
        <DialogContent>
          <TextField
            autofocus
            label="New Watchlist Name"
            value={newWatchlistName}
            onChange={(e) => setNewWatchlistName(e.target.value)}
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
          <Button onClick={handleEditNameDialogClose}>Cancel</Button>
          <Button onClick={handleEditNameDialogSubmit}>Submit</Button>
        </DialogActions>
      </Dialog>
    </React.Fragment>
  );
};

export default Watchlist;
