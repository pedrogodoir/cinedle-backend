package main

import (
	"cinedle-backend/internal/database"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func queryData(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "INSERT INTO movie (id, title) VALUES ($1, $2)", 1, "Inception")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Title: %s", id, title)
	}
}

func main() {
	db := database.New()
	queryData(db.GetConnection())
	//router.Run()
	// Example query to test the connection
	// Close the database connection when done
	defer db.Close()
}
