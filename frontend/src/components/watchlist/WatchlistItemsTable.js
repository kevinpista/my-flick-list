import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { DataGrid } from '@mui/x-data-grid';
import { formatReleaseDate, formatRuntime, formatFinancialData } from '../../utils/formatUtils';
import { getJwtTokenFromCookies } from '../../utils/authTokenUtils'
import axios from 'axios';

// MUI Dialog component to confirm watchlist item deletion
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';

import ClearSharpIcon from '@mui/icons-material/ClearSharp';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../../css/MuiThemeProvider.js';
import * as themeStyles from '../../styling/ThemeStyles';

// MUI Checkbox component for 'checkmarked' state
import Checkbox from '@mui/material/Checkbox';


// TODO
// My notes icon popup module

// watchlistItems is a JSON object holding 1 array containing individual movie data for each watchlistItem
const WatchlistItemsTable = ({ watchlistItems, onDeleteWatchlistItem, setWatchlistItems }) => {
  // Dialog component useState
  const [openConfirmation, setOpenConfirmation] = useState(false);
  const [deleteItemId, setDeleteItemId] = useState(null);    

  const [editRowsModel, setEditRowsModel] = useState({});

  const handleToWatchClick = async (event, row) => {
    event.stopPropagation();
    try {
      const token = getJwtTokenFromCookies();
      if (!token) {
        console.error('Token not available or expired');
        return Promise.reject('Token not available or expired');
      }
  
      const headers = {
        Authorization: `Bearer ${token}`,
      };
  
      const response = await axios.put(
        `http://localhost:8080/api/watchlist-item-checkmarked?id=${row.id}`, // row.id is the watchlistItemID
        { checkmarked: !row.toWatch },
        { headers }
      );
  
      if (response.status === 200) {
        // Update the state to reflect the new checkmarked status        
        setWatchlistItems((prevItems) => {
          const updatedWatchlistItems = prevItems['watchlist-items'].map((watchlistItem) => { // Performs the actual mapping
            // Create new array of updatedWatchlistItems by taking current old array and creating a copy
            // once it finds the row of the one that is being updated with if block below, it will only make chnages to that row
            if (watchlistItem.id === row.id) {
              return { ...watchlistItem, checkmarked: !row.toWatch }; // Spread operator keeps this original row data the same minus checkmarked field
            }
            return watchlistItem; // If row matching watchlistItem.id is not found, returns original state as updatedItems
          });
          return { 'watchlist-items': updatedWatchlistItems }; // Sent to Watchlist.js useState function which sets with the updated data
          // and is then eventually passed back to here to be re-rendered
        });
      } else {
        console.error('Failed to update checkmarked status');
      }
    } catch (error) {
        console.error('Error updating checkmarked status:', error);
    }
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
    {
      field: 'toWatch',
      headerName: 'Checkmark',
      width: 100,
      headerAlign: 'center',
      align: 'center',
      renderCell: (params) => (
        <Checkbox
          checked={params.row.toWatch}
          onChange={(e) => handleToWatchClick(e, params.row)}
        />
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
        <Link to={`/movie/${params.row.movie_id}`}>
          <img
            src={params.value}
            alt={`${params.row.title} Poster`}
            style={{ width: 80, height: 112 }}
          />
        </Link>
      ),
    },
    { 
      field: 'title', 
      headerName: 'Title', 
      width: 300, 
      headerAlign: 'center', 
      align: 'center',
      renderCell: (params) => (
        <Link to={`/movie/${params.row.movie_id}`} style={{ color: 'black', textDecoration: 'none' }}>
        {params.row.title}
        </Link>
      ),
     },
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
          <ThemeProvider theme={muiTheme}>
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
        hideFooterPagination
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
          <Button variant = "contained" onClick={handleCloseConfirmation}>Cancel</Button>
          <Button variant = "contained" onClick={() => handleConfirmDelete(deleteItemId)} autoFocus>
            Confirm
          </Button>
        </DialogActions>
      </Dialog>
      </ThemeProvider>
    </div>
  );
};

export default WatchlistItemsTable;