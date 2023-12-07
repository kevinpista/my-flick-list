import React, { useState, useEffect} from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Container, Paper, Typography, Button } from '@mui/material';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';
import { fetchWatchlistItems } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';

import MovieTable from './MovieTable';
// import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../../utils/formatUtils'; // Adjust the path to match your file structure

// Individual Watchlist that represents 1 single watchlist and holds up to 20 movies
// TODO
// Max hold 20 movies. Checkmark change. Able to fetch notes data as well.
// Need to be able to make changes to "toWatch" and send to backend and also notes
// Convert data Release date runtime, budget, revenue with helper functions
// Store movie ID url to titles so user is directed to the individual movie page to view

const Watchlist = () => {
  const { watchlistID } = useParams(); // Extract watchlistID from the URL params
  const [movies, setMovies] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetchWatchlistItems(watchlistID);
        setMovies(response);
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

  return (
    <React.Fragment>
      <NavBar />
    <Container maxWidth={"xl"} className="movie-data-grid-container">
      <h1 className="watchlist-name">My Watchlist</h1>
      {error ? (
        <p> Error loading watchlist: {error.message}</p>
      ) : (
        movies && <MovieTable className="movie-data-grid-table" movies={movies} />
      )}
    </Container>
    </React.Fragment>
  );
};

export default Watchlist;
