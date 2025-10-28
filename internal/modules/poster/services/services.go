package service

import (
	movies_service "cinedle-backend/internal/modules/movies/services"
	"cinedle-backend/internal/modules/poster/models"
	repository "cinedle-backend/internal/modules/poster/repositories"
	image_manipualtor "cinedle-backend/internal/utils/image_manipulator"
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"os"
)

type PosterGameService interface {
	GetPosterGameById(id int) (models.PosterGame, error)
	//ValidateGuess(movie_id int, date string, iteration int) (models.PosterGame, error)
	generatePosterImages(movie_id int, date string) ([]image.Image, error)
}
type posterGameService struct {
	repo repository.PosterGameRepository
}

func NewPosterGameService() PosterGameService {
	return &posterGameService{
		repo: repository.NewPosterGameRepository(),
	}
}

func (s *posterGameService) GetPosterGameById(id int) (models.PosterGame, error) {
	return s.repo.GetPosterGameById(id)
}

// func (s *posterGameService) ValidateGuess(movie_id int, date string, iteration int) (models.PosterGame, error) {
// 	s.repo.GetPosterGameById(movie_id)
// }

/*Nessa função gera as imagens do pôster para um filme específico e salva no Poster Game no dia específico*/
func (s *posterGameService) generatePosterImages(movie_id int, date string) ([]image.Image, error) {
	movie, err := movies_service.NewMoviesService().GetMovieById(movie_id)
	if err != nil {
		return nil, err
	}
	/*TODO
	GERAR 9 IMAGENS, cada um com um retangulo novo revelado.
	*/
	baseImg, err := image_manipualtor.HandleImage(movie.Poster)
	if err != nil {
		return nil, err
	}
	imgs := []image.Image{}
	// A ideia aqui é gerar 9 imagens, cada um com um retangulo novo revelado.
	total_attempts := 9
	blockSize := 15
	rgba := image_manipualtor.ToRGBA(baseImg)
	rects := image_manipualtor.GetAllRects(rgba)
	/*suffle*/
	for i := range rects {
		j := rand.Intn(i + 1)
		rects[i], rects[j] = rects[j], rects[i]
	}
	for i := 1; i <= total_attempts-1; i++ {
		image_manipualtor.PixelateNRegions(rgba, rects[:i], blockSize)
		outFile, _ := os.Create(fmt.Sprintf("output_%d.jpg", i))
		defer outFile.Close()
		jpeg.Encode(outFile, rgba, &jpeg.Options{Quality: 90})
	}
	fmt.Println("Rectangles order:", rects)
	for i := 1; i <= total_attempts; i++ {
		imgs = append(imgs, image_manipualtor.PixelateNRegions(baseImg, rects[:i], blockSize))
	}

	return imgs, nil
}
