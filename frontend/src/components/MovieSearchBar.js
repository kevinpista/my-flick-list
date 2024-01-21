import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';

import '../css/MovieSearchBar.css';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';
import SearchIcon from '@mui/icons-material/Search';


function MovieSearchBar() {
  const [searchQuery, setSearchQuery] = useState('');

  const handleSearch = () => {
    const encodedQuery = encodeURIComponent(searchQuery);
    window.location.href = `/movie-search?query=${encodedQuery}`;
  };

  return (
    <ThemeProvider theme={muiTheme}>
    <div className="search-bar">
      <TextField
        label="Search for a movie..."
        type="search"
        variant="filled"
        color="primary"
        style={{background:"#e0e0e0"}}
        value={searchQuery}
        onChange={(e) => setSearchQuery(e.target.value)}
        fullWidth
        onKeyDown={(event) => {
            if (event.key === 'Enter') {
                event.preventDefault();
                handleSearch();
            }
        }}
      />
      <Button 
        className="search-button" 
        variant="contained" 
        onClick={handleSearch}
      >
        <SearchIcon />
      </Button>
    </div>
    </ThemeProvider>
  );
}
export default MovieSearchBar;