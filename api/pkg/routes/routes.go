package routes

import (
	"app/feats/docs"
	"app/feats/metrics"
	"app/feats/search"
)

type Route interface {
	Setup()
}

type Routes []Route

func GetRoutes(metricRoutes metrics.MetricRoutes, docsRoutes docs.Routes, searchRoutes search.Routes) Routes {
	return Routes{
		metricRoutes,
		docsRoutes,
		searchRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
