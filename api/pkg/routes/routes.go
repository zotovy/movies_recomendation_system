package routes

import (
	"app/feats/docs"
	"app/feats/metrics"
)

type Route interface {
	Setup()
}

type Routes []Route

func GetRoutes(metricRoutes metrics.MetricRoutes, docsRoutes docs.Routes) Routes {
	return Routes{
		metricRoutes,
		docsRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
