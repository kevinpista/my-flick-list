import React, { useState } from 'react';
import { DataGrid } from '@mui/x-data-grid';
import { formatReleaseDate, formatRuntime, formatFinancialData } from '../../utils/formatUtils';

// MUI Dialog component to confirm watchlist item deletion
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';

import ClearSharpIcon from '@mui/icons-material/ClearSharp';
import { ThemeProvider } from '@mui/material/styles';
import * as themeStyles from '../../styling/ThemeStyles';

// TODO
// Movie title link to website.com/movie{id} - should already be added to DB
// My notes icon popup module

// watchlistItems is a JSON object holding 1 array containing individual movie data for each watchlistItem
const WatchlistItemsTable = ({ watchlistItems, onDeleteWatchlistItem }) => {
  // Dialog component useState
  const [openConfirmation, setOpenConfirmation] = useState(false);
  const [deleteItemId, setDeleteItemId] = useState(null);    

  const [editRowsModel, setEditRowsModel] = useState({});

  const handleToWatchClick = (event, row) => {
    event.stopPropagation();
    setEditRowsModel((prev) => ({
      ...prev,
      [row.id]: { ...prev[row.id], toWatch: !row.toWatch },
    }));
  };

  const handleDeleteClick = (event, watchlistItemId) => {
    event.stopPropagation();
    setDeleteItemId(watchlistItemId);
    setOpenConfirmation(true);
  };

  const handleCloseConfirmation = () => {
    setOpenConfirmation(false);
    setDeleteItemId(null);
  };

  const handleConfirmDelete = (watchlistItemId) => {
    handleCloseConfirmation();
    onDeleteWatchlistItem(watchlistItemId); // Call this function with the watchlistItemId to be deleted
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

    {
      field: 'deleteButton',
      headerName: 'Delete',
      width: 100,
      headerAlign: 'center',
      align: 'center',
      renderCell: (params) => (
        <div style={{ textAlign: 'center', cursor: 'pointer' }} onClick={(e) => handleDeleteClick(e, params.row.id)}>
         <ThemeProvider theme={themeStyles.formTheme}>
            <ClearSharpIcon fontSize="medium"/>
          </ThemeProvider> 
        </div>
      ),
    },
  ];
  
  // Will likely rename this to watchlist_item instead of "movie"
  const rows = watchlistItems['watchlist-items'].map((watchlistItem) => ({
    id: watchlistItem.id, // watchlist item id
    movie_id: watchlistItem.movie_id,
    toWatch: watchlistItem.checkmarked,
    moviePoster: `https://image.tmdb.org/t/p/w200${watchlistItem.poster_path}`, // Loading 200 width poster from API, resize to 80 width
    title: watchlistItem.original_title,
    releaseDate: formatReleaseDate(watchlistItem.release_date),
    runtime: formatRuntime(watchlistItem.runtime),
    rating: watchlistItem.rating,
    budget: formatFinancialData(watchlistItem.budget),
    revenue: formatFinancialData(watchlistItem.revenue),
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
      {/* Confirmation Dialog for Deletion */}
      <Dialog
        open={openConfirmation}
        onClose={handleCloseConfirmation}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">Confirm Deletion</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Are you sure you want to remove this movie from your watchlist?
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseConfirmation}>Cancel</Button>
          <Button onClick={() => handleConfirmDelete(deleteItemId)} autoFocus>
            Confirm
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default WatchlistItemsTable;