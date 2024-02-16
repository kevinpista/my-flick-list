import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Alert from '@mui/material/Alert';
import Snackbar from '@mui/material/Snackbar';
import Container from 'react-bootstrap/Container';

import { removeTokenFromCookie, getJwtTokenFromCookies } from '../utils/authTokenUtils';

import { ThemeProvider } from '@mui/material/styles';
import {muiTheme} from '../css/MuiThemeProvider.js';

// AppBar imports
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Tooltip from '@mui/material/Tooltip';
import SearchIcon from '@mui/icons-material/Search';
import TheatersOutlinedIcon from '@mui/icons-material/TheatersOutlined';


function NavBar() {
  // Logout alert message and navigation
  const navigate = useNavigate();
  const [snackbarOpen, setSnackbarOpen] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(getJwtTokenFromCookies() !== undefined);

  const handleLogout = () => {
    removeTokenFromCookie();
    setSnackbarOpen(true);
    setTimeout(() => {
      if (window.location.pathname === '/') {
        window.location.reload();
      } else {
        navigate('/');
      }
    }, 1500); // Redirect after 1.5 seconds delay
  };

  const handleCloseSnackbar = () => {
    setSnackbarOpen(false);
  };

  return (
    <ThemeProvider theme={muiTheme}>
    <AppBar position="static" sx={{marginTop:'-40px'}}>
      <Container maxWidth="xl" style={{ marginTop: '50px' }}>
        <Toolbar disableGutters >
          {/* Logo Icon + Name */}
          <TheatersOutlinedIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1.5, fontSize: '30px' }} />
          <Typography
            variant="h6"
            noWrap
            component="a"
            href="/"
            sx={{
              mr: 2,
              display: { xs: 'none', md: 'flex' },
              fontFamily: 'monospace',
              fontWeight: 700,
              letterSpacing: '.05rem',
              color: 'inherit',
              textDecoration: 'none',
            }}
          >
            My Flick List
          </Typography>

          {/* Logo Icon + Name for small screens */}
          <TheatersOutlinedIcon sx={{ display: { xs: 'flex', md: 'none' }, mr: 1.5, fontSize: '30px' }} />
          <Typography
            variant="h5"
            noWrap
            component="a"
            href="/"
            sx={{
              mr: 2,
              display: { xs: 'flex', md: 'none' },
              flexGrow: 1,
              fontFamily: 'monospace',
              fontWeight: 700,
              letterSpacing: '.05rem',
              color: 'inherit',
              textDecoration: 'none',
            }}
          >
            My Flick List
          </Typography>

          {/* Left side of nav */}
          <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
            <Button
              onClick={() => navigate('/movie-search')}
              sx={{ my: 2, color: 'white', display: 'block' }}
            >
              Movies
            </Button>
            <Button
              onClick={() => navigate('/watchlist')}
              sx={{ my: 2, color: 'white', display: 'block' }}
            >
              Watchlists
            </Button>
            <Button
              onClick={() => navigate('/about')}
              sx={{ my: 2, color: 'white', display: 'block' }}
            >
              About
            </Button>

            <Button
              onClick={() => { window.open('https://github.com/kevinpista/my-flick-list', '_blank'); }}
              sx={{ my: 2, color: 'white', display: 'block' }}
            >
              Git Repo
            </Button>
            {!isLoggedIn &&
              <Button
                onClick={() => navigate('/user-login')}
                sx={{ my: 2, color: 'white', display: 'block' }}
              >
                Demo
              </Button>
            }
          </Box>

          {/* Right side of nav */}
          <Box sx={{ flexGrow: 0, display: { xs: 'none', md: 'flex' } }}>
            {isLoggedIn ? (
              <Button
                onClick={handleLogout}
                sx={{ my: 2, color: 'white', display: 'block', marginRight: 1 }}
              >
                Logout
              </Button>
              ) : (
              <Button
                onClick={() => navigate('/user-login')}
                sx={{ my: 2, color: 'white', display: 'block', marginRight: 2 }}
              >
                Login
              </Button>
            )}
            <Tooltip title="Search for a movie">
              <IconButton onClick={() => navigate('/movie-search')} sx={{ p: 0 }}>
                <SearchIcon sx={{ my: 2, color: 'white', display: 'block', fontSize: '30px'}}
                />
              </IconButton>
            </Tooltip>
          </Box>

        </Toolbar>
      </Container>
    </AppBar>

    {/* Alerts for when user clicks LOGOUT */}
    <Snackbar
      open={snackbarOpen}
      autoHideDuration={2000}
      onClose={handleCloseSnackbar}
      anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
      style={{ top: '50px' }}
    >
      <Alert
        severity='info'
      >
        Logging you out ...
      </Alert>
    </Snackbar>
    </ThemeProvider>
  );
}
export default NavBar;
