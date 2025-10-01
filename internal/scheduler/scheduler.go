package scheduler

import (
	"fmt"
	"math/rand"
	"time"

	"cinedle-backend/internal/modules/classicGame/models"
	classic_game_service "cinedle-backend/internal/modules/classicGame/services"
	movie_service "cinedle-backend/internal/modules/movies/services"

	"github.com/robfig/cron/v3"
)

func StartFilmeDoDiaScheduler(svc_classic classic_game_service.ClassicGameService, svc_movie movie_service.MoviesService) *cron.Cron {
	c := cron.New()
	// Sortear filme todo dia meia noite
	c.AddFunc("0 0 * * *", func() {
		fmt.Println("Sorteando filme do dia em:", time.Now())

		movie_count, err := svc_movie.GetMovieCount()

		if err != nil {
			fmt.Println("Erro ao buscar quantidade de filmes")
			return
		}

		var randomId int

		for {
			randomId = rand.Intn(movie_count-1) + 1 // sorteia entre 1 e movie_count
			searchedGame, err := svc_classic.GetClassicGameById(randomId)
			fmt.Println("Oq veio:", searchedGame)

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

		// Data do próximo dia
		tomorrow := time.Now().AddDate(0, 0, 1)

		game := models.ClassicGame{
			ID:           randomId,
			Title:        "Filme sorteado",
			Date:         tomorrow,
			TotalGuesses: 0,
		}

		created, err := svc_classic.CreateClassicGame(game)
		if err != nil {
			fmt.Println("Erro ao criar jogo:", err)
			return
		}

		fmt.Println("Jogo criado:", created)
	})
	c.Start()
	return c
}
