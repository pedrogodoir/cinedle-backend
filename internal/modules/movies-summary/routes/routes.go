package routes

import (
	"cinedle-backend/internal/modules/movies-summary/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, h *handlers.MoviesSummaryHandler) {
	movies := route.Group("/search")
	{
		movies.GET("/:title", h.GetMovieSummaryByTitle)
		//movies.POST("/", h.CreateMovie) #exemplo
	}
}
