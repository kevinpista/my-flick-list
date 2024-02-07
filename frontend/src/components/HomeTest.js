// Home.js 
import { Container, Row, Col } from 'react-bootstrap';
import Button from '@mui/material/Button';
import { useNavigate } from 'react-router-dom';
import '../css/HomeTest.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import NavBar from './NavBar';
import MovieSearchBar from './MovieSearchBar.js';
import Rectangle_Movies from '../static/Rectangle_Movies.jpg';
import Sqaure_Movies from '../static/Square_Movies.jpg';
import Square from '../static/Square.jpg';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';
import Parallax from './Parallax.js';


import SC1 from '../static/Showcase_1_big.jpg';
import SC2 from '../static/Showcase_2_big.jpg';
import SC3 from '../static/Showcase_3_small.jpg';
import SC4 from '../static/Showcase_11_small.jpg';
import SC5 from '../static/Showcase_12_small.jpg';
import SC6 from '../static/Showcase_13_small.jpg';
import SC7 from '../static/first-image.png';


function HomeTest() {
  const navigate = useNavigate();

  const handleGetStartedClick = () => {
    navigate('/user-registration')
  };

  return (
    <div className="home">
      <NavBar />
      <Parallax />
        <div className="features-container">
            <div className="features">

                <div className="first-feature">
                    <div className="feature-info">
                        <div class="feature-subtitle">
                            Over 850,000 Titles
                        </div>
                        <div class="feature-headline">
                            Search for Any Movie
                        </div>
                        <div class="feature-body">
                            Find nearly every movie that was ever made.
                            Discover spin-offs and highly acclaimed fan made movies of your favorite flicks that you
                            didn't know existed. All powered by 'The Movie Database' (TMDB).
                        </div>
                        <div class="feature-info-footer">
                            <a href="#">
                                <button>
                                    Find a Movie
                                </button>
                            </a>
                        </div>   
                    </div>
            
                    <div className="feature-image">
                        <img src={SC7} />
                    </div>

                </div>

                <div className="first-feature">
                    <div className="feature-image">
                        <img src={SC7} />
                    </div>

                    <div className="feature-info">
                        <div class="feature-subtitle">
                            Make It Yours
                        </div>
                        <div class="feature-headline">
                            Create Your Watchlist
                        </div>
                        <div class="feature-body">
                            Curate a watchlist and add any movie to it. Personalize by mood,
                            genre, favorite director, actors, or anything that sparks your desire
                            to organize all your movies into one place.
                        </div>
                        <div class="feature-info-footer">
                            <a href="#">
                                <button>
                                    Get Started
                                </button>
                            </a>
                        </div>   
                    </div>
                </div>

                <div className="first-feature">
                    <div className="feature-info">
                        <div class="feature-subtitle">
                            Jot Down Your Thoughts
                        </div>
                        <div class="feature-headline">
                            Take Notes & Check Off
                        </div>
                        <div class="feature-body">
                            Add notes to each movie before or after you watch with timestamps on when you
                            last updated them. Stay organized by easily checking off when you finished watching a movie.
                        </div>
                        <div class="feature-info-footer">
                            <a href="#">
                                <button>
                                    Personalize Now
                                </button>
                            </a>
                        </div>   
                    </div>
                    <div className="feature-image">
                        <img src={SC7} />
                    </div>
                </div>

            </div>
        </div>
    </div>

  );
}

export default HomeTest;