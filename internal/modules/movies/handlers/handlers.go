package handlers

import (
	"net/http"
	"strconv"

	services "cinedle-backend/internal/modules/movies/services"

	"github.com/gin-gonic/gin"
)

type MoviesHandler struct {
	service services.MoviesService
}

func NewMoviesHandler(service services.MoviesService) *MoviesHandler {
	return &MoviesHandler{
		service: service,
	}
}
func (h *MoviesHandler) GetFullMovieById(c *gin.Context) {
	// Pega o ID da URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	movie, err := h.service.GetFullMovieById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if movie.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
		return
	}

	c.JSON(http.StatusOK, movie)
}
