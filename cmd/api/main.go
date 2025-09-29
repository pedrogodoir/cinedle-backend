package main

import (
	"cinedle-backend/internal/config"
	"cinedle-backend/internal/database"
	"cinedle-backend/internal/router"
)

func main() {
	config.LoadConfig()
	database.GetDBPool()
	//router.Run()
	// Example query to test the connection
	// Close the database connection when done

	router.Run()
}
