import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Alert from '@mui/material/Alert';
import Snackbar from '@mui/material/Snackbar';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import '../css/NavBar.css'; // Import the CSS file
import { removeTokenFromCookie, getJwtTokenFromCookies } from '../utils/authTokenUtils';

import { ThemeProvider } from '@mui/material/styles';
import {muiTheme} from '../css/MuiThemeProvider.js';

// AppBar imports
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import MenuIcon from '@mui/icons-material/Menu';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import Tooltip from '@mui/material/Tooltip';
import MenuItem from '@mui/material/MenuItem';
import AdbIcon from '@mui/icons-material/Adb';

const pages = ['Movies', 'Watchlists', 'Logout'];

// TODO - make height smaller, Add Logo
function NavBar() {
  // Nav bar itself
  const [anchorElNav, setAnchorElNav] = useState(null);
  const [anchorElUser, setAnchorElUser] = useState(null);

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };
  const handleOpenUserMenu = (event) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  // Messaging and redirect
  const navigate = useNavigate();
  // const handleClick = () => navigate('/');
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

    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          {/* Logo Text + Icon*/}
          <AdbIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
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
              letterSpacing: '.1rem',
              color: 'inherit',
              textDecoration: 'none',
            }}
          >
            My Flick List
          </Typography>

          <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleOpenNavMenu}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: 'bottom',
                horizontal: 'left',
              }}
              keepMounted
              transformOrigin={{
                vertical: 'top',
                horizontal: 'left',
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{
                display: { xs: 'block', md: 'none' },
              }}
            >
              {pages.map((page) => (
                <MenuItem key={page} onClick={handleCloseNavMenu}>
                  <Typography textAlign="center">{page}</Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>
          <AdbIcon sx={{ display: { xs: 'flex', md: 'none' }, mr: 1 }} />
          <Typography
            variant="h5"
            noWrap
            component="a"
            href="#app-bar-with-responsive-menu"
            sx={{
              mr: 2,
              display: { xs: 'flex', md: 'none' },
              flexGrow: 1,
              fontFamily: 'monospace',
              fontWeight: 700,
              letterSpacing: '.3rem',
              color: 'inherit',
              textDecoration: 'none',
            }}
          >
            LOGO
          </Typography>
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
            {isLoggedIn ? (
              <Button
                onClick={handleLogout}
                sx={{ my: 2, color: 'white', display: 'block' }}
              >
                Logout
              </Button>
              ) : (
              <Button
                onClick={() => navigate('/user-login')}
                sx={{ my: 2, color: 'white', display: 'block' }}
              >
                Login
              </Button>
            )}
          </Box>
          </Toolbar>
      </Container>
      
    </AppBar>

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
