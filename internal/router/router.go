package router

import (
	"cinedle-backend/internal/config"
	summary_handler "cinedle-backend/internal/modules/movies-summary/handlers"
	summary_repo "cinedle-backend/internal/modules/movies-summary/repositories"
	summary_router "cinedle-backend/internal/modules/movies-summary/routes"
	summary_service "cinedle-backend/internal/modules/movies-summary/services"
	movie_handler "cinedle-backend/internal/modules/movies/handlers"
	movie_repo "cinedle-backend/internal/modules/movies/repositories"
	movie_router "cinedle-backend/internal/modules/movies/routes"
	movie_service "cinedle-backend/internal/modules/movies/services"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.LoadConfig()
	r := gin.Default()

	//setup routes
	movieRepo := movie_repo.NewMoviesRepository()
	movieService := movie_service.NewMoviesService(movieRepo)
	movieHandler := movie_handler.NewMoviesHandler(movieService)
	movie_router.Routes(r, movieHandler)

	summaryRepo := summary_repo.NewMoviesSummaryRepository()
	summaryService := summary_service.NewMoviesSummaryService(summaryRepo)
	summaryHandler := summary_handler.NewMoviesSummaryHandler(summaryService)
	summary_router.Routes(r, summaryHandler)

	r.Run(":" + cfg.Port)
}
