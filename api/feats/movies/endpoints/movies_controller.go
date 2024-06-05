package endpoints

import (
	"app/feats/movies/repositories"
	"app/pkg/logger"
)

type MoviesController struct {
	logger     *logger.Logger
	repository *repositories.MoviesRepository
}

func NewMoviesController(logger *logger.Logger, repository *repositories.MoviesRepository) *MoviesController {
	return &MoviesController{logger: logger, repository: repository}
}
