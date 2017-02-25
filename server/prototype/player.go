package prototype

import (
    "time"
)

type Player interface {
    GetClient()        	Client
    SetClient(Client)
    GetStatus()		string
    SetStatus(string)
    // Username()      	string    	
    // Type()		string		
    DateCreated()   	time.Time   
}
