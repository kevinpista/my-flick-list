

import tmdb_logo from '../static/tmdb-short-logo.svg';
import tmdb_logo_long from '../static/tmdb-logo-long.svg';

import Typography from '@mui/material/Typography';
import TheatersOutlinedIcon from '@mui/icons-material/TheatersOutlined';

import '../css/Footer.css';

function Footer() {
    return (
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
    );
}

export default Footer;
