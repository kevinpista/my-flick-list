import React from 'react';
import { Container, Paper, Typography, Button, Rating } from '@mui/material';
import NavBar from './NavBar';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/Movie.css'; // Import the CSS file

const MoviePage = () => {
  return (
    <React.Fragment>
        <NavBar />
        <Container>
            <Paper elevation={3} className="movie-paper"> 
            <Typography variant="h4" gutterBottom className="movie-title"> 
                Movie Title
            </Typography>
            <Typography variant="body1" className="movie-description">  
                Description of the movie goes here. Lorem ipsum dolor sit amet,
                consectetur adipiscing elit.
            </Typography>
            <Typography variant="body2" className="movie-release-date"> 
                Release Date: January 1, 2023
            </Typography>
            <Typography variant="body2" className="movie-ratings"> 
                Ratings:
            </Typography>
            <Rating value={4.5} precision={0.5} readOnly className="movie-rating" /> 
            <Button variant="contained" color="primary" className="watch-button"> 
                Watch Trailer
            </Button>
            </Paper>
        </Container>
    </React.Fragment>

  );
};

export default MoviePage;
