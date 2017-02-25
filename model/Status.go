package schema

import ()

type Status struct {
	Id 				string	`json:"_id"`
	Status 			string  `json:"status"`
	Description 	string	`json:"description"`
}