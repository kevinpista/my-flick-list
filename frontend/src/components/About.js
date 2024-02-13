import '../css/About.css';
import NavBar from './NavBar.js';
import Footer from './Footer.js';
import React from 'react';

import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';
import tmdb_logo_long from '../static/tmdb-logo-long.svg';


const About = () => {
    return (
        <ThemeProvider theme={muiTheme}>
        <NavBar />
        <div className="about-wrapper">
        <div className="about-container">
            <div class="title-wrapper">
                <h1>About</h1>        
            </div>
            <div className="about-text">
                <p>
                    <b>'My Flick List'</b> is a web app created as a personal project. The web app allows a user to search for any movie
                    thanks to the data provied by <b>'TMDb API'</b>.
                    Every movie will contain information such as the title, overview, release date, trailer video,
                    along with other metrics such as its budget and revenue generated in the box office. A user can create a personalized watchlist
                    and add any movie to it. Within each watchlist, the user can check off whenever they finish watching a movie and also write their notes for each film.
                    These functions allow a cinema enthusiast to gather all their movies and thoughts in one place in order to make their watching
                    experience more organized and enjoyable.
                </p>
                <p>
                    I used Go with a PostgreSQL database for the backend. For the frontend, I used React.js along with the Material UI React component library.
                </p>
                <p>
                    <a href="https://github.com/kevinpista/my-flick-list" target="_blank">
                        Github Repo Here
                    </a>
                </p>

                <h2>
                    TMDb API Attribution
                </h2>
                <p>
                    This personal project uses TMDb API but is not endorsed or certified by TMDb.
                </p>
                <p>
                    <a href="https://www.themoviedb.org/" target="_blank">
                        <img src={tmdb_logo_long } />
                    </a>
                </p>
            </div>
        </div>
        <div className="home">
            <Footer />
        </div>
        </div>
        </ThemeProvider>
  );
};

export default About;