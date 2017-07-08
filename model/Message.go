package model

type Message struct {
	Action string `json:"action"`
}

type PartyMessage struct {
	Message
	Party `json:"party"`
	Users []User `json:"users"`
}

type HeartbeatMessage struct {
	Heartbeat `json:"heartbeat"`
}
