package routes

import (
	"app/feats/docs"
	"app/feats/metrics"
	"app/feats/movies"
	"app/feats/recommendations"
	"app/feats/search"
)

type Route interface {
	Setup()
}

type Routes []Route

func GetRoutes(
	metricRoutes metrics.MetricRoutes,
	docsRoutes docs.Routes,
	searchRoutes search.Routes,
	moviesRoutes *movies.Routes,
	recommendationsRoutes *recommendations.Routes,
) Routes {
	return Routes{
		metricRoutes,
		docsRoutes,
		searchRoutes,
		moviesRoutes,
		recommendationsRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
