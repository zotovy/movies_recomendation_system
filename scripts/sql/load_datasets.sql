COPY movies (id, budget, homepage, original_language, original_title, overview, popularity, release_date, revenue,
             runtime, status, tagline, title, vote_average,
             vote_count) FROM '/var/datasets/movies.csv' DELIMITER ',' CSV HEADER FORCE NOT NULL overview;