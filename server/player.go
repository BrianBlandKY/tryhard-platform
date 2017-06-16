package server

import (
	"time"
)

type player struct {
	client       *client
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Status       string    `json:"status"`
	CreationDate time.Time `json:"date_created"`
}
