package main

import (
	"cinedle-backend/internal/config"
	"cinedle-backend/internal/database"
	repository "cinedle-backend/internal/modules/movies/repositories"
	"cinedle-backend/internal/router"
	"fmt"
)

func main() {
	config.LoadConfig()
	database.GetDB()
	movi_repo := repository.NewMoviesRepository()
	movie, err := movi_repo.GetFullMovieById(1)
	if err != nil {
		fmt.Println("Error fetching movies:", err)
		return
	}
	fmt.Println(movie.Actors[0].Name)
	//router.Run()
	// Example query to test the connection
	// Close the database connection when done

	router.Run()
	defer database.CloseDB()
}
