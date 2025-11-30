package handlers

import (
	services "cinedle-backend/internal/modules/posterGame/services"
	"cinedle-backend/internal/utils"
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

	_, err = utils.ValidateDate(date)
	if err != nil {
		// Retornamos o erro exato que a função auxiliar gerou
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.ValidateGuess(id, date, iter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"res": res})
}

func (h *PosterGameHandler) GetPosterImageByDateAndIteration(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	iteration := c.DefaultQuery("iteration", "1")

	// Validação da Iteração
	iter, err := strconv.Atoi(iteration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Iteração inválida"})
		return
	}

	_, err = utils.ValidateDate(date)
	if err != nil {
		// Retornamos o erro exato que a função auxiliar gerou
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.GetPosterImageByDateAndIteration(date, iter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"res": res})
}

func (h *PosterGameHandler) GetPosterGameByDateAndIteration(c *gin.Context) {
	defaultDate := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	date := c.DefaultQuery("date", defaultDate)

	iteration := c.DefaultQuery("iteration", "9")
	iter, err := strconv.Atoi(iteration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Iteração inválida"})
		return
	}

	_, err = utils.ValidateDate(date)
	if err != nil {
		// Retornamos o erro exato que a função auxiliar gerou
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.GetPosterGameByDateAndIteration(date, iter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"res": res})
}
