package main

import (
	"bufio"
	"log"
	"os"
	mess "tryhard-platform/messenger"
	"tryhard-platform/model"
)

type client struct {
	messenger mess.Messenger
}

func (c *client) connect(url, id string) {
	err := c.messenger.Connect(url, id)
	if err != nil {
		panic(err)
	}
}

func (c *client) join() {
	cmd := mess.Command{
		Service: mess.PARTY,
		Action:  mess.JOIN,
	}

	cmd.Serialize(model.Party{
		Code: "PARTY_CODE",
		ID:   "TESTID",
	})

	var resCommand mess.Command
	err := c.messenger.Request(cmd, &resCommand)
	if err != nil {
		panic(err)
	}
}

func (c *client) leave() {
	cmd := mess.Command{
		Service: mess.PARTY,
		Action:  mess.DISBAND,
	}

	cmd.Serialize(model.Party{
		Code: "PARTY_CODE",
	})

	var resCommand mess.Command
	err := c.messenger.Request(cmd, &resCommand)
	if err != nil {
		panic(err)
	}
}

func (c *client) scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Bytes()
		log.Printf("command: %s \r\n", command)

		switch cmd := string(command[:len(command)]); cmd {
		case "join":
			c.join()
		case "leave":
			c.leave()
		default:
			log.Printf("unrecognized command %s \r\n", cmd)
		}
	}
}

func newClient() *client {
	return &client{
		messenger: mess.NewMessenger(),
	}
}

func main() {
	c := newClient()
	c.connect("nats://localhost:4222", "TEST_CLIENT")

	log.Println("Scanning...")
	c.scan()
}

// log.Println("Raw message", rawMsg)
// log.Println("Raw message (string)", string(rawMsg[:len(rawMsg)]))
