package utils

import (
	model_classic_game "cinedle-backend/internal/modules/classicGame/models"
	model_movie "cinedle-backend/internal/modules/movies/models"
	"fmt"
	"strings"
)

// capitalize the first letter of each word in a string
func ToTitle(s string) string {
	words := strings.Fields(strings.ToLower(s))
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(words, " ")
}

func CompareMovies(correct, guess model_movie.MovieRes) model_classic_game.ClassicGameGuess {
	var cg model_classic_game.ClassicGameGuess

	// Title
	if guess.Title == correct.Title {
		cg.Title = "correct"
	} else {
		cg.Title = "incorrect"
	}

	// ReleaseDate
	if guess.ReleaseDate.Equal(correct.ReleaseDate) {
		cg.ReleaseDate = "correct"
	} else {
		cg.ReleaseDate = "incorrect"
	}

	// Budget
	if guess.Budget.Equal(correct.Budget) {
		cg.Budget = "correct"
	} else {
		cg.Budget = "incorrect"
	}

	// TicketOffice
	if guess.TicketOffice.Compare(correct.TicketOffice) == 0 {
		cg.TicketOffice = "correct"
	} else {
		cg.TicketOffice = "incorrect"
		fmt.Println(guess.TicketOffice, correct.TicketOffice)
	}

	// VoteAverage
	if guess.VoteAverage == correct.VoteAverage {
		cg.VoteAverage = "correct"
	} else {
		cg.VoteAverage = "incorrect"
	}

	// Genres
	cg.Genres = statusSliceByIDStrict(extractIDsFromGenres(guess.Genres), extractIDsFromGenres(correct.Genres))

	// Companies
	cg.Companies = statusSliceByIDStrict(extractIDsFromCompanies(guess.Companies), extractIDsFromCompanies(correct.Companies))

	// Directors
	cg.Directors = statusSliceByIDStrict(extractIDsFromDirectors(guess.Directors), extractIDsFromDirectors(correct.Directors))

	// Actors
	cg.Actors = statusSliceByIDStrict(extractIDsFromActors(guess.Actors), extractIDsFromActors(correct.Actors))

	// overall correct
	cg.Correct = cg.Title == "correct" &&
		cg.ReleaseDate == "correct" &&
		cg.Budget == "correct" &&
		cg.TicketOffice == "correct" &&
		cg.VoteAverage == "correct" &&
		cg.Genres == "correct" &&
		cg.Companies == "correct" &&
		cg.Directors == "correct" &&
		cg.Actors == "correct"

	return cg
}

func statusSliceByIDStrict(guessIDs, correctIDs []int) string {
	if len(correctIDs) == 0 && len(guessIDs) == 0 {
		return "correct"
	}
	if equalIntSlices(guessIDs, correctIDs) {
		return "correct"
	}
	overlap := countOverlapIDs(guessIDs, correctIDs)
	if overlap > 0 {
		return "parcial"
	}
	return "incorrect"
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	count := map[int]int{}
	for _, v := range a {
		count[v]++
	}
	for _, v := range b {
		if count[v] == 0 {
			return false
		}
		count[v]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

func countOverlapIDs(a, b []int) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	set := map[int]struct{}{}
	for _, v := range b {
		set[v] = struct{}{}
	}
	ov := 0
	for _, v := range a {
		if _, ok := set[v]; ok {
			ov++
		}
	}
	return ov
}

// extractors assumem structs com ID e Name
func extractIDsFromGenres(gs []model_movie.Genre) []int {
	out := make([]int, 0, len(gs))
	for _, g := range gs {
		out = append(out, g.ID)
	}
	return out
}
func extractIDsFromCompanies(cs []model_movie.Company) []int {
	out := make([]int, 0, len(cs))
	for _, c := range cs {
		out = append(out, c.ID)
	}
	return out
}
func extractIDsFromDirectors(ds []model_movie.Director) []int {
	out := make([]int, 0, len(ds))
	for _, d := range ds {
		out = append(out, d.ID)
	}
	return out
}
func extractIDsFromActors(as []model_movie.Actor) []int {
	out := make([]int, 0, len(as))
	for _, a := range as {
		out = append(out, a.ID)
	}
	return out
}
