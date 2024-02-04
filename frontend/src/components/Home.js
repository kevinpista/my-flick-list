// Home.js 
import { Container, Row, Col } from 'react-bootstrap';
import Button from '@mui/material/Button';
import { useNavigate } from 'react-router-dom';
import '../css/Home.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import NavBar from './NavBar';
import MovieSearchBar from './MovieSearchBar.js';
import Rectangle_Movies from '../static/Rectangle_Movies.jpg';
import Sqaure_Movies from '../static/Square_Movies.jpg';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';

function Home() {
  const navigate = useNavigate();

  const handleGetStartedClick = () => {
    navigate('/user-registration')
  };

  return (
    <ThemeProvider theme={muiTheme}>
    <div className="home">
      <NavBar />
      <div 
        className="jumbotron"
        style={{ backgroundImage: `url(${Rectangle_Movies})` }} 
      >
        <div className="gradient-overlay">
          <div className="banner-details">
          <h1>Curate. Watch. Repeat.</h1>
          <p>
            Create endless watchlists of your favorite movies, 
            add notes, and track when you  finish watching.
          </p>
          <Button
            variant="contained" 
            onClick={handleGetStartedClick}
            back="primary" // Use bg instead of color
            sx={{
            boxShadow: `0px 0px 2px rgba(255, 255, 255, 1),
                        0px 0px 6px rgba(255, 255, 255, 0.1),
                        0px 0px 10px rgba(255, 255, 255, 0.05)`,   
            border: `1px solid rgba(255, 255, 255, 0.25)`,
            }}
          >
            Get Started
          </Button>
          </div>
        </div>
      </div>

      {/* React Bootstrap */}
      <Container style={{ marginTop: '20px' }}>
        <MovieSearchBar />
        <Row>
          <Col md={4}>
            <h2>Search</h2>
            <p>Search for new movies to add to your watchlist.</p>
          </Col>

          <Col md={4}>
            <h2>Watchlist</h2>
            <p>View and manage your personal watchlist.</p>
          </Col>

          <Col md={4}>
            <h2>History</h2>
            <p>See what movies you've already watched.</p>
          </Col>
        </Row>
      </Container>

    </div>
    </ThemeProvider>
  );
}

export default Home;