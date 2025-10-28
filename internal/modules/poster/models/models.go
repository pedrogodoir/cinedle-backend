package models

type PosterGame struct {
	ID        int    `json:"id"`
	MovieID   int    `json:"movie_id"`
	Name      string `json:"name"`
	Iteration int    `json:"iteration"`
	ImageURL  string `json:"image_ur l"`
}
type PosterGameRes struct {
	PosterGame PosterGame `json:"poster_game"`
	Correct    bool       `json:"correct"`
}
