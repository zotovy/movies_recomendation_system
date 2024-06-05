package recommendations

import (
	"app/feats/recommendations/endpoints"
	"app/pkg/logger"
	"app/pkg/router"
)

type Routes struct {
	logger     *logger.Logger
	router     *router.Router
	controller *endpoints.Controller
}

func NewRoutes(logger *logger.Logger, router *router.Router, controller *endpoints.Controller) *Routes {
	return &Routes{logger: logger, router: router, controller: controller}
}

func (routes *Routes) Setup() {
	routes.router.POST("/api/v1/recommendations/movies", routes.controller.GetMoviesRecommendations)
}
