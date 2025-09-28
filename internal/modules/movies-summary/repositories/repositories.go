// repository/movies_repository.go
package repository

import (
	"cinedle-backend/internal/database"
	"cinedle-backend/internal/modules/movies-summary/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type MoviesSummaryRepository interface {
	GetMovieSummaryByTitle(title string) ([]models.MovieSummary, error)
}

// moviesRepo é a implementação concreta do repositório
type moviesSummaryRepo struct {
	db *pgx.Conn
}

// NewMoviesRepository retorna uma instância do repositório
func NewMoviesSummaryRepository() *moviesSummaryRepo {
	return &moviesSummaryRepo{
		db: database.GetDB(),
	}
}

func (r *moviesSummaryRepo) GetMovieSummaryByTitle(title string) ([]models.MovieSummary, error) {
	fmt.Printf("Repository: Fetching movie summary for title: %s\n", title)
	var movieRes []models.MovieSummary

	rows, err := r.db.Query(database.GetCtx(),
		`SELECT * from search_movie where title LIKE '%' || $1 || '%'`,
		title,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.MovieSummary
		if err := rows.Scan(&m.ID, &m.Title); err != nil {
			return nil, err
		}
		movieRes = append(movieRes, m)
	}

	return movieRes, nil
}
