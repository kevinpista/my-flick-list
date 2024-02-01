import React, { useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/MovieSearchResults.css';
import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../utils/formatUtils';
import no_image_placeholder from '../static/no_image_placeholder.jpg';


// Render the movie results. Function by the MovieSearch component
const MovieSearchResults = ({ id, title, releaseDate, description, posterURL }) => {
    // posterURL will either be TMDB link or null
    const finalPosterURL = posterURL ? posterURL : no_image_placeholder;
    const formatedReleaseDate = formatReleaseDate(releaseDate)

    return (
      <div id={`card_movie_${id}`} className="card tight">
        <div className="poster-wrapper">
          <a
            data-id={id}
            className="result"
            href={`/movie/${id}?language=en-US`}
          >
              <img
                className="poster"
                src={finalPosterURL}
                alt={title}
              />
          </a>
        </div>

        <div className="details">
          <a
            data-id={id}
            className="result"
            href={`/movie/${id}?language=en-US`}
          >
            <h2 className="title">
              {title}
            </h2>
          </a>

          <span className="release_date">
            {formatedReleaseDate}
          </span>

          <p className="overview">
            {description}
          </p>

        </div>
    </div>
  );
};
export default MovieSearchResults;