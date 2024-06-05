package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @BasePath /api/v1

// GetMovieById godoc
// @Summary      Get movie by its id
// @Description  Responds with the movie
// @Tags         movie
// @Produce      json
// @Param        id path int true "movie id"
// @Success      200 {object} models.MovieDTO
// @Router       /movies/{id} [get]
func (c *MoviesController) GetMovieById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be int",
		})
		return
	}

	movie, err := c.repository.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if movie == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "movie not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, movie.ToDto())
}
