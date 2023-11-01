import React from 'react';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import '../css/NavBar.css'; // Import the CSS file


// TODO - make height smaller, Add Logo
function NavBar() {
  return (
    <Navbar data-bs-theme="dark" expand="lg" className="nav-bar-background">
      <Container className="navbar-container">
        <Navbar.Brand href="#home" className="custom-font">My Flick List</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
                <Nav.Link href="#features" className="custom-font">Movie Search</Nav.Link>
                <Nav.Link href="#pricing" className="custom-font">Watchlists</Nav.Link>
                <Nav.Link href="#pricing" className="custom-font">Account</Nav.Link>
            </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}
export default NavBar;
