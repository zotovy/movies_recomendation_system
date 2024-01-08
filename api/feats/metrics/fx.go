package metrics

import (
	"app/feats/metrics/endpoints"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Context = fx.Options(
	fx.Provide(endpoints.GetMetricController),
	fx.Provide(GetMetricRoutes),
	fx.WithLogger(func(logger *zap.Logger) *zap.Logger {
		return logger.Named("metrics")
	}),
)
