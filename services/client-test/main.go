package main

import (
	"bufio"
	"log"
	"os"

	"time"
	"tryhard-platform/data"

	"encoding/json"

	nats "github.com/nats-io/go-nats"
)

type client struct {
	nc *nats.Conn
	ec *nats.EncodedConn
}

func (c *client) connect(url string) {
	nc, err := nats.Connect(url)

	if err != nil {
		log.Panic(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Panic(err)
	}

	c.nc = nc
	c.ec = ec

	log.Println("Connected")
}

func (c *client) join() {
	log.Println("Joining party..")

	partyMessage := data.PartyMessage{
		Command: data.Command{
			Action:  data.JOIN,
			Service: data.PARTY,
			User: data.User{
				Username: "TEST PLAYER NAME",
			},
		},
		Party: data.Party{
			Code: "TEST",
		},
	}

	rawMsg, err := json.Marshal(partyMessage)
	if err != nil {
		log.Println("error sending request", err)
	}

	log.Println("Raw message", rawMsg)
	log.Println("Raw message (string)", string(rawMsg[:len(rawMsg)]))

	response, err := c.nc.Request(data.PARTY, rawMsg, 1000*time.Millisecond)
	if err != nil {
		log.Println("error sending request", err)
	}

	subject := data.PARTY + partyMessage.Party.Code
	log.Println("subject", subject)

	// subscribe to all party messages
	c.nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Println("Well, look who decided to join", string(msg.Data[:len(msg.Data)]))
	})

	log.Println("got response", response)

	// do something with response
	// var resMsg data.PartyMessage
	// err = json.Unmarshal(response.Data, &resMsg)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// log.Println("reply", resMsg)

	for {
	}
}

func (c *client) leave() {
	log.Println("Leaving party...")
	// c.sendCh <- &nats.Msg{
	// 	Data: []byte(`{
	// 		"command": "LEAVE",
	// 		"service": "PARTY",
	// 		"party": {
	// 			"code": "TEST_PARTY_CODE"
	// 		},
	// 		"player": {
	// 			"username": "TEST_USERNAME"
	// 		}
	// 	}`),
	// }
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
	return &client{}
}

func main() {
	c := newClient()

	c.connect("nats://10.0.0.111:4222")

	log.Println("Scanning...")
	c.scan()
}
