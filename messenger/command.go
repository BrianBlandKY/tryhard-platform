package messenger

// Command
type Command struct {
	Service string `json:"service"`
	Reply   string `json:"-"`
	Data    []byte `json:"data"`
}
