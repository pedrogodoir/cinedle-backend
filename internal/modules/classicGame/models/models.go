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
type ClassicGameGuess struct {
	Title        string `json:"title"`
	ReleaseDate  string `json:"releaseDate"`
	Budget       string `json:"budget"`
	TicketOffice string `json:"ticketOffice"`
	VoteAverage  string `json:"voteAverage"`
	Genres       string `json:"genres"`
	Companies    string `json:"companies"`
	Directors    string `json:"directors"`
	Actors       string `json:"actors"`
	Correct      bool   `json:"correct"`
}
