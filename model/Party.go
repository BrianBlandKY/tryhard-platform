package model

type Party struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	//Status mess.Status `json:"status"`
	Player []Player `json:"players"`
}
