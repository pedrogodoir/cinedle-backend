package routes

import (
	"cinedle-backend/internal/modules/posterGame/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, h *handlers.PosterGameHandler) {
	posterGames := route.Group("/poster-games")
	{
		posterGames.GET("/guess", h.ValidateGuess)
	}
}
