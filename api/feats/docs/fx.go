package docs

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Context = fx.Options(
	fx.Provide(GetDocsRoutes),
	fx.WithLogger(func(logger *zap.Logger) *zap.Logger {
		return logger.Named("docs")
	}),
)
