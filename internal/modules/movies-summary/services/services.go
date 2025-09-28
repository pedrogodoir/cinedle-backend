// service/movies_service.go
package service

import (
	"cinedle-backend/internal/modules/movies-summary/models"
	repository "cinedle-backend/internal/modules/movies-summary/repositories"
	"strings"
)

// MoviesSummaryService define os métodos do service
type MoviesSummaryService interface {
	GetMovieSummaryByTitle(title string) ([]models.MovieSummary, error)
}

// moviesSummaryService é a implementação concreta
type moviesSummaryService struct {
	repo repository.MoviesSummaryRepository
}

// NewMoviesSummaryService cria uma instância do service
func NewMoviesSummaryService(repo repository.MoviesSummaryRepository) MoviesSummaryService {
	return &moviesSummaryService{
		repo: repo,
	}
}

func (s *moviesSummaryService) GetMovieSummaryByTitle(title string) ([]models.MovieSummary, error) {
	var t = strings.Trim(title, " ")
	return s.repo.GetMovieSummaryByTitle(t)
}
