package main

import (
	"cinedle-backend/internal/database"
	"cinedle-backend/internal/database/schema"
)

func main() {
	db := database.New()
	schema.MigrateAll(db.GetConnection())
}
