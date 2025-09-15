package repository

import "cinedle-backend/internal/movies/domain/models"

type Repositories struct {
	Movie interface {
		GetAll() ([]models.Movie, error)
		//GetAll() ([]interface{}, error) #exemplo de metodo sem tipagem, o interface funciona para qualquer tipo
		Add(newMovie models.Movie) (int32, error)
	}
}

func New() *Repositories {
	return &Repositories{}
}
