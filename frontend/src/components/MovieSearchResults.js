import React, { useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/MovieSearchResults.css';
import { formatReleaseDate, formatRuntime, formatVoteCount, formatFinancialData } from '../utils/formatUtils';


// Render the movie results. Function by the MovieSearch component
const MovieSearchResults = ({ id, title, releaseDate, description, posterURL }) => {

  return (
    <div id={`card_movie_${id}`} className="card tight">
      <div className="wrapper">
        <div className="image">
          <div className="poster">
            <a
              data-id={id}
              data-media-type="movie"
              data-media-adult="false"
              className="result"
              href={`/movie/${id}?language=en-US`}
            >
              <img
                loading="lazy"
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
              <div>
                <a
                  data-id={id}
                  data-media-type="movie"
                  data-media-adult="false"
                  className="result"
                  href={`/movie/${id}?language=en-US`}
                >
                  <h2>{title}</h2>
                </a>
              </div>

              <span className="release_date">{releaseDate}</span>
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