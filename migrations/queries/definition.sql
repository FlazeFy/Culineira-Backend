-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE Countries (
    id VARCHAR PRIMARY KEY,
    country_code VARCHAR NOT NULL,
	country_name VARCHAR NOT NULL,
	country_continent VARCHAR,

	is_active BOOLEAN NOT NULL,  
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP 
);

CREATE TABLE Users (
    id VARCHAR PRIMARY KEY,
    country_code VARCHAR NOT NULL,
	slug_name VARCHAR NOT NULL,
	username VARCHAR NOT NULL,
	first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    job VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,

	is_active BOOLEAN NOT NULL,  
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP, 
	updated_by VARCHAR
);

CREATE TABLE Recipes (
    id VARCHAR PRIMARY KEY,
	country_code VARCHAR NOT NULL,
	recipe_name VARCHAR NOT NULL,
	recipe_desc VARCHAR,
	recipe_type VARCHAR NOT NULL,
	recipe_time_spend INT NOT NULL,
	recipe_cal INT,
	recipe_Level VARCHAR NOT NULL,
	recipe_image_url VARCHAR,
	recipe_video_url VARCHAR, 

	is_private BOOLEAN NOT NULL,  
	created_at TIMESTAMP NOT NULL,
	created_by VARCHAR NOT NULL,
	updated_at TIMESTAMP, 
	updated_by VARCHAR, 
	deleted_at TIMESTAMP, 
	deleted_by VARCHAR
);

CREATE TABLE Steps (
    id VARCHAR PRIMARY KEY,
    recipes_id VARCHAR NOT NULL,
	steps_body VARCHAR NOT NULL,
	steps_image_url VARCHAR,
    steps_video_url VARCHAR,

	is_optional BOOLEAN NOT NULL,  
	created_at TIMESTAMP NOT NULL,
	created_by VARCHAR NOT NULL,
	updated_at TIMESTAMP, 
	updated_by VARCHAR
);

CREATE TABLE Ingredients (
    id VARCHAR PRIMARY KEY,
    recipes_id VARCHAR NOT NULL,
	ingredient_name VARCHAR NOT NULL,
	ingredient_volume VARCHAR,

	is_optional BOOLEAN NOT NULL,  
	created_at TIMESTAMP NOT NULL,
	created_by VARCHAR NOT NULL,
	updated_at TIMESTAMP, 
	updated_by VARCHAR
);

CREATE TABLE Likes (
    id VARCHAR PRIMARY KEY,
    recipes_id VARCHAR NOT NULL,

	created_at TIMESTAMP NOT NULL,
	created_by VARCHAR NOT NULL
);

CREATE TABLE Comments (
    id VARCHAR PRIMARY KEY,
	recipes_id VARCHAR NOT NULL,
	comment_body VARCHAR NOT NULL,
	comment_file_url VARCHAR,

	created_at TIMESTAMP NOT NULL,
	created_by VARCHAR NOT NULL,
	deleted_at TIMESTAMP, 
	deleted_by VARCHAR
);

-- +migrate StatementEnd