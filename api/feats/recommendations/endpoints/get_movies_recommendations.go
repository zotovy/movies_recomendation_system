package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @BasePath /api/v1

// GetMoviesRecommendations godoc
// @Summary      Get movies recommendations
// @Description  Responds with the list of movies, which is the best recommendations for given user
// @Tags         recommendations
// @Accept  json
// @Produce      json
// @Param        userId query int true "user id"
// @Param        limit  query int false "query limit"
// @Success      200 {array} models.MoviePreview
// @Router       /recommendations/movies [post]
func (c *Controller) GetMoviesRecommendations(ctx *gin.Context) {
	userIdStr := ctx.Query("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
	}

	amountStr := ctx.Query("limit")
	amount := 100
	if amountStr != "" {
		amount, err = strconv.Atoi(amountStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user id",
			})
		}
	}

	if amount < 0 {
		amount = 1
	}

	recommendations, err := c.repository.GetRecommendations(userId, amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to get recommendations",
			"messages": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, recommendations)
}
