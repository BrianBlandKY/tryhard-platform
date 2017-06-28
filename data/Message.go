package data

// Command Struct
type Command struct {
	Action  string `json:"action"`
	Service string `json:"service"`
	User    User   `json:"user"`
}

// Message Struct
type Message struct {
	Command `json:"command"`
}

// PartyMessage Struct
type PartyMessage struct {
	Command `json:"command"`
	Party   `json:"party"`
}
