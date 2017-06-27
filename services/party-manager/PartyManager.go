package main

import (
	"fmt"
	data "tryhard-platform/data"

	"encoding/json"
	"log"

	"time"

	nats "github.com/nats-io/go-nats"
)

type partyManager struct {
	address string
	nc      *nats.Conn
	ec      *nats.EncodedConn
}

func (pm *partyManager) Connect() {
	nc, err := nats.Connect(pm.address)
	if err != nil {
		panic(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	pm.nc = nc
	pm.ec = ec
	log.Println("Connected")
}

func (pm *partyManager) Listen() {
	defer func() {
		log.Println("Stopped Listening")
	}()

	pm.nc.Subscribe(data.PARTY, func(rawMsg *nats.Msg) {
		// process message
		log.Println("raw msg", rawMsg.Data)
		log.Println("Raw message (string)", string(rawMsg.Data[:len(rawMsg.Data)]))

		// decoded message
		var msg data.MessageAdapter
		pm.decode(rawMsg.Data, &msg)
		log.Println("decoded msg", msg)

		if msg.Message.Command == data.JOIN {
			var joinMsg data.PartyMessage
			pm.decode(rawMsg.Data, &joinMsg)
			log.Println("join msg", joinMsg)

			// create Party
			// data storage?

			joinMsg.Party.DateCreated = time.Now()
			// pull party status from source
			// it may not always be in OPENED state.
			joinMsg.Party.Status = data.OPENED
			joinMsg.Player.DateCreated = time.Now()
			joinMsg.Player.Status = data.CONNECTED

			pm.ec.Publish(rawMsg.Reply, joinMsg)
		}
	})

	log.Println("Listening")
	for {
		// what is the best way to idle?
	}
}

func (pm *partyManager) decode(msg []byte, output interface{}) {
	err := json.Unmarshal(msg, output)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func NewPartyManager(address string) partyManager {
	return partyManager{
		address: address,
	}
}
