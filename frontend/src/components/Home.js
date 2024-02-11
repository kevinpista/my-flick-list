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
import Square from '../static/Square.jpg';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';

import SC6 from '../static/first-image.png';


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
        <div class="features-headline">

            <h2>Over 850,000 Movies</h2>
            <p>Search for nearly every movie that was ever made. All powered by 'The Movie Database' (TMDB).</p>

        </div>
      </Container>



      <Container class="features">
      <div class="features-row">
        <div class="text">
          <h2>Search for Any Movie</h2>
          <p>
            Use the search function powered by TMDB API. See detailed overviews, 
            release dates, ratings, and finacials for each movie
          </p>
        </div>
        <div class="image"><img src={SC6} alt="movie page"/></div>
      </div>
      <div class="features-row">
      <div class="image"><img src={SC6} alt="list of watchlist page" /></div>

        <div class="text">
          <h2>Create Your Watchlist</h2>
          <p>Curate a watchlist and add any movie. Personalize by your mood, genre,
            favorite director, actors, or anything that sparks you desire to group your movies into 
            one place.
             </p>
        </div>
      </div>
      <div class="features-row">
        <div class="text">
          <h2>Take Notes & Check Off</h2>
          <p>Jot down your own notes about each movie before or after you watch. Organize by checkmarking 
            whenever you've finished watching a movie. 
          </p>
        </div>
        <div class="image"><img src={SC6}  alt="watchlist page" /></div>
      </div>
        </Container>
    </div>
    </ThemeProvider>
  );
}

export default Home;