package socket

import (
    "time"
)
type player struct {
    client   		*client   
    Id				string		`json:"id"`
    Username		string		`json:"username"`
    Status			string	    `json:"status"`
    CreationDate 	time.Time   `json:"date_created"`
}
