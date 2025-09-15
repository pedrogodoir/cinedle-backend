package repositories

import (
	"cinedle-backend/internal/database/schema"
	"cinedle-backend/internal/movies/domain/models"

	"errors"

	"gorm.io/gorm"

	"cinedle-backend/internal/database"
)

type Movie struct {
	movies []models.Movie
	db     database.DB
}

func New(db database.DB) *Movie {
	return &Movie{movies: make([]models.Movie, 0), db: db}
}

func (m *Movie) GetAll() []models.Movie {

	return m.movies
}

func (m *Movie) Add(newMovie models.Movie) (int, error) {
	movie := &schema.Movie{
		Title:        newMovie.Title,
		ReleaseDate:  newMovie.ReleaseDate,
		Budget:       newMovie.Budget,
		TicketOffice: newMovie.TicketOffice,
		VoteAverage:  newMovie.VoteAverage,
	}
	response := gorm.G[schema.Movie](m.db.GetConnection()).Create(m.db.GetContext(), movie)
	m.movies = append(m.movies, newMovie)

	return movie.ID, errors.New(response.Error())
}
