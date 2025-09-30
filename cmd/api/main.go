package main

import (
	"cinedle-backend/internal/config"
	"cinedle-backend/internal/database"
	"cinedle-backend/internal/router"
	"cinedle-backend/internal/scheduler"
)

func main() {
	config.LoadConfig()
	database.GetDBPool()

	c := scheduler.StartFilmeDoDiaScheduler()
	defer c.Stop()

	router.Run()
}
