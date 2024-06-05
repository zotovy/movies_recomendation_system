CREATE TABLE IF NOT EXISTS movies
(
    id                int PRIMARY KEY,
    budget            bigserial    NOT NULL,
    homepage          varchar(255),
    original_language varchar(255) NOT NULL,
    original_title    text         NOT NULL,
    overview          text         NOT NULL DEFAULT '',
    popularity        float        NOT NULL,
    release_date      date         NOT NULL,
    revenue           bigserial    NOT NULL,
    runtime           float        NOT NULL,
    status            varchar(255) NOT NULL,
    tagline           text,
    title             text         NOT NULL,
    vote_average      float        NOT NULL,
    vote_count        int          NOT NULL
);

CREATE TABLE IF NOT EXISTS genres
(
    id   INT PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS keywords
(
    id   INT PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS production_companies
(
    id   INT PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS countries
(
    code varchar(2) PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS languages
(
    code varchar(2) PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "cast"
(
    id        INT PRIMARY KEY,
    character text         NOT NULL,
    name      varchar(255) NOT NULL,
    gender    INT          NOT NULL,
    "order"   INT          NOT NULL
);

CREATE TABLE IF NOT EXISTS crew
(
    id         INT PRIMARY KEY,
    name       varchar(255) NOT NULL,
    job        varchar(255) NOT NULL,
    gender     INT          NOT NULL,
    department varchar(255) NOT NULL
);

CREATE TABLE movies_genres
(
    movie_id INTEGER REFERENCES movies (id),
    genre_id INTEGER REFERENCES genres (id),
    CONSTRAINT movies_genres_pk PRIMARY KEY (movie_id, genre_id)
);

CREATE TABLE movies_keywords
(
    movie_id    INTEGER REFERENCES movies (id),
    keywords_id INTEGER REFERENCES keywords (id),
    CONSTRAINT movies_keywords_pk PRIMARY KEY (movie_id, keywords_id)
);

CREATE TABLE movies_production_companies
(
    movie_id              INTEGER REFERENCES movies (id),
    production_company_id INTEGER REFERENCES production_companies (id),
    CONSTRAINT movies_production_companies_pk PRIMARY KEY (movie_id, production_company_id)
);

CREATE TABLE movies_production_countries
(
    movie_id     INTEGER REFERENCES movies (id),
    country_code varchar(2) REFERENCES countries (code),
    CONSTRAINT movies_production_countries_pk PRIMARY KEY (movie_id, country_code)
);

CREATE TABLE movies_spoken_languages
(
    movie_id      INTEGER REFERENCES movies (id),
    language_code varchar(2) REFERENCES languages (code),
    CONSTRAINT movies_spoken_languages_pk PRIMARY KEY (movie_id, language_code)
);

CREATE TABLE movies_cast
(
    movie_id INTEGER REFERENCES movies (id),
    cast_id  INTEGER REFERENCES "cast" (id),
    CONSTRAINT movies_cast_pk PRIMARY KEY (movie_id, cast_id)
);

CREATE TABLE movies_crew
(
    movie_id INTEGER REFERENCES movies (id),
    crew_id  INTEGER REFERENCES crew (id),
    CONSTRAINT movies_crew_pk PRIMARY KEY (movie_id, crew_id)
);


CREATE TABLE ratings
(
    movie_id INTEGER REFERENCES movies(id),
    user_id INTEGER,
    rating FLOAT
)