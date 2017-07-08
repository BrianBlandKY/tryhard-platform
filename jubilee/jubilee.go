package main

import (
	"fmt"
	"tryhard-platform-old/data"
	mess "tryhard-platform/messenger"
	"tryhard-platform/model"

	"encoding/json"
	"log"
)

func main() {
	nickelodeonServer := "nats://localhost:4222"

	h := newCarnival(nickelodeonServer)
	h.Connect()
	h.Listen()
	log.Println("Done.")
}

// Remove references to NATS
type carnival struct {
	address   string
	messenger mess.Messenger
	recvSub   mess.Subscription
	recvChan  chan mess.Command
}

func (c *carnival) Connect() {
	c.messenger.Register(
		model.Shindig{},
	)

	err := c.messenger.Connect(c.address, "CARNIVAL_ID")
	if err != nil {
		panic(err)
	}

	c.bind()
}

func (c *carnival) bind() {
	ch := make(chan mess.Command)
	recvSub, err := c.messenger.BindRecvChan(data.PARTY, ch)
	if err != nil {
		fmt.Printf("error creating channel %v", err)
	}

	c.recvSub = recvSub
	c.recvChan = ch
}

func (c *carnival) Listen() {
	log.Println("Listening")

	for {
		select {
		case m := <-c.recvChan:
			if m.Action == data.JOIN {
				log.Println("received message", m)
				c.messenger.Reply()
			}
		}
	}
}

func (c *carnival) decode(msg []byte, output interface{}) {
	err := json.Unmarshal(msg, output)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func newCarnival(address string) carnival {
	c := carnival{
		address:   address,
		messenger: mess.NewMessenger(),
	}
	return c
}
