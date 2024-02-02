import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';

import '../css/MovieSearchBar.css';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';
import SearchIcon from '@mui/icons-material/Search';

// Args for search bar TextField's size & label
function MovieSearchBar({ sizeSx ="medium", labelText="Search for any movie..."}) { 
  const [searchQuery, setSearchQuery] = useState('');

  const handleSearch = () => {
    const encodedQuery = encodeURIComponent(searchQuery);
    window.location.href = `/movie-search?query=${encodedQuery}`;
  };
  // Checks length of
  const searchDoesNotQueryExists = searchQuery.trim().length === 0;

  return (
    <ThemeProvider theme={muiTheme}>
    <div className="search-bar-wrapper">
      <TextField
        className="search-text-field"
        label={labelText}
        type="search"
        variant="filled"
        color="primary"
        style={{background:"#ededed"}}
        value={searchQuery}
        onChange={(e) => setSearchQuery(e.target.value)}
        fullWidth
        onKeyDown={(event) => {
            if (event.key === 'Enter') {
                event.preventDefault();
                handleSearch();
            }
        }}
        size={sizeSx}
      />
      <Button 
        className="search-button" 
        variant="contained" 
        onClick={searchDoesNotQueryExists ? null: handleSearch}
      >
        <SearchIcon />
      </Button>
    </div>
    </ThemeProvider>
  );
}
export default MovieSearchBar;