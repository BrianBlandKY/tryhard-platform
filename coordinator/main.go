package main

import (
	"fmt"
	mess "tryhard-platform/messenger"

	"log"
	"tryhard-platform/model"
)

func main() {
	nickelodeonServer := "nats://localhost:4222"

	h := newCoordinator(nickelodeonServer)
	h.Connect()
	h.Listen()
	log.Println("Done.")
}

// Remove references to NATS
type coordinator struct {
	address   string
	messenger mess.Messenger
	sub       mess.Subscription
	subCh     chan mess.Command
}

func (c *coordinator) Connect() {
	err := c.messenger.Connect(c.address, "CARNIVAL_ID")
	if err != nil {
		panic(err)
	}

	c.subscribe()
}

func (c *coordinator) subscribe() {
	ch := make(chan mess.Command)
	key := c.messenger.Key(mess.PARTY, "*")
	recvSub, err := c.messenger.SubscribeChan(key, ch)
	if err != nil {
		fmt.Printf("error creating channel %v", err)
	}

	c.sub = recvSub
	c.subCh = ch
}

func (c *coordinator) processAction(cmd mess.Command, party model.Party) (err error) {
	switch cmd.Action {
	case mess.JOIN:
		c.join(cmd, party)
	case mess.DISBAND:
		c.disband(cmd, party)
	default:
		err = fmt.Errorf("invalid action %v", cmd.Action)
	}
	return
}

// JOIN
func (c *coordinator) join(cmd mess.Command, party model.Party) {
	c.messenger.Reply(cmd)
}

// DISBAND
func (c *coordinator) disband(cmd mess.Command, party model.Party) {
	c.messenger.Reply(cmd)
}

func (c *coordinator) Listen() {
	log.Println("Listening")

	for {
		select {
		case m := <-c.subCh:
			log.Println("received message", m)

			var party model.Party
			err := m.Deserialize(&party)
			if err != nil {
				log.Println("failed to deserialize message", err)
				break
			}

			err = c.processAction(m, party)
			if err != nil {
				log.Printf("process error %v \n\r", err)
				break
			}
		}
	}
}

func newCoordinator(address string) (c coordinator) {
	c = coordinator{
		address:   address,
		messenger: mess.NewMessenger(),
	}
	return
}
