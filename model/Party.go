package model

type Party struct {
	ID     string   `json:"id"`
	Code   string   `json:"code"`
	Status Status   `json:"status"`
	Player []Player `json:"players"`
}
