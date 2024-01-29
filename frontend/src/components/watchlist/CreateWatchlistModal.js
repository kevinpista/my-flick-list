import React, { useState } from 'react';
import { Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, Typography } from '@mui/material';
import LoadingButton from '@mui/lab/LoadingButton';
import InputAdornment from '@mui/material/InputAdornment';
import { createWatchlistAPI } from '../../api/watchlistAPI.js'
import * as errorConstants from '../../api/errorConstants';

import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../../css/MuiThemeProvider.js';


const CreateWatchlistModal = ({open, onClose}) => {
  
    const [newWatchlistName, setNewWatchlistName] = useState('');
    const [newWatchlistDescription, setNewWatchlistDescription] = useState('');
  
    const [dialogErrorMessage, setDialogErrorMessage] = useState('');
    const [loading, setLoading] = useState(false);


    const handleCreateWatchlistButtonClose = () => {
        onClose(); // This calls the prop function, which is setOpenCreateWatchlistModal
    };
    // createWatchlistAPI Call
    const handleCreateWatchlistDialogSubmit = async () => {
        setLoading(true)
        try {
        const response = await createWatchlistAPI(newWatchlistName, newWatchlistDescription);
        if (response) {

            console.log('it works')
        }
        } catch (error) {
        if (error.message === errorConstants.ERROR_BAD_REQUEST) {
            setLoading(false)
            setDialogErrorMessage('Error: Bad request. Please try again.');
        } else {
            setLoading(false)
            setDialogErrorMessage(`Error: ${error.message}`);
        } 
        }
    };
    return (
        <ThemeProvider theme={muiTheme}>
        {/* Modal creating a watchlist */}
        <Dialog
        open={open}
        onClose={onClose}
        maxWidth="md"
        fullWidth={true}
        >
        <DialogTitle><b>Create a New Watchlist</b></DialogTitle>
        <DialogContent>
            <TextField
            autoFocus
            id="watchlist-name"
            label="Watchlist Name"
            value={newWatchlistName}
            onChange={(e) => setNewWatchlistName(e.target.value)}
            multiline
            fullWidth
            margin="dense"
            variant="standard"
            // Display character limit and changes text to red if user goes over limit
            InputProps={{
                endAdornment: (
                <InputAdornment position="end">
                <span style={{ color: newWatchlistName.length > 60 ? 'red' : 'inherit' }}>
                    {newWatchlistName.length}/{60}
                </span>
                </InputAdornment>
                ),
            }}
            />
            <TextField
            autoFocus
            id="watchlist-description"
            label="Watchlist Description"
            value={newWatchlistDescription}
            onChange={(e) => setNewWatchlistDescription(e.target.value)}
            multiline
            fullWidth
            margin="dense"
            variant="standard"
            // Display character limit and changes text to red if user goes over limit
            InputProps={{
                endAdornment: (
                <InputAdornment position="end">
                <span style={{ color: newWatchlistDescription.length > 500 ? 'red' : 'inherit' }}>
                    {newWatchlistDescription.length}/{500}
                </span>
                </InputAdornment>
                ),
            }}
            />
            {dialogErrorMessage && (
            <Typography color="error" variant="body2">
                {dialogErrorMessage}
            </Typography>
            )}
        </DialogContent>
        <DialogActions style={{ paddingBottom: '20px', paddingRight: '18px' }}>
            <Button variant="contained" onClick={handleCreateWatchlistButtonClose}>
            Exit</Button>

            <LoadingButton 
            variant="contained"
            loading={loading}
            onClick={handleCreateWatchlistDialogSubmit}
            disabled={
            newWatchlistName.length > 60 || // Character limit for watchlist name
            newWatchlistDescription.length > 500 // Character limit for watchlist description
            }
            >
            Create
            </LoadingButton>

        </DialogActions>
        </Dialog>
        </ThemeProvider>
        );
};
export default CreateWatchlistModal;
