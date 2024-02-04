import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { DataGrid } from '@mui/x-data-grid';
import { formatReleaseDate, formatRuntime, formatFinancialData } from '../../utils/formatUtils';
import { getJwtTokenFromCookies } from '../../utils/authTokenUtils';
import { editWatchlistItemNoteAPI, createWatchlistItemNoteAPI } from '../../api/watchlistAPI';
import * as errorConstants from '../../api/errorConstants';
import axios from 'axios';

import { Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Button, Typography } from '@mui/material';
import InputAdornment from '@mui/material/InputAdornment';

import ClearSharpIcon from '@mui/icons-material/ClearSharp';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../../css/MuiThemeProvider.js';
import * as themeStyles from '../../styling/ThemeStyles';
import '../../css/WatchlistItemsTable.css';

// MUI Checkbox component for 'checkmarked' state
import Checkbox from '@mui/material/Checkbox';

// Item Note icons
import IconButton from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import ChatIcon from '@mui/icons-material/Chat';
import ChatBubbleIcon from '@mui/icons-material/ChatBubble';

import TextField from '@mui/material/TextField';


// watchlistItems is a JSON object holding 1 array containing individual movie data for each watchlistItem
const WatchlistItemsTable = ({ watchlistItems, onDeleteWatchlistItem, setWatchlistItems }) => {
  // useStates for Deletion Confirmation Dialog box
  const [openDeletionConfirmation, setOpenDeletionConfirmation] = useState(false);
  const [deleteItemId, setDeleteItemId] = useState(null);    

  // useStates for Note Dialog box
  const [openNoteDialog, setOpenNoteDialog] = useState(false);
  const [selectedNote, setSelectedNote] = useState('');
  const [selectedNoteUpdatedAt, setSelectedNoteUpdatedAt] = useState(null);
  const [editedNote, setEditedNote] = useState('');
  const [isEditingNote, setIsEditingNote] = useState(false);
  const [isCreatingNote, setIsCreatingNote] = useState(false);
  const [dialogNoteErrorMessage, setDialogNoteErrorMessage] = useState(''); // Dialog error display for editing notes 
  
  const [editRowsModel, setEditRowsModel] = useState({});
  const [selectedWatchlistItemId, setSelectedWatchlistItemId] = useState(null);

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
  
  const handleNoteIconClick = (itemNote, watchlistItemId, itemNoteUpdatedAt) => {
    setSelectedNote(itemNote);
    setSelectedWatchlistItemId(watchlistItemId)
    if (itemNote === null) {
      setEditedNote(''); // Sets initial editing TextField
    } else {
      setEditedNote(itemNote); // Set the initial value of the text field to the current note
    }
    if (itemNoteUpdatedAt !== null) {
      setSelectedNoteUpdatedAt(formatDate(itemNoteUpdatedAt)) 
    }
    setOpenNoteDialog(true);
  };

  const handleNoteDialogClose = () => {
    setOpenNoteDialog(false);
    setSelectedNote('');
    setSelectedNoteUpdatedAt(null);
    setDialogNoteErrorMessage('');

    if (isEditingNote && editedNote === selectedNote) {
      // Only set if dialog is open and there were no edits made
      setIsEditingNote(false);
    }
    if (isCreatingNote && editedNote === selectedNote) {
      // Only set if dialog is open and there were no creations made
      setIsCreatingNote(false);
      setIsEditingNote(false);
    }
  };

  const handleDeleteClick = (event, watchlistItemId) => {
    event.stopPropagation();
    setDeleteItemId(watchlistItemId);
    setOpenDeletionConfirmation(true);
  };

  const handleCloseConfirmation = () => {
    setOpenDeletionConfirmation(false);
    setDeleteItemId(null);
  };

  const handleConfirmDelete = (watchlistItemId) => {
    handleCloseConfirmation();
    onDeleteWatchlistItem(watchlistItemId); // Call this function with the watchlistItemId to be deleted
  };

  const handleNoteCreateSubmit = async () => {
    try {
      const response = await createWatchlistItemNoteAPI(selectedWatchlistItemId, editedNote); // selectedNote gets updated in dialog textfield
      if (response.status === 200) {
        // Create and stores the item_notes in the local state. Loop original array of data until edited row found
        const updatedWatchlistItems = watchlistItems['watchlist-items'].map((watchlistItem) => {
          if (watchlistItem.id === selectedWatchlistItemId) {
            return {
              ...watchlistItem,
              item_notes: response.data.item_notes, // Using the updated note from the API response
            };
          }
          return watchlistItem; // Catch all in case selecetedWatchlistItemId is not found
        });
        // Set the updated state
        setWatchlistItems({ 'watchlist-items': updatedWatchlistItems });
        setSelectedNote(response.data.item_notes);
        setIsEditingNote(false);
        setIsCreatingNote(false);
        setDialogNoteErrorMessage('');
        setSelectedNoteUpdatedAt(formatDate(new Date().toISOString())); // Set to current time
      }
    } catch (error) {
    if (error.message === errorConstants.ERROR_BAD_REQUEST) {
        setDialogNoteErrorMessage('Bad request, please try again.');
      } else {
        setDialogNoteErrorMessage(`Error creating note: ${error.message}`);
      }
    };
  };

  const handleNoteUpdateSubmit = async () => {
    try {
      const response = await editWatchlistItemNoteAPI(selectedWatchlistItemId, editedNote); // selectedNote gets updated in dialog textfield
      if (response.status === 200) {
        const updatedWatchlistItems = watchlistItems['watchlist-items'].map((watchlistItem) => {
          if (watchlistItem.id === selectedWatchlistItemId) {
            return {
              ...watchlistItem,
              item_notes: response.data.item_notes,
            };
          }
          return watchlistItem;
        });
        setWatchlistItems({ 'watchlist-items': updatedWatchlistItems });
        setSelectedNote(response.data.item_notes);
        setIsEditingNote(false);
        setDialogNoteErrorMessage('');
        setSelectedNoteUpdatedAt(formatDate(new Date().toISOString())); // Set to current time
      }
    } catch (error) {
    if (error.message === errorConstants.ERROR_BAD_REQUEST) {
        setDialogNoteErrorMessage('Bad request, please try again.');
      } else {
        setDialogNoteErrorMessage(`Error updating note: ${error.message}`);
      }
    };
  };
  // 'Last Updated At' date found in the Notes Dialog
  const formatDate = (dateString) => {
    const date = new Date(dateString);
    const options = {
      month: '2-digit',
      day: 'numeric',
      year: '2-digit',
      hour: 'numeric',
      minute: '2-digit',
      hour12: true,
    };
    const formattedDate = date.toLocaleString('en-US', options);

    // Extract date & time parts, removing comma from date part
    const datePart = formattedDate.slice(0, formattedDate.lastIndexOf(',')).trim();
    const timePart = formattedDate.slice(formattedDate.lastIndexOf(' ') -5);
    return `${datePart} at ${timePart}`;
  };
  
  const getRowId = (row) => row.id;
  const rowHeight = 140; // Fixed height for each row
  const columns = [
    {
      field: 'toWatch',
      renderHeader: () => (
        <span className="columnHeader">
          Watched
        </span>
      ),      
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
      renderHeader: () => (
        <span className="columnHeader">
          Poster
        </span>
      ),     
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
      renderHeader: () => (
        <span className="columnHeader">
          Title
        </span>
      ),        
      width: 300, 
      headerAlign: 'center', 
      align: 'center',
      renderCell: (params) => (
        <Link to={`/movie/${params.row.movie_id}`} style={{ color: 'black', textDecoration: 'none' }}>
        {params.row.title}
        </Link>
      ),
     },
    { field: 'releaseDate', renderHeader: () => (<span className="columnHeader">Release Date</span>), width: 150, headerAlign: 'center', align: 'center' },
    { field: 'runtime', renderHeader: () => (<span className="columnHeader">Runtime</span>), width: 120, headerAlign: 'center', align: 'center' },
    { field: 'rating', renderHeader: () => (<span className="columnHeader">Ratings</span>), width: 120, headerAlign: 'center', align: 'center' },
    { field: 'budget', renderHeader: () => (<span className="columnHeader">Budget</span>), width: 120, headerAlign: 'center', align: 'center' },
    { field: 'revenue', renderHeader: () => (<span className="columnHeader">Revenue</span>), width: 120, headerAlign: 'center', align: 'center' },
    { 
      field: 'noteIcon', 
      renderHeader: () => (
        <span className="columnHeader">
          Notes
        </span>
      ),       
      width: 100, 
      headerAlign: 'center', 
      align: 'center', 
      renderCell: (params) => (
        <Tooltip title={params.row.item_notes ? 'Click to view note' : 'No notes found'}>
          <IconButton onClick={() => handleNoteIconClick(params.row.item_notes, params.row.id, params.row.note_updated_at)}>
          {params.row.item_notes === null ? (
            <ChatBubbleIcon color='primary'/>  
          ) : params.row.item_notes === "" ? (
            <ChatBubbleIcon color='primary' />
          ) : (
            <ChatIcon color='primary' />
          )}
          </IconButton>
        </Tooltip>
      ),
    },

    {
      field: 'deleteButton',
      renderHeader: () => (
        <span className="columnHeader">
          Delete
        </span>
      ),             
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
    item_notes: watchlistItem.item_notes,
    note_updated_at: watchlistItem.note_updated_at
  }));

  // Components to render for item notes dialog based on 3 cases
  // Note exists, show current note, show edit option, submit sends PATCH request
  // Note exists but is empty "", show 'note is empty' message, show edit option, submit sends PATCH request
  // Note does not exist, show option to create a note, Textfield opens with create button, submit sends POST request
  const renderDialogContent = () => {
    if (isEditingNote) {
      return (
        <TextField
          multiline
          label="Add your notes for this movie.."
          fullWidth
          rows={10}
          value={editedNote}
          margin="dense"
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <span style={{ color: editedNote.length > 2000 ? 'red' : 'inherit' }}>
                  {editedNote.length}/{2000}
                </span>
              </InputAdornment>
            ),
          }}
          onChange={(e) => setEditedNote(e.target.value)}
        />
      );
    } else {
      if (selectedNote === null) {
        return renderCreateNoteContent();
      } else if (selectedNote === '') {
        return renderEmptyNoteContent();
      } else {
        return renderExistingNoteContent();
      }
    }
  };

  // Initial display if notes is NULL
  const renderCreateNoteContent = () => (
    <DialogContentText style={{ paddingLeft: '10px', paddingRight: '10px', fontStyle: 'italic', color: 'navy' }}>
      Let's create your first note for this movie!
    </DialogContentText>
  );
  // Initial display if notes is ''
  const renderEmptyNoteContent = () => (
    <DialogContentText style={{ paddingLeft: '10px', paddingRight: '10px', fontStyle: 'italic', color: 'purple' }}>
      Note is empty.. let's add something!
    </DialogContentText>
  );
  
  // Initial display if a note exists with text
  const renderExistingNoteContent = () => (
    <DialogContentText style={{ paddingLeft: '10px', paddingRight: '10px', color:'black' }}>
      {selectedNote}
    </DialogContentText>
  );
  
  // Buttons to accompany dialog boxes based on conditions
  const renderDialogActions = () => {
    if (isCreatingNote) {
      return (
        <DialogActions style={{ paddingBottom: '15px', paddingRight: '18px' }}>
          <Button variant="contained" onClick={handleNoteDialogClose} color="primary">
            Close
          </Button>
          <Button
            variant="contained"
            onClick={handleNoteCreateSubmit}
            color="primary"
            disabled={editedNote.length > 2000}
          >
            Create
          </Button>
        </DialogActions>
      );    
    } else if (isEditingNote) {
      return (
        <DialogActions style={{ paddingBottom: '15px', paddingRight: '18px' }}>
          <Button variant="contained" onClick={handleNoteDialogClose} color="primary">
            Close
          </Button>
          <Button
            variant="contained"
            onClick={handleNoteUpdateSubmit}
            color="primary"
            disabled={editedNote.length > 2000}
          >
            Save
          </Button>
        </DialogActions>
      );
    } else {
      if (selectedNote === null) {
        return renderCreateNoteActions();
      } else {
        return renderEditNoteActions();
      }
    }
  };

  const renderCreateNoteActions = () => (
    <DialogActions style={{ paddingBottom: '15px', paddingRight: '18px' }}>
      <Button variant="contained" onClick={handleNoteDialogClose} color="primary">
        Close
      </Button>
      <Button
        variant="contained"
        onClick={() => {
          setIsEditingNote(true); // Opens text field
          setIsCreatingNote(true); // Renders button to use POST endpoint
          setEditedNote('');
        }}
        color="primary"
      >
        Create Note
      </Button>
    </DialogActions>
  );

  // Appears if a note object is created, regardless if it is empty or not
  const renderEditNoteActions = () => (
    <DialogActions style={{ paddingBottom: '15px', paddingRight: '18px' }}>
      <Button variant="contained" onClick={handleNoteDialogClose} color="primary">
        Close
      </Button>
      <Button
        variant="contained"
        onClick={() => setIsEditingNote(true)}
        color="primary"
      >
        Edit Note
      </Button>
    </DialogActions>
  );

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
        open={openDeletionConfirmation}
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
        <DialogActions style={{ paddingBottom: '20px', paddingRight: '18px' }}>
          <Button variant = "contained" onClick={handleCloseConfirmation}>Cancel</Button>
          <Button variant = "contained" onClick={() => handleConfirmDelete(deleteItemId)} autoFocus>
            Confirm
          </Button>
        </DialogActions>
      </Dialog>

      {/* Dialog for displaying notes */}
      <Dialog 
      open={openNoteDialog} 
      onClose={handleNoteDialogClose}
      maxWidth="md"
      fullWidth={true}
      >
        <DialogTitle style={{ paddingLeft: '30px', paddingTop: '30px', paddingBottom: '18px', display: 'flex', alignItems: 'center'  }}>
          Your Movie Notes
          {selectedNoteUpdatedAt && (
            <Typography variant='body2' color='textSecondary' sx={{ ml: 1, fontStyle: 'italic'}}>
              - Last Updated: {selectedNoteUpdatedAt}
            </Typography>
          )}
        </DialogTitle>

        <DialogContent>
          {renderDialogContent()}

          {dialogNoteErrorMessage && (
            <Typography color="error" variant="body2">
              {dialogNoteErrorMessage}
            </Typography>
          )}
          
        </DialogContent>

        <DialogActions style={{ paddingBottom: '6px', paddingRight: '1px' }}>
          {renderDialogActions()}
        </DialogActions>

      </Dialog>
          </ThemeProvider>
    </div>
  );
};

export default WatchlistItemsTable;