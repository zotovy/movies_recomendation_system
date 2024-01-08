package docs

import (
	docs "app/docs"
	"app/pkg/config"
	"app/pkg/logger"
	"app/pkg/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type Routes struct {
	logger *zap.Logger
	router *router.Router
	config *config.Config
}

func GetDocsRoutes(logger *logger.Logger, router *router.Router, config *config.Config) Routes {
	return Routes{
		logger,
		router,
		config,
	}
}

func (routes Routes) Setup() {
	routes.logger.Info("Setup docs router")
	docs.SwaggerInfo.BasePath = "/api/v1"
	//url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", routes.config.Api.Host))
	routes.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
