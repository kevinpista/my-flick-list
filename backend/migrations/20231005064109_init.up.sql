-- Add up migration script here
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE, 
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE jwt_token (
  id SERIAL PRIMARY KEY,
  users_id UUID REFERENCES users(id),
  jwt_token_key TEXT,
  issued_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  expired_at TIMESTAMP  
);

CREATE TABLE movie (
  id INT PRIMARY KEY UNIQUE,
  original_title VARCHAR(255) NOT NULL,
  overview TEXT,
  tagline VARCHAR(255), 
  release_date DATE, 
  poster_path VARCHAR(255),
  backdrop_path VARCHAR(255),
  runtime INT, 
  adult BOOLEAN,
  budget NUMERIC(12,0),
  revenue NUMERIC(12,0),
  rating NUMERIC(3,2) CHECK (rating >= 0 AND rating <= 10),
  votes INT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE watchlist (
  id SERIAL PRIMARY KEY,
  users_id UUID REFERENCES users(id), 
  name VARCHAR(60) NOT NULL,
  description TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE watchlist_item (
  id SERIAL PRIMARY KEY,
  watchlist_id INT REFERENCES watchlist(id),
  movie_id INT REFERENCES movie(id), 
  checkmarked BOOLEAN DEFAULT false,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  UNIQUE (watchlist_id, movie_id)
);

CREATE TABLE watchlist_item_note (
  watchlist_item_id INT REFERENCES watchlist_item(id), 
  item_notes TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE genre (
  movie_id INT REFERENCES movie(id),
  genre_id INT NOT NULL,
  genre VARCHAR(255) NOT NULL
);