// service/movies_service.go
package service

import (
	"cinedle-backend/internal/modules/movies/models"
	repository "cinedle-backend/internal/modules/movies/repositories"
)

// MoviesService define os métodos do service
type MoviesService interface {
	GetFullMovieById(id int) (models.MovieRes, error)
}

// moviesService é a implementação concreta
type moviesService struct {
	repo repository.MoviesRepository
}

// NewMoviesService cria uma instância do service
func NewMoviesService(repo repository.MoviesRepository) MoviesService {
	return &moviesService{
		repo: repo,
	}
}

func (s *moviesService) GetFullMovieById(id int) (models.MovieRes, error) {
	return s.repo.GetFullMovieById(id)
}
