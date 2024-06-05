package models

type SimilarUsers struct {
	Distances    []float64 `json:"distances"`
	SimilarUsers []int     `json:"similar_users"`
}
