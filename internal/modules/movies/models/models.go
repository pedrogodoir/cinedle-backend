package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Movie struct {
	ID           int             `json:"id"`
	Title        string          `json:"title"`
	Poster       string          `json:"poster"`
	ReleaseDate  time.Time       `json:"releaseDate"`
	Budget       decimal.Decimal `json:"budget"`
	TicketOffice decimal.Decimal `json:"ticketOffice"`
	VoteAverage  float32         `json:"voteAverage"`
}
type MovieSummary struct {
	ID    int    `json:"id" `
	Title string `json:"title"`
}

type MovieRes struct {
	Movie
	Genres    []Genre    `json:"genres"`
	Companies []Company  `json:"companies"`
	Directors []Director `json:"directors"`
	Actors    []Actor    `json:"actors"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Director struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Actor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
