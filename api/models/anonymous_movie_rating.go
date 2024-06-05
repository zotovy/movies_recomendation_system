package models

type AnonymousMovieRating struct {
	MovieId int `json:"movie_id"`
	Rating  int `json:"rating"`
}
