package prototype

type Party interface {
    Code()              string
    Status()            string
    Write(msg Message)
    Listen()
    Reset()   
    Broadcast(msg Message)

    // subject
    // Add(c Client)
    // Drop(c Client)
    // Err(err error)
}
