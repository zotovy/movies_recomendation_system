package repositories

import (
	"app/models"
	"app/pkg/db"
	"app/pkg/logger"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"strings"
)

type Repository struct {
	db     *db.DB
	logger *logger.Logger
}

func NewRepository(db *db.DB, logger *logger.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (r *Repository) SearchMovies(query string) ([]*models.MoviePreview, error) {
	ctx := context.Background()

	movies := make([]*models.MoviePreview, 0)
	err := pgxscan.Select(ctx, r.db, &movies, "SELECT id, title, overview, vote_average FROM movies WHERE document @@ websearch_to_tsquery($1, '{0}:*') ORDER BY ts_rank(document, websearch_to_tsquery($1, '{0}:*')) desc LIMIT 15", query)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return movies, nil
}

func (r *Repository) AutoCompleteMovies(query string) ([]*models.MoviePreview, error) {
	ctx := context.Background()

	// tsQuery building
	queryWords := strings.Split(query, " ")
	tsQuery := strings.Join(queryWords, " & ") + ":*"

	movies := make([]*models.MoviePreview, 0)
	err := pgxscan.Select(
		ctx,
		r.db,
		&movies,
		autocompleteSQL,
		pgx.NamedArgs{
			"tsQuery": tsQuery,
			"query":   query,
		},
	)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return movies, nil
}

var autocompleteSQL = `
SELECT id, title, tagline, vote_average
FROM movies
WHERE document @@ to_tsquery(@tsQuery)
   OR similarity(lower(title), @query) > 0.1
ORDER BY ts_rank(document, to_tsquery(@tsQuery)) +
         similarity(lower(title), @query) +
         (movies.popularity / 900 / 2) desc
LIMIT 5;
`
