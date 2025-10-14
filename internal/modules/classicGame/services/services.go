// service/movies_service.go
package service

import (
	"cinedle-backend/internal/modules/classicGame/models"
	repository "cinedle-backend/internal/modules/classicGame/repositories"
	model_movie "cinedle-backend/internal/modules/movies/models"
	movie_service "cinedle-backend/internal/modules/movies/services"
	"cinedle-backend/internal/utils"
	"fmt"
	"math/rand"
	"time"
)

type ClassicGameService interface {
	GetClassicGameById(id int) (models.ClassicGame, error)
	CreateClassicGame(game models.ClassicGame) (models.ClassicGame, error)
	GetAllClassicGames() ([]models.ClassicGame, error)
	UpdateClassicGame(id int, game models.ClassicGame) error
	DeleteClassicGame(id int) error
	ValidateGuess(id int) (model_movie.MovieRes, models.ClassicGameGuess, error)
	GetTodaysClassicGame() (model_movie.MovieRes, error)
}
type classicGameService struct {
	repo repository.ClassicGameRepository
}

func NewClassicGameService() ClassicGameService {
	return &classicGameService{
		repo: repository.NewClassicGameRepository(),
	}
}

func (s *classicGameService) GetClassicGameById(id int) (models.ClassicGame, error) {
	return s.repo.GetClassicGameById(id)
}
func (s *classicGameService) CreateClassicGame(game models.ClassicGame) (models.ClassicGame, error) {
	var createdGame models.ClassicGame
	id, err := s.repo.CreateClassicGame(game)
	if err != nil {
		return models.ClassicGame{}, err
	}
	createdGame.ID = id
	createdGame.Title = game.Title
	createdGame.Date = game.Date
	createdGame.TotalGuesses = game.TotalGuesses
	return createdGame, nil
}
func (s *classicGameService) GetAllClassicGames() ([]models.ClassicGame, error) {
	return s.repo.GetAllClassicGames()
}
func (s *classicGameService) UpdateClassicGame(id int, game models.ClassicGame) error {
	return s.repo.UpdateClassicGame(id, game)
}
func (s *classicGameService) DeleteClassicGame(id int) error {
	return s.repo.DeleteClassicGame(id)
}
func (s *classicGameService) ValidateGuess(id int) (model_movie.MovieRes, models.ClassicGameGuess, error) {
	movie_service := movie_service.NewMoviesService()
	guess, err := movie_service.GetMovieById(id)
	if err != nil {
		return model_movie.MovieRes{}, models.ClassicGameGuess{}, err
	}
	correct, err := s.GetTodaysClassicGame()
	if err != nil {
		return model_movie.MovieRes{}, models.ClassicGameGuess{}, err
	}
	res := utils.CompareMovies(guess, correct)

	return guess, res, nil
}
func (s *classicGameService) GetTodaysClassicGame() (model_movie.MovieRes, error) {
	movie_service := movie_service.NewMoviesService()
	classic_game, err := s.repo.GetClassicGameByDate(time.Now())

	if err != nil {
		fmt.Println("ERRO")
		return model_movie.MovieRes{}, err
	}

	// Nenhum jogo encontrado para hoje
	if classic_game.ID == 0 {
		return movie_service.GetMovieById(s.drawMovie(movie_service, time.Now()))
	}

	// Retorna o filme associado ao jogo clássico
	return movie_service.GetMovieById(classic_game.ID)
}

func (s *classicGameService) drawMovie(svc_movie movie_service.MoviesService, date time.Time) int {

	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	fmt.Println("Sorteando filme do dia em:", date)

	movie_count, err := svc_movie.GetMovieCount()

	if err != nil {
		fmt.Println("Erro ao buscar quantidade de filmes")
		return -1
	}

	var randomId int
	for {
		randomId = rand.Intn(movie_count-1) + 1 // sorteia entre 1 e movie_count
		searchedGame, err := s.GetClassicGameById(randomId)
		if searchedGame.ID == 0 {
			fmt.Println("ID vago:", randomId)
			break
		}

		if err != nil {
			fmt.Println("Erro ao buscar classicGame:")
			break
		}
		fmt.Println("ID já existe, sorteando outro:", randomId)
	}

	searchedMovie, err := svc_movie.GetMovieById(randomId)
	if err != nil {
		return -1
	}

	game := models.ClassicGame{
		ID:           randomId,
		Title:        searchedMovie.Title,
		Date:         date,
		TotalGuesses: 0,
	}

	created, err := s.CreateClassicGame(game)
	if err != nil {
		fmt.Println("Erro ao criar filme: ", err)
		return -1
	}

	return created.ID
}
