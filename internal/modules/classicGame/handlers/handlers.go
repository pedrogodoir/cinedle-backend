package handlers

import (
	"net/http"
	"strconv"

	"cinedle-backend/internal/modules/classicGame/models"
	services "cinedle-backend/internal/modules/classicGame/services"

	"github.com/gin-gonic/gin"
)

type ClassicGameHandler struct {
	s services.ClassicGameService
}

func NewClassicGameHandler() *ClassicGameHandler {
	return &ClassicGameHandler{
		s: services.NewClassicGameService(),
	}
}
func (h *ClassicGameHandler) GetClassicGameById(c *gin.Context) {
	// Pega o ID da URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	classicGame, err := h.s.GetClassicGameById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if classicGame.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "\"Jogo clássico\" não encontrado"})
		return
	}

	c.JSON(http.StatusOK, classicGame)
}
func (h *ClassicGameHandler) CreateClassicGame(c *gin.Context) {
	var newGame models.ClassicGame
	if err := c.ShouldBindJSON(&newGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	classicGame, err := h.s.CreateClassicGame(newGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, classicGame)
}
func (h *ClassicGameHandler) GetAllClassicGames(c *gin.Context) {
	games, err := h.s.GetAllClassicGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}
func (h *ClassicGameHandler) UpdateClassicGame(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var updatedGame models.ClassicGame
	if err := c.ShouldBindJSON(&updatedGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	err = h.s.UpdateClassicGame(id, updatedGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "\"Jogo clássico\" atualizado com sucesso"})
}
func (h *ClassicGameHandler) DeleteClassicGame(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	err = h.s.DeleteClassicGame(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "\"Jogo clássico\" deletado com sucesso"})
}
func (h *ClassicGameHandler) ValidateGuess(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	movie, res, err := h.s.ValidateGuess(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"movie": movie, "res ": res})
}
func (h *ClassicGameHandler) GetTodayClassicGame(c *gin.Context) {
	classicGame, err := h.s.GetTodaysClassicGame()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classicGame)
}
