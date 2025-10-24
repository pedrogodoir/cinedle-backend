// service/movies_service.go
package service

import (
	"cinedle-backend/internal/modules/movies/models"
	repository "cinedle-backend/internal/modules/movies/repositories"
	"cinedle-backend/internal/utils"
	"strings"
)

// MoviesService define os métodos do service
type MoviesService interface {
	GetMovieById(id int) (models.MovieRes, error)
	GetMovieByTitle(title string) (models.MovieRes, error)
	GetMovieSummaryByTitle(title string) ([]models.MovieSummary, error)
	GetMovieCount() (int, error)
}

// moviesService é a implementação concreta
type moviesService struct {
	repo repository.MoviesRepository
}

// NewMoviesService cria uma instância do service
func NewMoviesService() MoviesService {
	return &moviesService{
		repo: repository.NewMoviesRepository(),
	}
}

func (s *moviesService) GetMovieById(id int) (models.MovieRes, error) {
	return s.repo.GetMovieById(id)
}

func (s *moviesService) GetMovieByTitle(title string) (models.MovieRes, error) {
	t := utils.ToTitle(title)
	t = strings.Trim(t, " ")
	return s.repo.GetMovieByTitle(t)
}
func (s *moviesService) GetMovieSummaryByTitle(title string) ([]models.MovieSummary, error) {
	t := utils.ToTitle(title)
	t = strings.Trim(t, " ")
	return s.repo.GetMovieSummaryByTitle(t)
}

func (s *moviesService) GetMovieCount() (int, error) {
	return s.repo.GetMovieCount()
}
