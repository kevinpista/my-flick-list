import React, { useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/MovieSearchResults.css';
import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../utils/formatUtils';


// Render the movie results. Function by the MovieSearch component
const MovieSearchResults = ({ id, title, releaseDate, description, posterURL }) => {
    
    const formatedReleaseDate = formatReleaseDate(releaseDate)

    return (
    <div id={`card_movie_${id}`} className="card tight">
      <div className="wrapper">
        <div className="image">
          <div className="poster">
            <a
              data-id={id}
              className="result"
              href={`/movie/${id}?language=en-US`}
            >
              <img
                className="poster"
                src={posterURL}
                srcSet={`${posterURL} 1x, ${posterURL} 2x`}
                alt={title}
              />
            </a>
          </div>
        </div>

        <div className="details">
          <div className="wrapper">
            <div className="title">
                <a
                  data-id={id}
                  className="result"
                  href={`/movie/${id}?language=en-US`}
                >
                  <h2>{title}</h2>
                </a>

              <span className="release_date">{formatedReleaseDate}</span>
            </div>
          </div>

          <div className="overview">
            <p>{description}</p>
          </div>
        </div>
      </div>
    </div>
  );
};
export default MovieSearchResults;