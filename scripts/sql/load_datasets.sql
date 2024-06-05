COPY movies (id, budget, homepage, original_language, original_title, overview, popularity, release_date, revenue,
             runtime, status, tagline, title, vote_average,
             vote_count) FROM '/var/datasets/movies.csv' DELIMITER ',' CSV HEADER FORCE NOT NULL overview;

COPY genres (id, name) FROM '/var/datasets/genres.csv' DELIMITER ',' CSV HEADER;

COPY keywords (id, name) FROM '/var/datasets/keywords.csv' DELIMITER ',' CSV HEADER;

COPY production_companies (id, name) FROM '/var/datasets/production_companies.csv' DELIMITER ',' CSV HEADER;

COPY countries (code, name) FROM '/var/datasets/countries.csv' DELIMITER ',' CSV HEADER;

COPY languages (code, name) FROM '/var/datasets/languages.csv' DELIMITER ',' CSV HEADER;

COPY "cast" (id, character, name, gender, "order") FROM '/var/datasets/cast.csv' DELIMITER ',' CSV HEADER;

COPY crew (id, name, job, gender, department) FROM '/var/datasets/crew.csv' DELIMITER ',' CSV HEADER;

COPY movies_genres (movie_id, genre_id)  FROM '/var/datasets/movies_genres.csv' DELIMITER ',' CSV HEADER;

COPY movies_keywords (movie_id, keywords_id)  FROM '/var/datasets/movies_keywords.csv' DELIMITER ',' CSV HEADER;

COPY movies_production_companies (movie_id, production_company_id) FROM '/var/datasets/movies_production_companies.csv' DELIMITER ',' CSV HEADER;

COPY movies_production_countries (movie_id, country_code) FROM '/var/datasets/movies_production_countries.csv' DELIMITER ',' CSV HEADER;

COPY movies_spoken_languages (movie_id, language_code) FROM '/var/datasets/movies_spoken_languages.csv' DELIMITER ',' CSV HEADER;

COPY movies_cast (movie_id, cast_id) FROM '/var/datasets/movies_cast.csv' DELIMITER ',' CSV HEADER;

COPY movies_crew (movie_id, crew_id) FROM '/var/datasets/movies_crew.csv' DELIMITER ',' CSV HEADER;

COPY ratings (user_id, movie_id, rating) FROM '/var/datasets/ratings.csv' DELIMITER ',' CSV HEADER;