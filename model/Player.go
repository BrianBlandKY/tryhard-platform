package schema

import ()

type Player struct {
	Id 				string      `json:"_id"`
	PartyStatus 	Status	    `json:"party_status"`
	Status 		    Status	    `json:"status"`
	Signature 		string 	    `json:"signature"`
	Handle 			string	    `json:"handle"`
	DateCreated 	string	    `json:"date_created"`
	DateUpdated 	string      `json:"date_updated"`
}