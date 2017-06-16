package data

import "time"

// Party Model
type Party struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
}
