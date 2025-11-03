package repository

import (
	"cinedle-backend/internal/database"
	"cinedle-backend/internal/modules/posterGame/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PosterGameRepository interface {
	GetPosterGameById(id int) (models.PosterGame, error)
	CreatePosterGame(posterGame models.PosterGame) (int, error)
	//feio, eu sei
	GetPosterGameByDateAndIteration(date string, iteration int) (models.PosterGame, error)
	UpdatePosterGame(posterGame models.PosterGame) error
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

/*isso aqui tá pegando como array para impedir que de erro caso não ache dentro do bd*/
func (r *posterGameRepo) GetPosterGameByDateAndIteration(date string, iteration int) (models.PosterGame, error) {
	var posterGame []models.PosterGame
	query := `SELECT id, movie_id, name, image_url FROM poster_games WHERE date = $1 AND iteration = $2`
	rows, err := r.db.Query(database.GetCtx(), query, date, iteration)
	if err != nil {
		return models.PosterGame{ID: 0}, err
	}
	defer rows.Close()
	for rows.Next() {
		var pg models.PosterGame
		err := rows.Scan(
			&pg.ID,
			&pg.MovieID,
			&pg.Name,
			&pg.ImageURL,
		)
		if err != nil {
			return models.PosterGame{ID: 0}, err
		}
		posterGame = append(posterGame, pg)
	}
	if len(posterGame) == 0 {
		return models.PosterGame{ID: 0}, nil
	}

	return posterGame[0], nil
}

func (r *posterGameRepo) UpdatePosterGame(posterGame models.PosterGame) error {
	query := `UPDATE poster_games SET movie_id = $1, name = $2, image_url = $3, iteration = $4 WHERE id = $5`
	_, err := r.db.Exec(database.GetCtx(), query,
		posterGame.MovieID,
		posterGame.Name,
		posterGame.ImageURL,
		posterGame.Iteration,
		posterGame.ID,
	)
	return err
}

func (r *posterGameRepo) CreatePosterGame(posterGame models.PosterGame) (int, error) {
	var id int
	query := `INSERT INTO poster_games (movie_id, name, iteration, date, image_url) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(database.GetCtx(), query,
		posterGame.MovieID,
		posterGame.Name,
		posterGame.Iteration,
		posterGame.Date,
		posterGame.ImageURL,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
