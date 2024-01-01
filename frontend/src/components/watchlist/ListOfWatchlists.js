import React, { useState, useEffect} from 'react';
import { Container } from '@mui/material';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';
import { fetchWatchlists } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';
import axios from 'axios';
import { getJwtTokenFromCookies } from '../../utils/authTokenUtils'

import ListOfWatchlistsTable from './ListOfWatchlistsTable.js';

// List of Watchlists; states name of watchlist, its description, and how many movies inside
// Accessed via URL of /watchlists

const ListOfWatchlists = () => {
  const [watchlistData, setWatchlistData] = useState(null); // In JSON object format
  const [error, setError] = useState(null);

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

  return (
    <React.Fragment>
      <NavBar />
    <Container maxWidth={"xl"} className="watchlist-item-grid-container">
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
    </React.Fragment>
  );
};

export default ListOfWatchlists;