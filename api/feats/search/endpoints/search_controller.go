package endpoints

import (
	"app/feats/search/repositories"
	"app/pkg/db"
	"app/pkg/logger"
)

type SearchController struct {
	logger     *logger.Logger
	db         *db.DB
	repository *repositories.Repository
}

func NewSearchController(logger *logger.Logger, db *db.DB, repository *repositories.Repository) *SearchController {
	return &SearchController{logger: logger, db: db, repository: repository}
}
