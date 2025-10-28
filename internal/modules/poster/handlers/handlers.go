package handlers

import (
	services "cinedle-backend/internal/modules/poster/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MoviesHandler struct {
	service services.PosterGameService
}

func NewMoviesHandler(service services.PosterGameService) *MoviesHandler {
	return &MoviesHandler{
		service: service,
	}
}
func (h *MoviesHandler) GetMovieById(c *gin.Context) {
	// Pega o ID da URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	movie, err := h.service.GetPosterGameById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if movie.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poster Game não encontrado"})
		return
	}

	c.JSON(http.StatusOK, movie)
}
