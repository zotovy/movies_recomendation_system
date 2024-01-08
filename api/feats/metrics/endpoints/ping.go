package endpoints

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// Ping godoc
// @Summary ping
// @Description do ping
// @Tags metrics
// @Accept json
// @Produce json
// @Success 200
// @Router /metrics/ping [get]
func (ctrl *MetricController) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
