package models

import (
	"time"
)

type ClassicGame struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	TotalGuesses int       `json:"total_guesses"`
}
