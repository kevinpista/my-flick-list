import React, { useState, useEffect} from 'react';
import { Container, Paper, Typography, Button } from '@mui/material';
import NavBar from '../NavBar.js';
import '../../css/Watchlist.css';

import MovieTable from './MovieTable';
// import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../../utils/formatUtils'; // Adjust the path to match your file structure

// Individual Watchlist that represents 1 single watchlist and holds up to 20 movies
// TODO
// 

const Watchlist = () => {
  // Hardcode movie data for now to style
  // Need to create use state + api call to fetch movies within a particular watchlist
  const movies = [
    {
      id: 1,
      toWatch: true,
      title: 'Movie 1',
      releaseDate: '2022-01-01',
      runtime: 120,
      budget: 50000000,
      revenue: 100000000,
    },
    {
      id: 2,
      toWatch: true,
      title: 'Movie 2',
      releaseDate: '2022-01-01',
      runtime: 90,
      budget: 123123,
      revenue: 4512362,
    },
    {
      id: 3,
      toWatch: true,
      title: 'Movie 3',
      releaseDate: '2022-01-01',
      runtime: 90,
      budget: 123123,
      revenue: 4512362,
    },
    {
      id: 4,
      toWatch: true,
      title: 'Movie 4',
      releaseDate: '2022-01-01',
      runtime: 90,
      budget: 123123,
      revenue: 4512362,
    },
    {
      id: 5,
      toWatch: true,
      title: 'Movie 5',
      releaseDate: '2022-01-01',
      runtime: 90,
      budget: 123123,
      revenue: 4512362,
    },
    {
      id: 6,
      toWatch: true,
      title: 'Movie 6',
      releaseDate: '2022-01-01',
      runtime: 90,
      budget: 123123,
      revenue: 4512362,
    },

    {
      id: 7,
      toWatch: true,
      title: 'Movie 7',
      releaseDate: '2022-01-01',
      runtime: 90,
      budget: 123123,
      revenue: 4512362,
    },

  ];

  return (
    <React.Fragment>
      <NavBar />
    <Container maxWidth={"xl"} className="movie-data-grid-container">
      <h1 className="watchlist-name">My Watchlist</h1>
      <MovieTable className="movie-data-grid-table" movies={movies} />
    </Container>

    </React.Fragment>

  );
};

export default Watchlist;
