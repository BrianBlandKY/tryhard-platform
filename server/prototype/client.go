package prototype

import ()

type Client interface {
	Id()        			string  // internal id for server management
	SetPartyCode(string)		 	// 
	GetPartyCode()			string  //
	Done()              			// close client and delete
	Write(Message)     				// send a message to the client
	Listen()        	    		// set client to listen for read/write messages
}