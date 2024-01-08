import React, { useState } from 'react';
import { DataGrid} from '@mui/x-data-grid';
import Link from '@mui/material/Link';

// MUI Dialog component to confirm watchlist deletion
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';

import ClearSharpIcon from '@mui/icons-material/ClearSharp';
import { ThemeProvider } from '@mui/material/styles';
import * as themeStyles from '../../styling/ThemeStyles';


// watchlistData is a JSON object holding 1 array containing individual watchlists
const ListOfWatchlistsTable = ({ watchlistData, onDeleteWatchlist, setWatchlistData }) => {
  // Dialog component for deleting a watchlist
  const [openConfirmation, setOpenConfirmation] = useState(false);
  const [deleteWatchlistId, setDeleteWatchlistId] = useState(null);    

  const [editRowsModel, setEditRowsModel] = useState({});

  
  const handleDeleteClick = (event, watchlistId) => {
    event.stopPropagation();
    setDeleteWatchlistId(watchlistId);
    setOpenConfirmation(true);
  };

  const handleCloseConfirmation = () => {
    setOpenConfirmation(false);
    setDeleteWatchlistId(null);
  };

  const handleConfirmDelete = (watchlistId) => {
    handleCloseConfirmation();
    onDeleteWatchlist(watchlistId); // Call this function with the watchlist ID to be deleted
  };
  

  const getRowId = (row) => row.id;

  // Set rows to be sorted by most recently updated by default
  const [sortModel, setSortModel] = React.useState([
    {
      field: 'updated_at',
      sort: 'desc',
    },
  ]);

  // Make description expandable if it overflows 150 chars
  function ExpandableCell({ value }) {
    const [expanded, setExpanded] = useState(false);
    return (
      <div>
        {expanded ? value : value.slice(0, 150)}&nbsp;
        {value.length > 150 && (
          <Link
            type="button"
            component="button"
            sx={{ fontSize: 'inherit' }}
            onClick={() => setExpanded(!expanded)}
          >
            {expanded ? 'view less' : 'view more'}
          </Link>
        )}
      </div>
    );
  }

  const columns = [
    { 
        field: 'name', 
        headerName: 'Watchlist Name', 
        width: 400, 
        headerAlign: 'left', 
        align: 'left',
        renderCell: (params) => (
            <Link href={`/watchlist/${params.row.id}`} underline="hover">
            {params.value}
          </Link>
        )
     },
    { 
        field: 'description', 
        headerName: 'Description', 
        width: 600, 
        headerAlign: 'left', 
        align: 'left', 
        renderCell: (params) => <ExpandableCell {...params} />,
    },
    { 
        field: 'updated_at', 
        headerName: 'Last Updated', 
        width: 160, 
        headerAlign: 'center', 
        align: 'center',
        valueFormatter: (params) => {
            const date = new Date(params.value);
            return date.toLocaleDateString('en-US', { month: 'long', day: 'numeric', year: 'numeric' });
        },
     },
    { 
        field: 'created_at', 
        headerName: 'Created', 
        width: 160, 
        headerAlign: 'center', 
        align: 'center' ,
        valueFormatter: (params) => {
            const date = new Date(params.value);
            return date.toLocaleDateString('en-US', { month: 'long', day: 'numeric', year: 'numeric' });
        },
    },
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
  
  const rows = watchlistData['watchlists'].map((watchlist) => ({
    id: watchlist.id, // watchlist id
    name: watchlist.name,
    description: watchlist.description,
    updated_at: new Date(watchlist.updated_at), // Set as ISO string first. Use valueFormatter to convert into Mon, DD, YYYY format
    created_at: new Date(watchlist.created_at),
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
        sortModel={sortModel}
        onSortModelChange={(model) => setSortModel(model)}
        autoHeight={true}
        disableMultipleRowSelection={true}
        getRowId={getRowId}
        getRowHeight={() => 'auto'}
        sx={{
            '&.MuiDataGrid-root--densityCompact .MuiDataGrid-cell': { py: '8px' },
            '&.MuiDataGrid-root--densityStandard .MuiDataGrid-cell': { py: '15px' },
            '&.MuiDataGrid-root--densityComfortable .MuiDataGrid-cell': { py: '22px' },
          }}
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
            Are you sure you want to delete this watchlist?
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button variant="contained" onClick={handleCloseConfirmation}>Cancel</Button>
          <Button variant="contained" onClick={() => handleConfirmDelete(deleteWatchlistId)} autoFocus>
            Confirm
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default ListOfWatchlistsTable;