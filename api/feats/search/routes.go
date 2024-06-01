package search

import (
	"app/feats/search/endpoints"
	"app/pkg/logger"
	"app/pkg/router"
)

type Routes struct {
	logger     *logger.Logger
	router     *router.Router
	controller *endpoints.SearchController
}

func NewSearchRoutes(logger *logger.Logger, router *router.Router, controller *endpoints.SearchController) Routes {
	return Routes{logger: logger, router: router, controller: controller}
}

func (routes Routes) Setup() {
	routes.logger.Info("Setup search routes")
	routes.router.GET("/api/v1/search/movies", routes.controller.SearchMovies)
	routes.router.GET("/api/v1/search/movies/autocomplete", routes.controller.AutoCompleteMovies)
}
