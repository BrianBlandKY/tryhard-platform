package model

type Player struct {
	ID          string `json:"id"`
	PartyStatus Status `json:"party_status"`
	Status      Status `json:"status"`
	Signature   string `json:"signature"`
	Handle      string `json:"handle"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}
