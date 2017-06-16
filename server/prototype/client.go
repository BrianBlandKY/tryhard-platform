package prototype

// Client Interface
type Client interface {
	ID() string          // internal ID for server management
	SetPartyCode(string) //
	PartyCode() string   //
	Done()               // close client and delete
	Write(Message)       // send a message to the client
	Listen()             // set client to listen for read/write messages
}
