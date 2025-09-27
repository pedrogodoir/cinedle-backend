package router

import (
	"cinedle-backend/internal/config"
	"cinedle-backend/internal/modules/movies/handlers"
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
	movieHandler := handlers.NewMoviesHandler(movieService)
	movie_router.Routes(r, movieHandler)

	r.Run(":" + cfg.Port)
}
