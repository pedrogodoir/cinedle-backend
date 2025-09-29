package router

import (
	"cinedle-backend/internal/config"
	classic_game_handler "cinedle-backend/internal/modules/classicGame/handlers"
	classic_game_router "cinedle-backend/internal/modules/classicGame/routes"
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
	classicGameHandler := classic_game_handler.NewClassicGameHandler()
	classic_game_router.Routes(r, classicGameHandler)

	r.Run(":" + cfg.Port)
}
