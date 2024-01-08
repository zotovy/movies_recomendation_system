package app

import (
	"app/feats/docs"
	"app/feats/metrics"
	"app/pkg"
	"app/pkg/config"
	"app/pkg/logger"
	"app/pkg/router"
	"app/pkg/routes"
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(
		pkg.Module,

		// Utils
		fx.Provide(
			routes.GetRoutes,
			logger.GetLogger,
		),

		// Feats
		metrics.Context,
		docs.Context,

		// Lifecycle
		fx.Invoke(registerHooks),

		// Logger
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)
}

func registerHooks(
	lifecycle fx.Lifecycle,
	router *router.Router,
	config *config.Config,
	logger *zap.Logger,
	routes routes.Routes,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				// Log the start of the application with the configured host and port
				logger.Info("Starting application...")

				// ======== SET UP COMPONENTS ========
				// Perform any necessary setup or initialization tasks for the middlewares

				// Perform any necessary setup or initialization tasks for the routes
				routes.Setup()

				port := fmt.Sprintf(":%d", config.Api.Port)
				logger.Debug(port)

				// Start the router by running it in a separate goroutine
				go func() {
					err := router.Run(port)
					if err != nil {
						logger.Fatal("Failed to start app", zap.Error(err))
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				// Log the stop of the application and any associated error
				logger.Fatal(
					fmt.Sprintf(
						"Stopping application. Error: %s", ctx.Err(),
					),
				)

				return nil
			},
		},
	)
}
