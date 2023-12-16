import React, { useState, useEffect} from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Container, Paper, Typography, Button } from '@mui/material';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';
import { fetchWatchlistItems } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';
import axios from 'axios';
import { getJwtTokenFromCookies } from '../../utils/authTokenUtils'

import WatchlistItemsTable from './WatchlistItemsTable';
// import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../../utils/formatUtils'; // Adjust the path to match your file structure

// Individual Watchlist that represents 1 single watchlist and holds up to 20 movies
// TODO
// Max hold 20 movies. Checkmark change. Able to fetch notes data as well.
// Need to be able to make changes to "toWatch" and send to backend and also notes
// Convert data Release date runtime, budget, revenue with helper functions
// Store movie ID url to titles so user is directed to the individual movie page to view

const Watchlist = () => {
  const { watchlistID } = useParams(); // Extract watchlistID from the URL params
  const [watchlistItems, setWatchlistItems] = useState(null); // In JSON object format
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetchWatchlistItems(watchlistID);
        setWatchlistItems(response);
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

  return (
    <React.Fragment>
      <NavBar />
    <Container maxWidth={"xl"} className="watchlist-item-grid-container">
      <h1 className="watchlist-name">My Watchlist</h1>
      {error ? (
        <p> Error loading watchlist: {error.message}</p>
      ) : (
        watchlistItems && (
        <WatchlistItemsTable 
          watchlistItems={watchlistItems}
          onDeleteWatchlistItem={handleDeleteWatchlistItem} // onDeleteWatchlistItem function gets passed to component
        />
        )
      )}
    </Container>
    </React.Fragment>
  );
};

export default Watchlist;
