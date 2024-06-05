package endpoints

import (
	"app/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// @BasePath /api/v1

// GetMoviesRecommendations godoc
// @Summary      Get movies recommendations
// @Description  Responds with the list of movies, which is the best recommendations for given user
// @Tags         recommendations
// @Accept  json
// @Produce      json
// @Param        request body []models.AnonymousMovieRating true "user ratings"
// @Success      200 {array} models.MoviePreview
// @Router       /recommendations/movies [post]
func (c *Controller) GetMoviesRecommendations(ctx *gin.Context) {
	jsonBody, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return
	}

	var ratings *[]models.AnonymousMovieRating
	err = json.Unmarshal(jsonBody, &ratings)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to marshal request body",
		})
		return
	}

	if len(*ratings) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Provide at least one movie rating",
		})
		return
	}

	recommendations, err := c.repository.GetRecommendations(ratings)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to get recommendations",
			"messages": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, recommendations)
}
