package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Movie struct {
	ID           int             `json:"id" gorm:"primaryKey;autoIncrement"`
	Title        string          `json:"title" gorm:"type:varchar(255);not null;uniqueIndex:idx_title_release_date"`
	Poster       string          `json:"poster" gorm:"type:text;not null"`
	ReleaseDate  time.Time       `json:"releaseDate" gorm:"type:timestamp;not null;uniqueIndex:idx_title_release_date"`
	Budget       decimal.Decimal `json:"budget" gorm:"type:numeric(12,2);not null"`
	TicketOffice decimal.Decimal `json:"ticketOffice" gorm:"type:numeric(12,2);not null"`
	VoteAverage  float32         `json:"voteAverage" gorm:"type:real;not null"`
}
type MovieRes struct {
	Movie
	Genres    []Genre    `json:"genres"`
	Companies []Company  `json:"companies"`
	Directors []Director `json:"directors"`
	Actors    []Actor    `json:"actors"`
}

type Genre struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(100);not null;unique"`
}
type Company struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(100);not null;unique"`
}
type Director struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(100);not null;unique"`
}
type Actor struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(100);not null;unique"`
}
