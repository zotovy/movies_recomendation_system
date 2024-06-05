package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type MoviePreview struct {
	Id          int64   `json:"id" example:"123"`
	Title       string  `json:"title" example:"Avatar"`
	Tagline     *string `json:"tagline" example:"A paraplegic Marine is dispatched to the moon"`
	VoteAverage float32 `json:"voteAverage" db:"vote_average" example:"7.8"`
}

type Movie struct {
	Id               int
	Budget           int64
	Homepage         *string
	OriginalLanguage string
	OriginalTitle    string
	Overview         string
	Popularity       float32
	ReleaseDate      pgtype.Date
	Revenue          int
	Runtime          float32
	Status           string
	Tagline          *string
	Title            string
	VoteAverage      float32
	VoteCount        int

	Genres              *[]Genre
	Keywords            *[]Keyword
	ProductionCompanies *[]Company
	Countries           *[]Country
	Languages           *[]Language
	Cast                *[]Cast
	Crew                *[]Crew
}

type MovieDTO struct {
	Id               int     `json:"id" example:"123"`
	Budget           int64   `json:"budget" example:"10000"`
	Homepage         *string `json:"homepage" example:"http://www.sinistermovie.com/"`
	OriginalLanguage string  `json:"originalLanguage" example:"English"`
	OriginalTitle    string  `json:"originalTitle" example:"Avatar"`
	Overview         string  `json:"overview" example:"A paraplegic Marine is dispatched to the moon"`
	Popularity       float32 `json:"popularity" example:"321.52"`
	ReleaseDate      string  `json:"releaseDate" example:"2006-10-06"`
	Revenue          int     `json:"revenue" example:"14821658"`
	Runtime          float32 `json:"runtime" example:"123"`
	Status           string  `json:"status" example:"Released"`
	Tagline          *string `json:"tagline" example:"A paraplegic Marine is dispatched to the moon"`
	Title            string  `json:"title" example:"Avatar"`
	VoteAverage      float32 `json:"voteAverage" example:"7.8"`
	VoteCount        int     `json:"voteCount" example:"32132"`

	Genres              *[]Genre    `json:"genres"`
	Keywords            *[]Keyword  `json:"keywords"`
	ProductionCompanies *[]Company  `json:"productionCompanies"`
	Countries           *[]Country  `json:"countries"`
	Languages           *[]Language `json:"languages"`
	Cast                *[]CastDTO  `json:"cast"`
	Crew                *[]CrewDTO  `json:"crew"`
}

func (movie *Movie) ToDto() MovieDTO {
	return MovieDTO{
		Id:                  movie.Id,
		Budget:              movie.Budget,
		Homepage:            movie.Homepage,
		OriginalLanguage:    movie.OriginalLanguage,
		OriginalTitle:       movie.OriginalTitle,
		Overview:            movie.Overview,
		Popularity:          movie.Popularity,
		ReleaseDate:         movie.ReleaseDate.Time.Format(time.RFC3339),
		Revenue:             movie.Revenue,
		Runtime:             movie.Runtime,
		Status:              movie.Status,
		Tagline:             movie.Tagline,
		Title:               movie.Title,
		VoteAverage:         movie.VoteAverage,
		VoteCount:           movie.VoteCount,
		Genres:              movie.Genres,
		Keywords:            movie.Keywords,
		ProductionCompanies: movie.ProductionCompanies,
		Countries:           movie.Countries,
		Languages:           movie.Languages,
		Cast:                CastArrayToDto(movie.Cast),
		Crew:                CrewArrayToDto(movie.Crew),
	}
}
