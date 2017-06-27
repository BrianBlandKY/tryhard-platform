package data

import "time"

// Player Model
type Player struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
}
