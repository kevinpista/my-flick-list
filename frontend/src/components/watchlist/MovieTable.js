import React, { useState } from 'react';
import { DataGrid } from '@mui/x-data-grid';
import { formatReleaseDate, formatRuntime, formatFinancialData } from '../../utils/formatUtils';

// TODO
// Movie title link to website.com/movie{id} - should already be added to DB
// My notes icon popup module


const MovieTable = ({ movies }) => {
  // Assuming 'movies' is an array of movie objects

  const [editRowsModel, setEditRowsModel] = useState({});

  const handleToWatchClick = (event, row) => {
    event.stopPropagation();
    setEditRowsModel((prev) => ({
      ...prev,
      [row.id]: { ...prev[row.id], toWatch: !row.toWatch },
    }));
  };

  const getRowId = (row) => row.id;

  const rowHeight = 140; // Fixed height for each row

  const columns = [
    { field: 'movie_id', headerName: 'Movie ID' },
    {
      field: 'toWatch',
      headerName: 'Checkmark',
      width: 100,
      headerAlign: 'center',
      align: 'center',
      renderCell: (params) => (
        <div
          style={{ textAlign: 'center', cursor: 'pointer' }}
          onClick={(e) => params.row.id && handleToWatchClick(e, params.row)}
        >
          {params.row.toWatch ? '✔' : 'X'}
        </div>
      ),
      editable: true,
    },
    {
      field: 'moviePoster',
      headerName: 'Poster',
      width: 80,
      headerAlign: 'center',
      align: 'center',
      renderCell: (params) => (
        <img
          src={params.value}
          alt={`${params.row.title} Poster`}
          style={{ width: 80, height: 112 }}
        />
      ),
    },
    { field: 'title', headerName: 'Title', width: 300, headerAlign: 'center', align: 'center' },
    { field: 'releaseDate', headerName: 'Release Date', width: 150, headerAlign: 'center', align: 'center' },
    { field: 'runtime', headerName: 'Runtime', width: 120, headerAlign: 'center', align: 'center' },
    { field: 'rating', headerName: 'Ratings', width: 120, headerAlign: 'center', align: 'center' },
    { field: 'budget', headerName: 'Budget', width: 120, headerAlign: 'center', align: 'center' },
    { field: 'revenue', headerName: 'Revenue', width: 120, headerAlign: 'center', align: 'center' },
    { field: 'notes', headerName: 'Notes', width: 120, headerAlign: 'center', align: 'center' },
  ];
  
  // Will likely rename this to watchlist_item instead of "movie"
  const rows = movies['watchlist-items'].map((movie) => ({
    id: movie.id,
    movie_id: movie.movie_id,
    toWatch: movie.checkmarked,
    moviePoster: `https://image.tmdb.org/t/p/w200${movie.poster_path}`, // Loading 200 width poster from API, resize to 80 width
    title: movie.original_title,
    releaseDate: formatReleaseDate(movie.release_date),
    runtime: formatRuntime(movie.runtime),
    rating: movie.rating,
    budget: formatFinancialData(movie.budget),
    revenue: formatFinancialData(movie.revenue),
  }));

  return (
    <div style={{ height: '100%', width: '100%' }}>
      <DataGrid
        rows={rows}
        columns={columns}
        pageSize={5}
        disableRowSelectionOnClick
        editRowsModel={editRowsModel}
        onEditRowsModelChange={(newModel) => setEditRowsModel(newModel)}
        autoHeight={false}
        disableMultipleRowSelection={true}
        getRowId={getRowId}
        rowHeight={rowHeight}
      />
    </div>
  );
};

export default MovieTable;