// service/movies_service.go
package service

import (
	"cinedle-backend/internal/modules/classicGame/models"
	repository "cinedle-backend/internal/modules/classicGame/repositories"
)

type ClassicGameService interface {
	GetClassicGameById(id int) (models.ClassicGame, error)
	CreateClassicGame(game models.ClassicGame) (int, error)
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
func (s *classicGameService) CreateClassicGame(game models.ClassicGame) (int, error) {
	return s.repo.CreateClassicGame(game)
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
