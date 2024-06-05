package movies

import (
	"app/feats/movies/endpoints"
	"app/pkg/logger"
	"app/pkg/router"
)

type Routes struct {
	logger     *logger.Logger
	router     *router.Router
	controller *endpoints.MoviesController
}

func NewRoutes(logger *logger.Logger, router *router.Router, controller *endpoints.MoviesController) *Routes {
	return &Routes{logger: logger, router: router, controller: controller}
}

func (routes Routes) Setup() {
	routes.router.GET("/api/v1/movies/:id", routes.controller.GetMovieById)
}
