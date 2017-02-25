package model

import (
	"time"
)

type App struct {
	Id          string    `json:"_id"`
	Name        string    `json:"name"`
	IsEnabled   bool      `json:"is_enabled"`
	DateCreated time.Time `json:"date_created"`
}
