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


import movies_page_img from '../static/Movies-Page.png'; // transparent bar
import SC14 from '../static/Feature5.png'; // white bar
import SC15 from '../static/Feature6.png'; // yellow bar
import movies_search_img from '../static/Movie-Search.png';
import watchlist_img from '../static/Watchlist.png'; 



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
                        <img src={movies_search_img} />
                    </div>

                </div>

                <div className="first-feature">
                    <div className="feature-image">
                        <img src={movies_page_img} />
                    </div>

                    <div className="feature-info">
                        <div class="feature-subtitle">
                            Info On Every Film
                        </div>
                        <div class="feature-headline">
                            Movie Details & Video Trailer
                        </div>
                        <div class="feature-body">
                            Quickly view in-depth details on any movie all in one place. See user ratings, runtime, 
                            revenue, budget, and its trailer. Add to your watchlists once you decide it's a movie
                            worth watching!
                        </div>
                        <div class="feature-info-footer">
                            <a href="#">
                                <button>
                                    Try Search Tool
                                </button>
                            </a>
                        </div>   
                    </div>
                </div>

                <div className="first-feature">
                    <div className="feature-info">
                        <div class="feature-subtitle">
                            Take Notes & Check Off
                        </div>
                        <div class="feature-headline">
                            Create a Watchlist
                        </div>
                        <div class="feature-body">
                            Curate a watchlist and personalize it by mood,
                            genre, actors, or anything that sparks your desire.
                            Add notes to each movie with timestamps.
                            Stay organized by easily checking off when you finish watching a movie.
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
                        <img src={watchlist_img} />
                    </div>
                </div>

            </div>
        </div>
    </div>

  );
}

export default HomeTest;