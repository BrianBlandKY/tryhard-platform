package prototype

import (
	"time"
)

// Player Interface
type Player interface {
	Client() Client
	SetClient(Client)
	Status() string
	SetStatus(string)
	// Username()      	string
	// Type()		    string
	DateCreated() time.Time
}
