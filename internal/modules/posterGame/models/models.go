package models

type PosterGame struct {
	ID        int    `json:"id"`
	MovieID   int    `json:"movie_id"`
	Name      string `json:"name"`
	Iteration int    `json:"iteration"`
	Date      string `json:"date"`
	ImageURL  string `json:"image_url"`
}
type PosterGameRes struct {
	PosterGame PosterGame `json:"poster_game"`
	Correct    bool       `json:"correct"`
	NextImage  string     `json:"next_image,omitempty"`
}
