package endpoints

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// @BasePath /api/v1

// AutoCompleteMovies godoc
// @Summary      Autocomplete movies title
// @Description  Responds with the list of all books, which title satisfies query
// @Tags         search
// @Produce      json
// @Param        q query string true "query to search"
// @Success      200 {array} models.MoviePreview
// @Router       /search/movies/autocomplete [get]
func (c *SearchController) AutoCompleteMovies(ctx *gin.Context) {
	q := ctx.Query("q")

	if strings.TrimSpace(q) == "" {
		ctx.JSON(400, gin.H{
			"error": "empty query",
		})
		return
	}

	movies, err := c.repository.AutoCompleteMovies(q)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if movies == nil {
		return
	}

	ctx.JSON(200, movies)
}
