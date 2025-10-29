package handlers

import (
	services "cinedle-backend/internal/modules/posterGame/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PosterGameHandler struct {
	service services.PosterGameService
}

func NewPosterGameHandler() *PosterGameHandler {
	return &PosterGameHandler{
		service: services.NewPosterGameService(),
	}
}

func (h *PosterGameHandler) ValidateGuess(c *gin.Context) {
	movieID := c.Query("movie_id")
	defaultDate := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	date := c.DefaultQuery("date", defaultDate)
	iteration := c.DefaultQuery("iteration", "1")

	id, err := strconv.Atoi(movieID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	iter, err := strconv.Atoi(iteration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Iteração inválida"})
		return
	}

	res, err := h.service.ValidateGuess(id, date, iter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not implemented"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"res": res})

}
