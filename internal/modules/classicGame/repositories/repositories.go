// repository/movies_repository.go
package repository

import (
	"cinedle-backend/internal/database"
	"cinedle-backend/internal/modules/classicGame/models"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClassicGameRepository interface {
	GetClassicGameById(id int) (models.ClassicGame, error)
	CreateClassicGame(game models.ClassicGame) (int, error)
	GetAllClassicGames() ([]models.ClassicGame, error)
	UpdateClassicGame(id int, game models.ClassicGame) error
	DeleteClassicGame(id int) error
	GetClassicGameByDate(date time.Time) (models.ClassicGame, error)
}

type classicGameRepo struct {
	db *pgxpool.Pool
}

func NewClassicGameRepository() ClassicGameRepository {
	return &classicGameRepo{
		db: database.GetDBPool(),
	}
}

func (r *classicGameRepo) GetClassicGameById(id int) (models.ClassicGame, error) {
	var query string = `SELECT movie_id, title, date, total_guesses FROM classic_games WHERE movie_id = $1;`
	var movieRes models.ClassicGame

	row := r.db.QueryRow(database.GetCtx(),
		query,
		id,
	)

	err := row.Scan(
		&movieRes.ID,
		&movieRes.Title,
		&movieRes.Date,
		&movieRes.TotalGuesses,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			// Retorna struct vazia se não encontrar
			return models.ClassicGame{}, nil
		}
		return models.ClassicGame{}, err
	}

	return movieRes, nil
}
func (r *classicGameRepo) CreateClassicGame(game models.ClassicGame) (int, error) {
	var query string = `INSERT INTO classic_games (movie_id, title, date, total_guesses) VALUES ($1, $2, $3, $4) RETURNING movie_id;`
	var newID int
	err := r.db.QueryRow(database.GetCtx(),
		query,
		game.ID,
		game.Title,
		game.Date,
		game.TotalGuesses,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}
func (r *classicGameRepo) GetAllClassicGames() ([]models.ClassicGame, error) {
	var query string = `SELECT movie_id, title, date, total_guesses FROM classic_games;`
	rows, err := r.db.Query(database.GetCtx(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	games := []models.ClassicGame{}
	for rows.Next() {
		var game models.ClassicGame
		if err := rows.Scan(&game.ID, &game.Title, &game.Date, &game.TotalGuesses); err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}
func (r *classicGameRepo) UpdateClassicGame(id int, game models.ClassicGame) error {
	var query string = `UPDATE classic_games SET title = $1, date = $2, total_guesses = $3 WHERE movie_id = $4;`
	_, err := r.db.Exec(database.GetCtx(),
		query,
		game.Title,
		game.Date,
		game.TotalGuesses,
		id,
	)

	return err
}
func (r *classicGameRepo) DeleteClassicGame(id int) error {
	var query string = `DELETE FROM classic_games WHERE movie_id = $1;`
	_, err := r.db.Exec(database.GetCtx(),
		query,
		id,
	)
	return err
}
func (r *classicGameRepo) GetClassicGameByDate(date time.Time) (models.ClassicGame, error) {
	var data string = strings.Split(date.String(), " ")[0]
	var query string = `SELECT movie_id, title, date, total_guesses FROM classic_games WHERE date = $1`
	var movieRes models.ClassicGame
	row := r.db.QueryRow(database.GetCtx(),
		query,
		data,
	)
	err := row.Scan(
		&movieRes.ID,
		&movieRes.Title,
		&movieRes.Date,
		&movieRes.TotalGuesses,
	)

	if err != nil {

		return models.ClassicGame{ID: 1, Title: "Inception", Date: time.Now(), TotalGuesses: 0}, nil
	}
	return movieRes, nil
}
