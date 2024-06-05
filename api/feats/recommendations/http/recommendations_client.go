package http

import (
	"app/models"
	"app/pkg/config"
	"app/pkg/logger"
	"bytes"
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

func (c *RecommendationsHttpClient) Recommend(ratings *[]models.AnonymousMovieRating) (*models.SimilarUsers, error) {
	// Marshal ratings
	jsonBody, _ := json.Marshal(ratings)

	// Build url
	url := fmt.Sprintf("%s/recommend", c.config.Services.RecommendationsAPI.BaseUrl)

	// Do request
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))

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
	var similarUsers *models.SimilarUsers
	err = json.Unmarshal(body, &similarUsers)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to invoke RecommendationsHttpClient.Recommend:\n%s", err.Error()))
		return nil, err
	}

	return similarUsers, nil
}
