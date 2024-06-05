package recommendations

import (
	"app/feats/recommendations/endpoints"
	"app/feats/recommendations/http"
	"app/feats/recommendations/repositories"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Context = fx.Options(
	fx.Provide(endpoints.NewController),
	fx.Provide(NewRoutes),
	fx.Provide(repositories.NewRecommendationRepository),
	fx.Provide(http.NewRecommendationsHttpClient),
	fx.WithLogger(func(logger *zap.Logger) *zap.Logger {
		return logger.Named("recommendations")
	}),
)
