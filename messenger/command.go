package messenger

import (
	"encoding/json"
)

// Command
type Command struct {
	Service string `json:"service"`
	Action  string `json:"action"`
	Reply   string `json:"-"`
	Data    []byte `json:"data"`
}

func (c *Command) Deserialize(output interface{}) (err error) {
	err = json.Unmarshal(c.Data, &output)
	return
}

func (c *Command) Serialize(input interface{}) (err error) {
	res, err := json.Marshal(input)
	c.Data = res
	return
}
