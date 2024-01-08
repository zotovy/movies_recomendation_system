package router

import (
	"app/pkg/config"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type Router = gin.Engine

func GetRouter(config *config.Config, logger *zap.Logger) *Router {
	router := gin.New()

	if config.Environment == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger, true))

	router.Use(gin.Recovery())

	err := router.SetTrustedProxies(nil)
	if err != nil {
		logger.Fatal("Failed to set trusted proxies", zap.Error(err))
	}

	return router
}
