package data

// Message Struct
type Message struct {
	Client  Client `json:"client"`
	ID      string `json:"id"`
	Command string `json:"command"`
}

// MessageAdapter Struct
type MessageAdapter struct {
	Message `json:"message"`
}

// HeartbeatMessage Struct
type HeartbeatMessage struct {
	Message   `json:"message"`
	Heartbeat `json:"heartbeat"`
}

// PartyMessage Struct
type PartyMessage struct {
	Message `json:"message"`
	Player  `json:"player"`
	Party   `json:"party"`
}

// PartyBroadcastMessage Struct
type PartyBroadcastMessage struct {
	Message `json:"message"`
	Party   `json:"party"`
	Players []Player `json:"players"`
}
