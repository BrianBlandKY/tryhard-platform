package main

import (
	"bufio"
	"log"
	"os"
	config "tryhard-platform/config"
	mess "tryhard-platform/messenger"
	"tryhard-platform/model"
)

type client struct {
	mess.Node
}

// Would be cool if we could override the console in Node
func (c *client) connect() {
	c.SetHandler(c.handler)
	c.SetProcessor(c.processor)
	c.SetConsole(c.scan)
	c.Run()
}

func (c *client) handler(n *mess.Node) {
	for {
		select {}
	}
}

func (c *client) processor() {
	for {
		select {}
	}
}

func (c *client) disconnect() {
	c.Close()
}

func (c *client) join() {
	cmd := mess.Command{
		Service: mess.Services.Party,
		Action:  mess.JOIN,
	}

	cmd.Serialize(model.Party{
		Code: "PARTY_CODE",
		ID:   "TESTID",
	})

	var resCommand mess.Command
	err := c.Messenger.Request(cmd, &resCommand)
	if err != nil {
		log.Printf("error %v", err)
	}
}

func (c *client) leave() {
	cmd := mess.Command{
		Service: mess.Services.Party,
		Action:  mess.DISBAND,
	}

	cmd.Serialize(model.Party{
		Code: "PARTY_CODE",
	})

	var resCommand mess.Command
	err := c.Messenger.Request(cmd, &resCommand)
	if err != nil {
		panic(err)
	}
}

func (c *client) scan() {
	// Show some cool output?
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Bytes()
		log.Printf("command: %s \r\n", command)

		switch cmd := string(command[:len(command)]); cmd {
		case "join":
			go c.join()
		default:
			log.Printf("unrecognized command %s \r\n", cmd)
		}
	}
}

func newClient(cfg config.Config) (s client) {
	s = client{
		Node: mess.DefaultNode(cfg),
	}
	return s
}
