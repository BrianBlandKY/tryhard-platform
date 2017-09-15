package model

type Player struct {
	ID    string `json:"id"`
	Alias string `json:"alias"`
	//Status      mess.Status `json:"status"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}
