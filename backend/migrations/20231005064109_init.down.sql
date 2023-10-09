-- Add down migration script here
-- Add down migration script here
drop schema public cascade;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS jwt_tokens;
DROP TABLE IF EXISTS watchlists;
DROP TABLE IF EXISTS watchlists_items;
DROP TABLE IF EXISTS watchlists_items_notes;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS genre;
