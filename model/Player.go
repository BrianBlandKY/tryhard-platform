package model

type Player struct {
	ID          string `json:"id"`
	DeviceID    string `json:"device_id"`
	Alias       string `json:"alias"`
	Status      Status `json:"status"`
	Signature   string `json:"signature"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}
