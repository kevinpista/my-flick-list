import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { NavLink } from 'react-router-dom';
import Alert from '@mui/material/Alert';
import Snackbar from '@mui/material/Snackbar';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import '../css/NavBar.css'; // Import the CSS file
import { removeTokenFromCookie } from '../utils/authTokenUtils';


// TODO - make height smaller, Add Logo
function NavBar() {

  const navigate = useNavigate();
  // const handleClick = () => navigate('/');
  const [snackbarOpen, setSnackbarOpen] = useState(false);

  const handleLogout = () => {
    removeTokenFromCookie();
    setSnackbarOpen(true);
    setTimeout(() => {
      navigate('/');
    }, 1500); // Redirect after 1.5 seconds delay
  };

  const handleCloseSnackbar = () => {
    setSnackbarOpen(false);
  };

  return (
    <>
    <Navbar data-bs-theme='dark' expand='lg' className='nav-bar-background'>
      <Container className='navbar-container'>
        <Navbar.Brand onClick={() => navigate('/')} className='custom-font'>
          My Flick List
        </Navbar.Brand>
        <Navbar.Toggle aria-controls='basic-navbar-nav' />
        <Navbar.Collapse id='basic-navbar-nav'>
            <Nav className='me-auto'>
              <Nav.Link onClick={() => navigate('/movie-search')} className='custom-font'>
                Movie Search
              </Nav.Link>
              <Nav.Link onClick={() => navigate('/watchlist')} className='custom-font'>
                Watchlists
              </Nav.Link>
              <Nav.Link onClick={() => navigate('/user-login')} className='custom-font'>
                Login
              </Nav.Link>
              <Nav.Link onClick={handleLogout} className='custom-font'>
                Logout
              </Nav.Link>                
            </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
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
        Logging you out...
      </Alert>
</Snackbar>
</>
  );
}
export default NavBar;
