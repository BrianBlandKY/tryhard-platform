package model

type Service struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status int64  `json:"status"`
}
