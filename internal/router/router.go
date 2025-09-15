package router

import (
	"cinedle-backend/internal/config"
	movie_router "cinedle-backend/internal/movies/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config")
	}
	r := gin.Default()

	//setup routes
	movie_router.Routes(r)

	r.Run(":" + cfg.Port)
}
