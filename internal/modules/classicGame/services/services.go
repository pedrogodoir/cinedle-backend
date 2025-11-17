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
	ValidateGuess(id int, date string) (model_movie.MovieRes, models.ClassicGameGuess, error)
	GetTodaysClassicGame() (model_movie.MovieRes, error)
	GetClassicGameByDate(date time.Time) (models.ClassicGame, error)
	DrawMovie(date time.Time) int
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
func (s *classicGameService) ValidateGuess(movie_id int, date string) (model_movie.MovieRes, models.ClassicGameGuess, error) {
	movie_service := movie_service.NewMoviesService()
	guess, err := movie_service.GetMovieById(movie_id)
	if err != nil {
		return model_movie.MovieRes{}, models.ClassicGameGuess{}, err
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return model_movie.MovieRes{}, models.ClassicGameGuess{}, err
	}
	daysGame, err := s.GetClassicGameByDate(parsedDate)
	if err != nil {
		return model_movie.MovieRes{}, models.ClassicGameGuess{}, err
	}
	correct, err := movie_service.GetMovieById(daysGame.ID)
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
		return movie_service.GetMovieById(s.DrawMovie(time.Now()))
	}

	// Retorna o filme associado ao jogo clássico
	return movie_service.GetMovieById(classic_game.ID)
}

func (s *classicGameService) DrawMovie(date time.Time) int {
	movie_service := movie_service.NewMoviesService()

	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	fmt.Println("Sorteando filme do dia em:", date)

	// Buscar apenas IDs de filmes que ainda não estão na classic_games
	availableIDs, err := movie_service.GetAvailableClassicMovieIDs()
	if err != nil {
		fmt.Println("Erro ao buscar filmes disponíveis:", err)
		return -1
	}
	if len(availableIDs) == 0 {
		fmt.Println("Nenhum filme disponível para sortear")
		return -1
	}

	// Sortear um ID entre os disponíveis
	randomID := availableIDs[rand.Intn(len(availableIDs))]

	fmt.Println("Filme sorteado:", randomID)

	// Buscar informações completas do filme
	searchedMovie, err := movie_service.GetMovieById(randomID)
	if err != nil {
		return -1
	}

	// 4. Registrar como ClassicGame
	game := models.ClassicGame{
		ID:           randomID,
		Title:        searchedMovie.Title,
		Date:         date,
		TotalGuesses: 0,
	}

	created, err := s.CreateClassicGame(game)
	if err != nil {
		fmt.Println("Erro ao criar jogo clássico:", err)
		return -1
	}

	return created.ID
}

func (s *classicGameService) GetClassicGameByDate(date time.Time) (models.ClassicGame, error) {
	game, err := s.repo.GetClassicGameByDate(date)

	// Filme não encontrado para o dia. Sortear.
	if game.ID == 0 {
		draw_id := s.DrawMovie(date)
		return s.repo.GetClassicGameById(draw_id)
	}
	if err != nil {
		fmt.Println("Erro ao buscar classic game")
	}

	return s.repo.GetClassicGameByDate(date)
}
