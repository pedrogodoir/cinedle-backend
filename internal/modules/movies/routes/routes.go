package routes

import (
	"cinedle-backend/internal/modules/movies/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, h *handlers.MoviesHandler) {
	movies := route.Group("/movies")
	{
		movies.GET("/:id", h.GetMovieById)
		movies.GET("/summary/:title", h.GetMovieSummaryByTitle)
		movies.GET("/search/:title", h.GetMovieByTitle)

		//movies.POST("/", h.CreateMovie) #exemplo
	}
}
