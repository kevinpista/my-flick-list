import React, { useState, useEffect} from 'react';
import { Container, Paper, Typography, Button } from '@mui/material';
import NavBar from './NavBar';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/MovieSearch.css';
import { movieSearchTMDBAPI } from '../api/movieSearchTMDB';
import { useLocation, useNavigate } from 'react-router-dom';

import MovieSearchBar from './MovieSearchBar.js';
import MovieSearchResults from './MovieSearchResults';

import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';

const MovieSearch = () => {

    const [error, setError] = useState(null);
    const [noMoviesFound, setNoMoviesFound] = useState(false)

    const navigate = useNavigate();

    // Fetch query parameters
    const { search } = useLocation(); // Base code to fetch location of object containing query parameter in form of a string
    const queryParams = new URLSearchParams(search); // Built in JS constructor to work with query paremters

    const query = queryParams.get('query'); // Retrieves value of query parameter called '?query' and so on if more than 1
    const pageParam = (queryParams.get('page')) || 1; // Retrieves value of query parameter called '?query' and so on if more than 1

    const [currentPage, setCurrentPage] = useState(pageParam);
    const [totalPages, setTotalPages] = useState(1);
    const [searchResults, setSearchResults] = useState([]);
    const [totalResultsCount, setTotalResultsCount] = useState(0);

    useEffect(() => {
        if (query !== null) {
        const fetchData = async () => {
            try {
                const response = await movieSearchTMDBAPI(query, currentPage);
                if (response.status === 204) {
                    setNoMoviesFound(true);
                  } else {
                    setSearchResults(response.data.results || []);
                    setTotalPages(response.data.total_pages || 1);
                    setTotalResultsCount(response.data.total_results || 0);
                  }
            } catch (error) {
                setError(error);
            }
        };

        fetchData();
        }
    }, [query, currentPage]);

    // Pagination handling
    const handleNextPage = () => {
        if (currentPage < totalPages) {
            updateURL(currentPage + 1);
        }
    };

    const handlePrevPage = () => {
        if (currentPage > 1) {
            updateURL(currentPage - 1);
        }
    };
    
    const updateURL = (page) => {
        const newSearch = `?query=${query}&page=${page}`;
        navigate(`/movie-search${newSearch}`);
        setCurrentPage(page);
    };

  // Renders No Results message if no movies found by TMDB API
    if (noMoviesFound) {
        return (
            <React.Fragment>
                <NavBar />
                <Container >
                <MovieSearchBar/>

                <ThemeProvider theme={muiTheme}>
                    <Container maxWidth="lg" style={{ marginTop: '15px', textAlign: 'center' }}>
                    <Paper elevation={15} style={{ padding: '25px' }}>
                    <Typography variant="h6" color='#032541' fontWeight='bold'>
                        No Results Found
                    </Typography>

                    <Typography variant="h7" >
                    There were no movies found for your search query.
                    </Typography>
                    </Paper>
                    </Container>
                </ThemeProvider>

                </Container>
            </React.Fragment>
        );
    };
        
    // RENDER COMPONENT
    return (
        <React.Fragment>
            <NavBar />
            <Container >
            <MovieSearchBar/>
            {query === null ? (
                    // Render a message when query is null
                        <ThemeProvider theme={muiTheme}>
                          <Container maxWidth="lg" style={{ marginTop: '15px', textAlign: 'center' }}>
                          <Paper elevation={15} style={{ padding: '25px' }}>
                            <Typography variant="h6" color='#032541' fontWeight='bold'>
                              Use the search bar above to find any movie
                            </Typography>

                            <Typography variant="h7" >
                              Then view the movie and add it to your watchlist.
                            </Typography>
                          </Paper>
                          </Container>
                        </ThemeProvider>
            ) : (
                error ? (
                    <h1 className='error'><u>Error:</u> {error.message}</h1>
                ) : (
                    <React.Fragment>
                        {searchResults && (
                        searchResults.map((movie) => (
                            <MovieSearchResults
                                key = {movie.id}
                                id = {movie.id}
                                title = {movie.original_title}
                                releaseDate = {movie.release_date}
                                description = {movie.overview}
                                posterURL={movie.poster_path ? `https://image.tmdb.org/t/p/w300_and_h450_bestv2${movie.poster_path}` : null }
                                />
                            )) 
                        )}
                        <div className="pagination">
                            <Button onClick={handlePrevPage} disabled={currentPage === 1}>
                                Previous
                            </Button>
                            <Typography variant="h6" className="page-indicator" style={{ margin: '0 15px' }}>
                                Page {currentPage} of {totalPages}
                            </Typography>
                            <Button onClick={handleNextPage} disabled={currentPage === totalPages}>
                                Next Page
                            </Button>
                            <Typography variant="h6" className="page-indicator" style={{ margin: '0 15px' }}>
                            Total results: {totalResultsCount}
                            </Typography>
                        </div>
                    </React.Fragment>
                )
                )}
            </Container>
        </React.Fragment>
  );
};
export default MovieSearch;