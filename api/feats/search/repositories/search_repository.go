package repositories

import (
	"app/feats/search/models"
	"app/pkg/db"
	"app/pkg/logger"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
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
	err := pgxscan.Select(ctx, r.db, &movies, "SELECT id, title, overview FROM movies WHERE document @@ websearch_to_tsquery($1, '{0}:*') ORDER BY ts_rank(document, websearch_to_tsquery($1, '{0}:*')) desc LIMIT 15", query)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return movies, nil
}

func (r *Repository) AutoCompleteMovies(query string) ([]*models.MoviePreview, error) {
	ctx := context.Background()

	movies := make([]*models.MoviePreview, 0)
	err := pgxscan.Select(ctx, r.db, &movies, "SELECT id, title, overview FROM movies WHERE title ILIKE ('%' || $1 || '%') ORDER BY popularity desc LIMIT 5", query)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return movies, nil
}
