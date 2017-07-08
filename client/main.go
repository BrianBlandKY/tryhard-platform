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
		Service: model.PARTY,
		Action:  model.JOIN,
		Data:    []byte("TEST DATA"),
	}
	var resCommand mess.Command
	err := c.messenger.Request(cmd, &resCommand)
	if err != nil {
		panic(err)
	}

	log.Println("got a response", resCommand)
}

func (c *client) leave() {

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
