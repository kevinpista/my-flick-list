import React from 'react';
import { Container, Paper, Typography, Button } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import NavBar from './NavBar';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/Movie.css'; // Import the CSS file

// TODO
// 1. Change AddIcon button and icon to "Added" with CheckMark icon when successfully added to someone's watchlist
// popup menu of which watchlist to add to -- or dropdown to let them pick which watchlist they want to add to.

// 2. Make revenue and budget side by side with nicer icons. Unit converion to k, or m with 1 decimal place

// 3. Move the moive poster image more to the right along with the movie content accordingly-; padding / flex adjustments

const Test = () => {
  return (
    <React.Fragment>
      <NavBar />
      <Container maxWidth="fluid">
        <Paper elevation={3} className="movie-paper">

              <div className="movie-poster">
                  <img class="poster-small" src="https://image.tmdb.org/t/p/w300_and_h450_bestv2/c54HpQmuwXjHq2C9wmoACjxoom3.jpg" alt="Movie Poster" />
              </div>

              <div className="movie-details">
                  <Typography variant="h3" className="movie-title" fontWeight="bold">
                    Harry Potter and the Deathly Hallows: Part 2 (2011)              
                    </Typography>
                    <div className="movie-description">
                    <Typography variant="body4">
                    January 1, 2023 | Fantasy, Adventure | 2h 10m              
                    </Typography>
                    </div>

                  <div className="movie-ratings">
                  <Typography variant="body3" >Ratings: 7.4/10 // 12.1k votes</Typography>
                  </div>

                  <Typography variant="body4" gutterBottom className="movie-tagline">It all ends here.</Typography>
                    <Typography variant="h5" className="movie-description">
                      Overview
                    </Typography>
                    <Typography variant="body1" paragraph>
                    Harry, Ron and Hermione continue their quest to vanquish the evil Voldemort once and for all. Just as things begin to look hopeless for the young wizards, Harry discovers a trio of magical objects that endow him with powers to rival Voldemort's formidable skills.
                    </Typography>

                    <div className="movie-financials">
                  <Typography variant="body1" gutterBottom>Revenue: $10m // Budget: $6m</Typography>
                  </div>
                  <Button 
                      variant="contained"
                      color="primary"
                      size="large" 
                      className="add-to-watchlist-btn"
                      endIcon={<AddIcon />}
                    >
                      ADD TO WATCHLIST
                    </Button>
          </div>
        </Paper>
      </Container>
    </React.Fragment>
  );
};
export default Test;
