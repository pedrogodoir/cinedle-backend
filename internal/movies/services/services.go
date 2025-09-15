package services

import (
	"cinedle-backend/internal/movies/domain/models"
	"cinedle-backend/internal/movies/domain/repository"
)

type Services struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *Services {
	return &Services{
		repos: repos,
	}
}

func (s *Services) GetAll() ([]models.Movie, error) {
	return s.repos.Movie.GetAll()
}

func (s *Services) Add(newMovie models.Movie) (int32, error) {
	repoReq := models.Movie{
		Title:        newMovie.Title,
		ReleaseDate:  newMovie.ReleaseDate,
		Budget:       newMovie.Budget,
		TicketOffice: newMovie.TicketOffice,
		VoteAverage:  newMovie.VoteAverage,
	}
	return s.repos.Movie.Add(repoReq)
}
