import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
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
    }, 2000); // Redirect after 2 seconds delay
  };

  const handleCloseSnackbar = () => {
    setSnackbarOpen(false);
  };

  return (
    <>
    <Navbar data-bs-theme="dark" expand="lg" className="nav-bar-background">
      <Container className="navbar-container">
        <Navbar.Brand href="/" className="custom-font">My Flick List</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
                <Nav.Link href="/movie-search" className="custom-font">Movie Search</Nav.Link>
                <Nav.Link href="/watchlist" className="custom-font">Watchlists</Nav.Link>
                <Nav.Link href="user-login" className="custom-font">Login</Nav.Link>
                <Nav.Link onClick={handleLogout} className="custom-font">Logout</Nav.Link>
            </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
    <Snackbar
      open={snackbarOpen}
      autoHideDuration={3000}
      onClose={handleCloseSnackbar}
      anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
      style={{ top: '50px' }}
    >
      <Alert
          severity="success"
      >
        Logout successful...
      </Alert>
</Snackbar>
</>
  );
}
export default NavBar;
