// repository/movies_repository.go
package repository

import (
	"cinedle-backend/internal/database"
	"cinedle-backend/internal/modules/movies/models"

	"github.com/jackc/pgx/v5"
)

type MoviesRepository interface {
	GetMovieById(id int) (models.MovieRes, error)
}

// moviesRepo é a implementação concreta do repositório
type moviesRepo struct {
	db *pgx.Conn
}

// NewMoviesRepository retorna uma instância do repositório
func NewMoviesRepository() MoviesRepository {
	return &moviesRepo{
		db: database.GetDB(),
	}
}

func (r *moviesRepo) GetMovieById(id int) (models.MovieRes, error) {
	var query string = `SELECT 
    m.id, m.title, m.release_date, m.budget, m.ticket_office, m.vote_average,
    COALESCE(JSON_AGG(DISTINCT g) FILTER (WHERE g.id IS NOT NULL), '[]') AS genres,
    COALESCE(JSON_AGG(DISTINCT c) FILTER (WHERE c.id IS NOT NULL), '[]') AS companies,
    COALESCE(JSON_AGG(DISTINCT d) FILTER (WHERE d.id IS NOT NULL), '[]') AS directors,
    COALESCE(JSON_AGG(DISTINCT a) FILTER (WHERE a.id IS NOT NULL), '[]') AS actors
		FROM movies m
		LEFT JOIN movie_genre mg ON mg.movie_id = m.id
		LEFT JOIN genres g ON g.id = mg.genre_id
		LEFT JOIN movie_company mc ON mc.movie_id = m.id
		LEFT JOIN companies c ON c.id = mc.company_id
		LEFT JOIN movie_director md ON md.movie_id = m.id
		LEFT JOIN directors d ON d.id = md.director_id
		LEFT JOIN movie_actor ma ON ma.movie_id = m.id
		LEFT JOIN actors a ON a.id = ma.actor_id
		WHERE m.id = $1
		GROUP BY m.id;`
	var movieRes models.MovieRes

	row := r.db.QueryRow(database.GetCtx(),
		query,
		id,
	)

	err := row.Scan(
		&movieRes.ID,
		&movieRes.Title,
		&movieRes.ReleaseDate,
		&movieRes.Budget,
		&movieRes.TicketOffice,
		&movieRes.VoteAverage,
		&movieRes.Genres,
		&movieRes.Companies,
		&movieRes.Directors,
		&movieRes.Actors,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			// Retorna struct vazia se não encontrar
			return models.MovieRes{}, nil
		}
		return models.MovieRes{}, err
	}

	return movieRes, nil
}
