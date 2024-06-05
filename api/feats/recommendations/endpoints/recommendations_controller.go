package endpoints

import (
	"app/feats/recommendations/repositories"
	"app/pkg/db"
	"app/pkg/logger"
)

type Controller struct {
	db         *db.DB
	logger     *logger.Logger
	repository *repositories.RecommendationRepository
}

func NewController(db *db.DB, logger *logger.Logger, repository *repositories.RecommendationRepository) *Controller {
	return &Controller{db: db, logger: logger, repository: repository}
}
