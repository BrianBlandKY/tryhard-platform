package prototype

type Message interface {
    GetClient()         Client
    SetClient(Client)
    Command()           string
    PartyCode()         string  
    Data()		[]byte
}
