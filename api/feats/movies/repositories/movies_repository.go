package repositories

import (
	"app/models"
	"app/pkg/db"
	"app/pkg/logger"
	"app/pkg/redis"
	"context"
	"encoding/json"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"time"
)

type MoviesRepository struct {
	logger *logger.Logger
	db     *db.DB
	redis  *redis.Client
}

func NewMoviesRepository(logger *logger.Logger, db *db.DB, redis *redis.Client) *MoviesRepository {
	return &MoviesRepository{logger: logger, db: db, redis: redis}
}

func (r *MoviesRepository) GetById(id int) (*models.Movie, error) {
	cached, err := r.redis.Get(context.Background(), fmt.Sprintf("movie-%d", id)).Bytes()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == nil {
		r.logger.Info("Found cached movie")
		var cachedMovie *models.Movie
		err := json.Unmarshal(cached, &cachedMovie)
		if err == nil {
			return cachedMovie, nil
		}
	}

	// Get movie
	movie, err := r.getMovieById(id)
	if err != nil {
		return nil, err
	}

	if movie == nil {
		return nil, err
	}

	// Genres
	genres, err := r.getMovieGenres(id)
	if err != nil {
		return nil, err
	} else {
		movie.Genres = genres
	}

	// Keywords
	keywords, err := r.getMovieKeywords(id)
	if err != nil {
		return nil, err
	} else {
		movie.Keywords = keywords
	}

	// Companies
	companies, err := r.getMovieCompanies(id)
	if err != nil {
		return nil, err
	} else {
		movie.ProductionCompanies = companies
	}

	// Countries
	countries, err := r.getMovieCountries(id)
	if err != nil {
		return nil, err
	} else {
		movie.Countries = countries
	}

	// Languages
	languages, err := r.getMovieLanguages(id)
	if err != nil {
		return nil, err
	} else {
		movie.Languages = languages
	}

	// Cast
	cast, err := r.getMovieCast(id)
	if err != nil {
		return nil, err
	} else {
		movie.Cast = cast
	}

	// Crew
	crew, err := r.getMovieCrew(id)
	if err != nil {
		return nil, err
	} else {
		movie.Crew = crew
	}

	// Save to cache
	jsonMovie, err := json.Marshal(movie)
	if err != nil {
		return nil, err
	}
	r.redis.Set(
		context.Background(),
		fmt.Sprintf("movie-%d", id),
		jsonMovie, 60*time.Minute,
	)

	return movie, err
}

func (r *MoviesRepository) getMovieById(id int) (*models.Movie, error) {
	var movie models.Movie

	err := pgxscan.Get(context.Background(), r.db, &movie, selectMovieByIdSQL, id)

	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		r.logger.Error(err.Error())
		return nil, err
	}

	return &movie, nil
}

func (r *MoviesRepository) getMovieGenres(id int) (*[]models.Genre, error) {
	var genres []models.Genre

	err := pgxscan.Select(context.Background(), r.db, &genres, selectMoviesGenresSQL, id)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &genres, nil
}

func (r *MoviesRepository) getMovieKeywords(id int) (*[]models.Keyword, error) {
	var keywords []models.Keyword

	err := pgxscan.Select(context.Background(), r.db, &keywords, selectMovieKeywordsSQL, id)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &keywords, nil
}

func (r *MoviesRepository) getMovieCompanies(id int) (*[]models.Company, error) {
	var companies []models.Company

	err := pgxscan.Select(context.Background(), r.db, &companies, selectMovieCompaniesSQL, id)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &companies, nil
}

func (r *MoviesRepository) getMovieCountries(id int) (*[]models.Country, error) {
	var countries []models.Country

	err := pgxscan.Select(context.Background(), r.db, &countries, selectMovieCountriesSQL, id)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &countries, nil
}

func (r *MoviesRepository) getMovieLanguages(id int) (*[]models.Language, error) {
	var languages []models.Language

	err := pgxscan.Select(context.Background(), r.db, &languages, selectMovieLanguagesSQL, id)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &languages, nil
}

func (r *MoviesRepository) getMovieCast(id int) (*[]models.Cast, error) {
	var casts []models.Cast

	err := pgxscan.Select(context.Background(), r.db, &casts, selectMovieCastSQL, id)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &casts, nil
}

func (r *MoviesRepository) getMovieCrew(id int) (*[]models.Crew, error) {
	var crews []models.Crew

	err := pgxscan.Select(context.Background(), r.db, &crews, selectMovieCrewSQL, id)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &crews, nil
}

var selectMovieByIdSQL = `
SELECT id,
       budget,
       homepage,
       original_language,
       original_title,
       overview,
       popularity,
       release_date,
       revenue,
       runtime,
       status,
       tagline,
       title,
       vote_average,
       vote_count
FROM movies
WHERE id = $1
`

var selectMoviesGenresSQL = `
SELECT g.*
FROM genres g
         JOIN movies_genres mg ON g.id = mg.genre_id
WHERE mg.movie_id = $1;
`

var selectMovieKeywordsSQL = `
SELECT k.*
FROM keywords k
         JOIN movies_keywords mk on k.id = mk.keywords_id
WHERE mk.movie_id = $1;
`

var selectMovieCompaniesSQL = `
SELECT pc.*
FROM production_companies pc
         JOIN movies_production_companies mpc on pc.id = mpc.production_company_id
WHERE mpc.movie_id = $1;
`

var selectMovieCountriesSQL = `
SELECT c.*
FROM countries c
         JOIN movies_production_countries mpc on c.code = mpc.country_code
WHERE mpc.movie_id = $1;
`

var selectMovieLanguagesSQL = `
SELECT l.*
FROM languages l
         JOIN movies_spoken_languages ml on l.code = ml.language_code
WHERE ml.movie_id = $1;
`

var selectMovieCastSQL = `
SELECT c.*
FROM "cast" c
         JOIN movies_cast mc
              ON c.id = mc.cast_id
WHERE mc.movie_id = $1
ORDER BY c."order"
LIMIT 10;
`

var selectMovieCrewSQL = `
SELECT c.*
FROM crew c
         JOIN movies_crew mc on c.id = mc.crew_id
WHERE mc.movie_id = $1
LIMIT 10;
`
