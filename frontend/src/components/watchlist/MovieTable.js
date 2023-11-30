import React from 'react';
import { DataGrid } from '@mui/x-data-grid';

// TODO
// Movie title link to website.com/movie{id} - should already be added to DB
// Genre listed
// My notes icon popup module


const columns = [
  {
    field: 'toWatch',
    headerName: 'To-Watch',
    width: 100,
    renderCell: (params) => (
      <div style={{ textAlign: 'center' }}>
        {params.value ? '✔' : ''}
      </div>
    ),
  },
  { field: 'title', headerName: 'Title', width: 300 },
  { field: 'releaseDate', headerName: 'Release Date', width: 150 },
  { field: 'runtime', headerName: 'Runtime', width: 120 },
  { field: 'rating', headerName: 'Ratings', width: 120 },
  { field: 'budget', headerName: 'Budget', width: 120 },
  { field: 'revenue', headerName: 'Revenue', width: 120 },
];

const MovieTable = ({ movies }) => {
  // Assuming 'movies' is an array of movie objects

  const rows = movies.map((movie) => ({
    id: movie.id, // movie id will the idenfitier
    toWatch: movie.toWatch, // will be a boolean value
    title: movie.title,
    releaseDate: movie.releaseDate,
    runtime: `${movie.runtime} mins`,
    budget: `$${movie.budget}`,
    revenue: `$${movie.revenue}`,
  }));

  return (
    <div style={{ height: '100%', width: '100%' }}>
      <DataGrid
        rows={rows}
        columns={columns}
        pageSize={5}
        checkboxSelection
        disableSelectionOnClick
        autoHeight={true}
      />
    </div>
  );
};

export default MovieTable;