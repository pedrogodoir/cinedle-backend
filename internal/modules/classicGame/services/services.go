// service/movies_service.go
package service

import (
	"cinedle-backend/internal/modules/classicGame/models"
	repository "cinedle-backend/internal/modules/classicGame/repositories"
)

type ClassicGameService interface {
	GetClassicGameById(id int) (models.ClassicGame, error)
	CreateClassicGame(game models.ClassicGame) (models.ClassicGame, error)
	GetAllClassicGames() ([]models.ClassicGame, error)
	UpdateClassicGame(id int, game models.ClassicGame) error
	DeleteClassicGame(id int) error
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
