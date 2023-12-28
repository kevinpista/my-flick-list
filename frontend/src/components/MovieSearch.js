import React, { useState, useEffect} from 'react';
import { Container, Paper, Typography, Button } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import NavBar from './NavBar';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/MovieSearch.css';
import { movieSearchTMDB } from '../api/movieSearchTMDB';
import { useParams, useLocation } from 'react-router-dom';

import MovieSearchBar from './MovieSearchBar.js';
import MovieSearchResults from './MovieSearchResults';

const MovieSearch = () => {

    const [error, setError] = useState(null);
    // Fetch query parameters
    const { search } = useLocation(); // Base code to fetch location of object containing query parameter in form of a string
    const queryParams = new URLSearchParams(search); // Built in JS constructor to work with query paremters

    const query = queryParams.get('query'); // Retrieves value of query parameter called '?query' and so on if more than 1
    const [searchResults, setSearchResults] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await movieSearchTMDB(query);
                setSearchResults(data.search_results || []);

            } catch (error) {
                setError(error);
            }
        };

        fetchData();
        }, [query]);
      
    // RENDER COMPONENT
    return (
        <React.Fragment>
            <NavBar />
            <Container >
            <MovieSearchBar/>
            {error ? (
                <h1 className='error'><u>Error:</u> {error.message}</h1>
                ) : (
                searchResults && (
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
                ))}
            </Container>
        </React.Fragment>
  );
};
export default MovieSearch;