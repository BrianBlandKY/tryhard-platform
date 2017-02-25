package messenger

// Command
type Command struct {
	Subject string      `json:"subject"`
	Reply   string      `json:"reply"`
	Data    interface{} `json:"data"`
}
