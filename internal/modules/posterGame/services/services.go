package service

import (
	"bytes"
	"cinedle-backend/internal/config"
	movies_service "cinedle-backend/internal/modules/movies/services"
	"cinedle-backend/internal/modules/posterGame/models"
	repository "cinedle-backend/internal/modules/posterGame/repositories"
	image_manipualtor "cinedle-backend/internal/utils/image_manipulator"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
)

type PosterGameService interface {
	GetPosterGameById(id int) (models.PosterGame, error)
	GetPosterGameByDateAndIteration(date string, iteration int) (models.PosterGame, error)
	createPosterGame(posterGame models.PosterGame) ([]int, error)
	UpdatePosterGame(posterGame models.PosterGame) error
	ValidateGuess(movie_id int, date string, iteration int) (models.PosterGameRes, error)
	generatePosterImages(movie_id int) ([]image.Image, error)
	saveGeneratedImages(movie_id int, date string) ([]string, error)
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
func (s *posterGameService) UpdatePosterGame(posterGame models.PosterGame) error {
	return s.repo.UpdatePosterGame(posterGame)
}

func (s *posterGameService) ValidateGuess(movie_id int, date string, iteration int) (models.PosterGameRes, error) {
	posterGame, err := s.repo.GetPosterGameByDateAndIteration(date, iteration)

	if err != nil {
		return models.PosterGameRes{}, err
	}
	correct := false
	if posterGame.MovieID == movie_id {
		correct = true
	}
	next, err := s.repo.GetPosterGameByDateAndIteration(date, iteration+1)
	if err != nil {
		return models.PosterGameRes{}, err
	}

	return models.PosterGameRes{
		PosterGame: posterGame,
		Correct:    correct,
		NextImage:  next.ImageURL,
	}, nil
}

/*Nessa função gera as imagens do pôster para um filme específico e salva no Poster Game no dia específico*/
func (s *posterGameService) generatePosterImages(movie_id int) ([]image.Image, error) {
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
func (s *posterGameService) saveGeneratedImages(movie_id int, date string) ([]string, error) {
	imgs, err := s.generatePosterImages(movie_id)
	results := []string{}
	if err != nil {
		return nil, err
	}
	iteration := len(imgs)
	// abre conexão com imagekit
	key := config.LoadConfig().ImageKitKey
	client := imagekit.NewClient(
		option.WithPrivateKey(key),
	)
	var buf bytes.Buffer
	// Save the generated images
	for _, img := range imgs {
		err := jpeg.Encode(&buf, img, nil)
		if err != nil {
			log.Fatal(err)
		}

		var file io.Reader = &buf
		response, err := client.Files.Upload(context.TODO(), imagekit.FileUploadParams{
			File:     file,
			FileName: fmt.Sprintf("%s-%d.jpg", date, iteration),
		})
		if err != nil {
			log.Fatalf("Erro ao fazer upload do arquivo: %v", err)
		}
		results = append(results, response.URL)
		iteration--
		buf.Reset()

	}
	return results, nil
}

/* nunca se cria sómente um posterGame,  cria um para cada iteração*/
func (s *posterGameService) createPosterGame(posterGame models.PosterGame) ([]int, error) {
	imageURLs, err := s.saveGeneratedImages(posterGame.MovieID, posterGame.Date)
	if err != nil {
		return nil, err
	}
	var ids []int
	i := len(imageURLs)
	for _, url := range imageURLs {
		id, err := s.repo.CreatePosterGame(models.PosterGame{
			MovieID:   posterGame.MovieID,
			Name:      posterGame.Name,
			Date:      posterGame.Date,
			ImageURL:  url,
			Iteration: i,
		})
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
		i--
	}
	return ids, nil
}

func (s *posterGameService) GetPosterGameByDateAndIteration(date string, iteration int) (models.PosterGame, error) {
	return s.repo.GetPosterGameByDateAndIteration(date, iteration)
}
