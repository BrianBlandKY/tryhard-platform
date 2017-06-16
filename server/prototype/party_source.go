package prototype

// PartySource Interface
type PartySource interface {
	Generate(string) Party
	Host(Server) Party
}
