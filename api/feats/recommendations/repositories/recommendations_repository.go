package repositories

import (
	"app/feats/recommendations/http"
	"app/models"
	"app/pkg/db"
	"app/pkg/logger"
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type RecommendationRepository struct {
	db                        *db.DB
	logger                    *logger.Logger
	recommendationsHttpClient *http.RecommendationsHttpClient
}

func NewRecommendationRepository(db *db.DB, logger *logger.Logger, recommendationsHttpClient *http.RecommendationsHttpClient) *RecommendationRepository {
	return &RecommendationRepository{db: db, logger: logger, recommendationsHttpClient: recommendationsHttpClient}
}

func (r *RecommendationRepository) GetRecommendations(ratings *[]models.AnonymousMovieRating) (*[]models.MoviePreview, error) {
	similarUsers, err := r.recommendationsHttpClient.Recommend(ratings)
	if err != nil {
		return nil, err
	}

	// Extract movie ids
	movieIds := make([]int, len(*ratings))
	for i := range *ratings {
		movieIds[i] = (*ratings)[i].MovieId
	}

	var movies []models.MoviePreview
	err = pgxscan.Select(
		context.Background(),
		r.db,
		&movies,
		getRecommendationsSQL,
		pgx.NamedArgs{
			"userIds":   similarUsers.SimilarUsers,
			"distances": similarUsers.Distances,
			"movieIds":  movieIds,
		},
	)

	if err != nil {
		r.logger.Error(fmt.Sprintf("Failed to select movies recommendations:\n%s", err.Error()))
		return nil, err
	}

	return &movies, err
}

var getRecommendationsSQL = `
WITH similar_users AS (
    SELECT unnest(@userIds::int[]) AS user_id, unnest(@distances::float8[]) AS distance
),
     weighted_ratings AS (
         SELECT
             ur.movie_id,
             SUM(ur.rating * (1 - su.distance)) AS weighted_sum,
             SUM(1 - su.distance) AS total_distance
         FROM ratings ur
            JOIN similar_users su ON ur.user_id = su.user_id
         GROUP BY
             ur.movie_id
     ),
     predicted_ratings AS (
         SELECT
             m.id,
             m.title,
			 m.tagline,
             m.vote_average,
             wr.weighted_sum / wr.total_distance AS predicted_rating
         FROM
             movies m
                 JOIN weighted_ratings wr ON m.id = wr.movie_id
     )
SELECT
    id,
    title,
    tagline,
	vote_average
FROM
    predicted_ratings
WHERE id NOT IN(SELECT unnest(@movieIds::int[]))
ORDER BY
    predicted_rating DESC
LIMIT 50;
`
