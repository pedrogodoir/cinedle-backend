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

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
)

type PosterGameService interface {
	GetPosterGameById(id int) (models.PosterGame, error)
	GetPosterGameByDateAndIteration(date string, iteration int) (models.PosterGame, error)
	createPosterGame(posterGame models.PosterGameCreate) ([]models.PosterGame, error)
	UpdatePosterGame(posterGame models.PosterGame) error
	ValidateGuess(movie_id int, date string, iteration int) (models.PosterGameGuessRes, error)
	generatePosterImages(movie_id int) ([]image.Image, error)
	saveGeneratedImages(movie_id int, date string) ([]string, error)
	GetPosterImageByDateAndIteration(date string, iteration int) (models.PosterImageImage, error)
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

func (s *posterGameService) DrawMovie(date string) int {
	movie_service := movies_service.NewMoviesService()

	fmt.Println("Sorteando filme para poster game do dia em:", date)

	// Buscar apenas IDs de filmes que ainda não estão na classic_games
	availableIDs, err := movie_service.GetAvailablePosterMovieIDs()
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

	// Registrar como Poster Game create
	s.createPosterGame(
		models.PosterGameCreate{
			MovieID: randomID,
			Date:    date,
		},
	)

	return randomID
}

func (s *posterGameService) ValidateGuess(movie_id int, date string, iteration int) (models.PosterGameGuessRes, error) {

	if iteration < 1 || iteration > 9 {
		return models.PosterGameGuessRes{}, fmt.Errorf("iteration %d fora do intervalo válido (1–9)", iteration)
	}
	posterGame, err := s.GetPosterGameByDateAndIteration(date, iteration)
	if err != nil {
		return models.PosterGameGuessRes{}, err
	}

	var next models.PosterGame

	if posterGame.MovieID == movie_id { // ACERTOU O FILME

		posterFinal, err := s.GetPosterGameByDateAndIteration(date, 9)
		if err != nil {
			return models.PosterGameGuessRes{}, err
		}
		return models.PosterGameGuessRes{
			CurrentImage: posterFinal.ImageURL,
			Correct:      true,
			NextImage:    "",
		}, nil

	}
	// ERROU O FILME
	next, err = s.GetPosterGameByDateAndIteration(date, iteration+1)
	if err != nil {
		return models.PosterGameGuessRes{}, err
	}

	return models.PosterGameGuessRes{
		CurrentImage: posterGame.ImageURL,
		Correct:      false,
		NextImage:    next.ImageURL,
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
	blockSize := 25
	rects := image_manipualtor.GetAllRects(baseImg)

	/*suffle*/
	for i := range rects {
		j := rand.Intn(i + 1)
		rects[i], rects[j] = rects[j], rects[i]
	}
	for i := 1; i <= total_attempts; i++ {
		copy := image_manipualtor.ToRGBA(baseImg) // cópia da imagem original
		imgs = append(imgs, image_manipualtor.PixelateNRegions(copy, rects[:i], blockSize))
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
	if iteration == 0 {
		return nil, fmt.Errorf("no images generated")
	}
	// // abre conexão com imagekit
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
		buf.Reset()

		iteration--

	}
	return results, nil
}

/* nunca se cria somente um posterGame,  cria um para cada imagem borrada (iteração)*/
func (s *posterGameService) createPosterGame(posterGame models.PosterGameCreate) ([]models.PosterGame, error) {
	imageURLs, err := s.saveGeneratedImages(posterGame.MovieID, posterGame.Date)
	if err != nil {
		return nil, err
	}
	var res []models.PosterGame
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
		res = append(res, models.PosterGame{
			ID:        id,
			MovieID:   posterGame.MovieID,
			Name:      posterGame.Name,
			Date:      posterGame.Date,
			ImageURL:  url,
			Iteration: i,
		})

		i--
	}
	return res, nil
}

func (s *posterGameService) GetPosterGameByDateAndIteration(date string, iteration int) (models.PosterGame, error) {

	game, err := s.repo.GetPosterGameByDateAndIteration(date, iteration)

	// Filme não encontrado para o dia. Sortear.
	if game.ID == 0 {
		fmt.Println("Nenhum poster game encontrado para a data e iteração. Sorteando novo filme.")
		s.DrawMovie(date)
		return s.repo.GetPosterGameByDateAndIteration(date, iteration)
	}
	if err != nil {
		fmt.Println("Erro ao buscar poster game: ", err)
	}
	return s.repo.GetPosterGameByDateAndIteration(date, iteration)
}

func (s *posterGameService) GetPosterImageByDateAndIteration(date string, iteration int) (models.PosterImageImage, error) {
	if iteration < 1 || iteration > 9 {
		iteration = 1
	}
	game, err := s.GetPosterGameByDateAndIteration(date, iteration)
	fmt.Println("GAME: ", game)
	return models.PosterImageImage{ImageURL: game.ImageURL}, err
}
