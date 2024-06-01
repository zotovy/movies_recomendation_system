package pkg

import (
	"app/pkg/config"
	"app/pkg/db"
	"app/pkg/router"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(
		router.GetRouter,
		config.NewConfig,
		db.NewDB,
	),
	fx.WithLogger(func(logger *zap.Logger) *zap.Logger {
		return logger.Named("pkg")
	}),
)
