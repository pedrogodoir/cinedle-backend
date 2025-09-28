package handlers

import (
	"net/http"

	services "cinedle-backend/internal/modules/movies-summary/services"

	"github.com/gin-gonic/gin"
)

type MoviesSummaryHandler struct {
	service services.MoviesSummaryService
}

func NewMoviesSummaryHandler(service services.MoviesSummaryService) *MoviesSummaryHandler {
	return &MoviesSummaryHandler{
		service: service,
	}
}
func (h *MoviesSummaryHandler) GetMovieSummaryByTitle(c *gin.Context) {
	// Pega o title da URL
	titleParam := c.Param("title")

	movies, err := h.service.GetMovieSummaryByTitle(titleParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}
