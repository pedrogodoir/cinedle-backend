// service/movies_service.go
package service

import (
	"cinedle-backend/internal/modules/movies/models"
	repository "cinedle-backend/internal/modules/movies/repositories"
	"cinedle-backend/internal/utils"
	cache_movie "cinedle-backend/internal/utils/cache/cacheMovie"
	"strings"
)

// MoviesService define os métodos do service
type MoviesService interface {
	GetMovieById(id int) (models.MovieRes, error)
	GetMovieByTitle(title string) (models.MovieRes, error)
	GetMovieSummaryByTitle(title string) ([]models.MovieSummary, error)
	GetMovieCount() (int, error)
	GetAvailableClassicMovieIDs() ([]int, error)
	GetAvailablePosterMovieIDs() ([]int, error)
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

// package-level cache instance so we don't recreate it on every request
var movieCache = cache_movie.NewCacheMovie()

func (s *moviesService) GetMovieById(id int) (models.MovieRes, error) {
	// primeiro usa a cache (instância única em nível de pacote)
	movie, hit := movieCache.GetMovieCache(id)
	if !hit {

		var err error
		movie, err = s.repo.GetMovieById(id)
		if err != nil {
			return models.MovieRes{}, err
		}
		movieCache.SetMovieCache(id, movie)
	}
	return movie, nil
}

func (s *moviesService) GetAvailableClassicMovieIDs() ([]int, error) {
	return s.repo.GetAvailableClassicMovieIDs()
}

func (s *moviesService) GetAvailablePosterMovieIDs() ([]int, error) {
	return s.repo.GetAvailablePosterMovieIDs()
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
