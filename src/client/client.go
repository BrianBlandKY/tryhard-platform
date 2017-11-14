package main

import (
	"bufio"
	"log"
	"os"
	"time"
	cfg "tryhard-platform/src/config"
	msg "tryhard-platform/src/messenger"
	service "tryhard-platform/src/service"
)

type client struct {
	service.Service
}

func (c *client) scan() {
	stopActvityCh := make(chan bool)

	fakeActivity := func(msgCh <-chan msg.Command) {
		for {
			select {
			case m := <-msgCh:
				log.Printf("received msg %v \r\n", m)
				c.Reply(m)
			case _ = <-stopActvityCh:
				return
			}
		}
	}

	fakeActivityGenerator := func() {
		for {
			select {
			case _ = <-stopActvityCh:
				return
			default:
				time.Sleep(1 * time.Second)
				repMsg := &msg.Command{}
				err := c.Request(msg.Command{
					Service: "TEST_CLIENT",
					Action:  "FAKE_ACTIVITY",
					Data:    []byte("fake data"),
				}, repMsg)
				if err != nil {
					// panic(err)
				}
			}
		}
	}

	// Show some cool output?
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Bytes()
		c.Printf("command: %s \r\n", command)

		switch cmd := string(command[:len(command)]); cmd {
		case "print config":
			c.Printf("\r\n %v \r\n", cfg.WriteConfig(c.Config()))
		case "start":
			msgCh := c.Start()
			go fakeActivity(msgCh)
			go fakeActivityGenerator()
		case "stop":
			stopActvityCh <- true
			c.Stop()
		default:
			c.Println("unrecognized command", cmd)
		}
	}
}

func newClient(c cfg.Config) (s client) {
	s = client{
		Service: service.DefaultService(c),
	}
	return s
}
