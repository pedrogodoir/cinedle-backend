package scheduler

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// Inicia o cron do filme do dia
func StartFilmeDoDiaScheduler() *cron.Cron {
	c := cron.New(cron.WithSeconds())
	// Sortear filme às 16h
	c.AddFunc("0 0 16 * * *", func() {
		fmt.Println("Sorteando filme do dia em:", time.Now())
		// lógica de salvar no banco
	})
	c.Start()
	return c
}
