package models

type AnonymousMovieRating struct {
	MovieId int     `json:"movie_id"`
	Rating  float64 `json:"rating"`
}
