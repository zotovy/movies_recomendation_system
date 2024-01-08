package metrics

import (
	"app/feats/metrics/endpoints"
	"app/pkg/router"
	"go.uber.org/zap"
)

type MetricRoutes struct {
	logger           *zap.Logger
	router           *router.Router
	metricController endpoints.MetricController
}

func GetMetricRoutes(
	logger *zap.Logger,
	router *router.Router,
	controller endpoints.MetricController,
) MetricRoutes {
	return MetricRoutes{
		logger:           logger,
		router:           router,
		metricController: controller,
	}
}

func (routes MetricRoutes) Setup() {
	routes.logger.Info("Setup metric routes")
	routes.router.GET("/api/v1/metrics/ping", routes.metricController.Ping)
}
