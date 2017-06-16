package prototype

// Message Interface
type Message interface {
	Client() Client
	SetClient(Client)
	Command() string
	PartyCode() string
	Data() []byte
}
