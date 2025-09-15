package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Essa struct é um exemplo de modelo, não o definido pelo nosso grande jordan
type Movie struct {
	ID           int
	Title        string
	ReleaseDate  *time.Time
	Budget       *decimal.Decimal
	TicketOffice *decimal.Decimal
	VoteAverage  *decimal.Decimal
}

type CreateMovie struct {
	Title        string
	ReleaseDate  *time.Time
	Budget       *decimal.Decimal
	TicketOffice *decimal.Decimal
	VoteAverage  *decimal.Decimal
}
