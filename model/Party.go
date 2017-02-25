package model
import (
	"time"
)

type Party struct {
	Id      	string 		`json:"_id"`
	Code    	string 		`json:"code"`
	Status  	Status		`json:"status"`
	App     	App 		`json:"app"`
	DateCreated time.Time 	`json:"date_created"`
	DateUpdated time.Time 	`db:"date_updated"`
}
