package search

import (
	"app/feats/search/endpoints"
	"app/feats/search/repositories"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Context = fx.Options(
	fx.Provide(endpoints.NewSearchController),
	fx.Provide(NewSearchRoutes),
	fx.Provide(repositories.NewRepository),
	fx.WithLogger(func(logger *zap.Logger) *zap.Logger {
		return logger.Named("search")
	}),
)
