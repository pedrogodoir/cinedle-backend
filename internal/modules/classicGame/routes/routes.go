package routes

import (
	"cinedle-backend/internal/modules/classicGame/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, h *handlers.ClassicGameHandler) {
	classicGames := route.Group("/classic-games")
	{
		classicGames.GET("/:id", h.GetClassicGameById)
		classicGames.POST("/", h.CreateClassicGame)
		classicGames.GET("/", h.GetAllClassicGames)
		classicGames.PUT("/:id", h.UpdateClassicGame)
		classicGames.DELETE("/:id", h.DeleteClassicGame)
	}
}
