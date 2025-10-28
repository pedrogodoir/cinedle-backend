package repository

import (
	"cinedle-backend/internal/database"
	"cinedle-backend/internal/modules/poster/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PosterGameRepository interface {
	GetPosterGameById(id int) (models.PosterGame, error)
	//feio, eu sei
	GetPosterGameByDateAndIteration(date string, iteration int) (models.PosterGame, error)
}

// moviesRepo é a implementação concreta do repositório
type posterGameRepo struct {
	db *pgxpool.Pool
}

// NewPosterGameRepository retorna uma instância do repositório
func NewPosterGameRepository() PosterGameRepository {
	return &posterGameRepo{
		db: database.GetDBPool(),
	}
}

func (r *posterGameRepo) GetPosterGameById(id int) (models.PosterGame, error) {
	var posterGame models.PosterGame
	query := `SELECT id, movie_id, name, image_url FROM poster_games WHERE id = $1`
	err := r.db.QueryRow(database.GetCtx(), query, id).Scan(
		&posterGame.ID,
		&posterGame.MovieID,
		&posterGame.Name,
		&posterGame.ImageURL,
	)
	if err != nil {
		return models.PosterGame{ID: 0}, err
	}
	return posterGame, nil
}

func (r *posterGameRepo) GetPosterGameByDateAndIteration(date string, iteration int) (models.PosterGame, error) {
	var posterGame models.PosterGame
	query := `SELECT id, movie_id, name, image_url FROM poster_games WHERE date = $1 AND iteration = $2`
	err := r.db.QueryRow(database.GetCtx(), query, date, iteration).Scan(
		&posterGame.ID,
		&posterGame.MovieID,
		&posterGame.Name,
		&posterGame.ImageURL,
	)
	if err != nil {
		return models.PosterGame{ID: 0}, err
	}
	return posterGame, nil
}
