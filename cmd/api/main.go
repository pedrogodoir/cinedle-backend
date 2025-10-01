package main

import (
	"cinedle-backend/internal/config"
	"cinedle-backend/internal/database"
	classicService "cinedle-backend/internal/modules/classicGame/services"
	movieService "cinedle-backend/internal/modules/movies/services"
	"cinedle-backend/internal/router"
	"cinedle-backend/internal/scheduler"
)

func main() {
	config.LoadConfig()
	database.GetDBPool()

	// cria o service
	classicService := classicService.NewClassicGameService()
	movieService := movieService.NewMoviesService()

	// injeta no scheduler
	c := scheduler.StartFilmeDoDiaScheduler(classicService, movieService)
	defer c.Stop()

	router.Run()
}
