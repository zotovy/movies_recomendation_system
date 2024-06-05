package movies

import (
	"app/feats/movies/endpoints"
	"app/feats/movies/repositories"
	"app/pkg/logger"
	"go.uber.org/fx"
)

var Context = fx.Options(
	fx.Provide(repositories.NewMoviesRepository),
	fx.Provide(endpoints.NewMoviesController),
	fx.Provide(NewRoutes),
	fx.WithLogger(func(logger *logger.Logger) *logger.Logger {
		return logger.Named("movies")
	}),
)
