// Home.js 
import '../css/HomeTest.css';
import '../css/Footer.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import NavBar from './NavBar';
import Parallax from './Parallax.js';


import movies_page_img from '../static/Movies-Page.png'; // transparent bar
import SC14 from '../static/Feature5.png'; // white bar
import SC15 from '../static/Feature6.png'; // yellow bar
import movies_search_img from '../static/Movie-Search.png';
import watchlist_img from '../static/Watchlist.png'; 
import arrow_right from '../static/arrow-right.png'; 
import tmdb_logo from '../static/tmdb-short-logo.svg';
import tmdb_logo_long from '../static/tmdb-logo-long.svg';

import Typography from '@mui/material/Typography';
import TheatersOutlinedIcon from '@mui/icons-material/TheatersOutlined';


function HomeTest() {

  return (
    <div className="home">
      <NavBar />
      <Parallax />
        <div className="features-container">
            <div className="features">

                <div className="feature">
                    <div className="feature-info">
                        <div class="feature-subtitle">
                            Over 850,000 Titles
                        </div>
                        <div class="feature-headline">
                            Search Any Movie
                        </div>
                        <div class="feature-body">
                            Find nearly every movie that was ever made.
                            Discover spin-offs and highly acclaimed fan made movies of your favorite flicks that you
                            didn't know existed. All powered by
                            <span style={{marginLeft: '10px'}}>
                                <a href="/about">
                                <img src={tmdb_logo_long } />.
                                </a>
                            </span>
                            <div class="number-step-bg" id="first-number-step">01</div>

                        </div>
                        <div class="feature-info-footer">
                            <a href="/movie-search">
                                <button>
                                    Find a Movie
                                    <img src={arrow_right} />
                                </button>
                            </a>
                            
                        </div>   
                    </div>
            
                    <div className="feature-image">
                        <img src={movies_search_img} />
                    </div>

                </div>

                <div className="feature">
                    <div className="feature-image">
                        <img src={movies_page_img} />
                    </div>

                    <div className="feature-info">
                        <div class="feature-subtitle">
                            Info On Every Film
                        </div>
                        <div class="feature-headline">
                            Movie Details
                        </div>
                        <div class="feature-body">
                            Quickly view in-depth details on any movie all in one place. See user ratings, runtime, 
                            revenue, budget, and its trailer. Add to your watchlists once you decide it's a movie
                            worth watching!
                            <div class="number-step-bg" id="second-number-step">02</div>
                        </div>
                        <div class="feature-info-footer">
                            <a href="/movie-search">
                                <button>
                                    Try a Search
                                    <img src={arrow_right} />
                                </button>
                            </a>
                        </div>   
                    </div>
                </div>

                <div className="feature">
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
                            Stay organized by easily checking off when you finish a movie.
                            <div class="number-step-bg" id="first-number-step">03</div>
                        </div>
                        <div class="feature-info-footer">
                            <a href="/user-login">
                                <button>
                                    Personalize
                                    <img src={arrow_right} />
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

    <footer>
        <div class="footer">
        <div class="footer-section-1">
            <div className="footer-logo">
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
                textDecoration: 'none',
                color: 'white',
                }}
            >
                My Flick List
            </Typography>
            </div>

            <div className="footer-body">
            Personalize & curate endless watchlists of your favorite movies.
            </div>

          <div className="footer-privacy">
            <a href="#"> Copyright 2024 My Flick List, Inc. Terms & Privacy</a>
          </div>
        </div>
        <div className="lists">
            <div className="footer-navigation">
                <div className="footer-navigation-header">
                    <a href="#footer">Navigation Menu</a>
                </div>
                <ul>
                    <li><a href="/movie-search">Movie Search</a></li>
                    <li><a href="/watchlist">Watchlists</a></li>
                    <li><a href="/about">About</a></li>
                    <li><a href="https://github.com/kevinpista/my-flick-list" target="_blank">Git Repo</a></li>
                </ul>
            </div>

            <div class="footer-section-2">
                <div className="tmdb-footer-logo">
                {/* Logo Icon + Name */}
                <img href="https://www.themoviedb.org/" target="_blank" src={tmdb_logo_long } />
                </div>
                <div className="tmdb-footer-body">
                    This web app uses TMDb API but is not endorsed or certified by TMDb.
                </div>
            </div>
        </div>
        </div>
    </footer>
    </div>

  );
}

export default HomeTest;