package prototype

type PartySource interface {
    Get(string) (Party)
    Host(Server) (Party)
}
