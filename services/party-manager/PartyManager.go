package main

import (
	"fmt"
	data "tryhard-platform/data"

	"encoding/json"
	"log"

	nats "github.com/nats-io/go-nats"
)

type partyManager struct {
	address string
	nc      *nats.Conn
	ec      *nats.EncodedConn
	testCh  chan []byte
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

func (pm *partyManager) ListenTest() {
	for {
		select {
		case msg := <-pm.testCh:
			log.Println("test chan", msg)
			err := pm.nc.Publish("PARTYTEST", msg)
			if err != nil {
				log.Println("error", err)
			}
		}
	}
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
		var msg data.Message
		pm.decode(rawMsg.Data, &msg)
		log.Println("decoded msg", msg)

		if msg.Command.Action == data.JOIN {
			var joinMsg data.PartyMessage
			pm.decode(rawMsg.Data, &joinMsg)
			log.Println("join msg", joinMsg)

			// create Party
			// data storage?
			// pull party status from source
			// it may not always be in OPENED state.
			joinMsg.Party.Status = data.OPENED

			repMsg, err := json.Marshal(joinMsg)
			if err != nil {
				fmt.Println("error:", err)
			}

			log.Println("rep msg", rawMsg.Reply, repMsg)

			err = pm.nc.Publish(rawMsg.Reply, repMsg)
			if err != nil {
				fmt.Println("error:", err)
			}

			log.Println("sent reply", repMsg)

			joinMsg.Action = data.JOINED
			// broadcastMsg, err := json.Marshal(joinMsg)
			// if err != nil {
			// 	fmt.Println("error:", err)
			// }

			// emit connected player to party
			// this breaks the req/rep for some reason
			subject := data.PARTY + joinMsg.Party.Code
			log.Println("emit message to party", subject)
			pm.testCh <- []byte("TEST")
			//pm.nc.Publish(subject, []byte("TEST BROADCAST"))
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
		testCh:  make(chan []byte),
	}
}
