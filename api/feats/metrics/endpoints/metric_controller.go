package endpoints

import (
	"app/pkg/db"
	"go.uber.org/zap"
)

type MetricController struct {
	logger *zap.Logger
	db     *db.DB
}

func GetMetricController(logger *zap.Logger, db *db.DB) MetricController {
	return MetricController{logger: logger, db: db}
}
