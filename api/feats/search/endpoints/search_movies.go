package endpoints

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// @BasePath /api/v1

// SearchMovies             godoc
// @Summary      Search movies by query
// @Description  Responds with the list of all books, which satisfies query
// @Tags         search
// @Produce      json
// @Param        q query string true "query to search"
// @Success      200 {array} models.MoviePreview
// @Router       /search/movies [get]
func (c *SearchController) SearchMovies(ctx *gin.Context) {
	q := ctx.Query("q")

	if strings.TrimSpace(q) == "" {
		ctx.JSON(400, gin.H{
			"error": "empty query",
		})
		return
	}

	movies, err := c.repository.SearchMovies(q)

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
