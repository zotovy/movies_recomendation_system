package http

import (
	"app/models"
	"app/pkg/config"
	"app/pkg/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RecommendationsHttpClient struct {
	logger *logger.Logger
	config *config.Config
}

func NewRecommendationsHttpClient(logger *logger.Logger, config *config.Config) *RecommendationsHttpClient {
	return &RecommendationsHttpClient{logger: logger, config: config}
}

func (c *RecommendationsHttpClient) Recommend(userId int, amount int) (*[]models.AnonymousMovieRating, error) {
	// Build url
	url := fmt.Sprintf("%s/recommend?user_id=%d&amount=%d", c.config.Services.RecommendationsAPI.BaseUrl, userId, amount)

	// Do request
	response, err := http.Get(url)

	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to invoke RecommendationsHttpClient.Recommend:\n%s", err.Error()))
		return nil, err
	}
	defer response.Body.Close()

	// Reading response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to invoke RecommendationsHttpClient.Recommend:\n%s", err.Error()))
		return nil, err
	}

	// Unmarshal response body
	var movieRatings *[]models.AnonymousMovieRating
	err = json.Unmarshal(body, &movieRatings)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to invoke RecommendationsHttpClient.Recommend:\n%s", err.Error()))
		return nil, err
	}

	return movieRatings, nil
}
