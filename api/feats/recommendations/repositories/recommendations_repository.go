package repositories

import (
	"app/feats/recommendations/http"
	"app/models"
	"app/pkg/db"
	"app/pkg/logger"
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type RecommendationRepository struct {
	db                        *db.DB
	logger                    *logger.Logger
	recommendationsHttpClient *http.RecommendationsHttpClient
}

func NewRecommendationRepository(db *db.DB, logger *logger.Logger, recommendationsHttpClient *http.RecommendationsHttpClient) *RecommendationRepository {
	return &RecommendationRepository{db: db, logger: logger, recommendationsHttpClient: recommendationsHttpClient}
}

func (r *RecommendationRepository) GetRecommendations(userId int, amount int) (*[]models.MoviePreviewWithRating, error) {
	ratings, err := r.recommendationsHttpClient.Recommend(userId, amount)
	if err != nil {
		return nil, err
	}

	var values []interface{}
	var placeholders string
	for i, rating := range *ratings {
		if i > 0 {
			placeholders += ", "
		}
		placeholders += fmt.Sprintf("($%d::int, $%d::float)", 2*i+1, 2*i+2)
		values = append(values, rating.MovieId, rating.Rating)
	}

	var movies []models.MoviePreviewWithRating
	err = pgxscan.Select(
		context.Background(),
		r.db,
		&movies,
		fmt.Sprintf(getRecommendationsSQL, placeholders),
		values...,
	)

	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to select movies recommendations:\n%s", err.Error()))
		return nil, err
	}

	return &movies, err
}

var getRecommendationsSQL = `
WITH ratings AS (
SELECT * FROM (VALUES %s) AS temp (movie_id, rating))
SELECT m.id, m.tagline, m.vote_average, m.title, r.rating
FROM movies m
JOIN ratings r ON m.id = r.movie_id;
`
