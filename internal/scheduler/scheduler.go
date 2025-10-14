package scheduler

import (
	"time"

	classic_game_service "cinedle-backend/internal/modules/classicGame/services"
	movie_service "cinedle-backend/internal/modules/movies/services"

	"github.com/robfig/cron/v3"
)

func StartFilmeDoDiaScheduler(svc_classic classic_game_service.ClassicGameService, svc_movie movie_service.MoviesService) *cron.Cron {
	c := cron.New()
	// Sortear filme do próximo dia meia noite

	c.AddFunc("0 0 * * *", func() {
		nextDay := time.Now().Add(24 * time.Hour)
		svc_classic.DrawMovie(nextDay)
	})

	c.Start()
	return c
}
