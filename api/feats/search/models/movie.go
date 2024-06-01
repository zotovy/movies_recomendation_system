package models

// TODO: move to movie package later

type MoviePreview struct {
	Id       int64  `json:"id" example:"123"`
	Title    string `json:"title" example:"Avatar"`
	Overview string `json:"overview" example:"A paraplegic Marine is dispatched to the moon"`
}
