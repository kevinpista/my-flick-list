import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';

import '../css/MovieSearchBar.css';

function MovieSearchBar() {
  const [searchQuery, setSearchQuery] = useState('');

  const handleSearch = () => {
    const encodedQuery = encodeURIComponent(searchQuery);
    window.location.href = `/movie-search?query=${encodedQuery}`;
  };

  return (
    <div className="search-bar">
      <TextField
        label="Search for a movie..."
        variant="outlined"
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
      <Button className="search-button" variant="contained" onClick={handleSearch}>
        Search
      </Button>
    </div>
  );
}

export default MovieSearchBar;