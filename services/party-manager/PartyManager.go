package main

import (
	"fmt"
	data "tryhard-platform/data"

	"encoding/json"
	"log"

	nats "github.com/nats-io/go-nats"
)

type partyManager struct {
	address      string
	nc           *nats.Conn
	partySubChan chan *nats.Msg
}

func (pm *partyManager) Connect() {
	nc, err := nats.Connect(pm.address)
	if err != nil {
		panic(err)
	}

	pm.nc = nc
	log.Println("Connected")
}

func (pm *partyManager) bind() {
	ch := make(chan *nats.Msg, 64)
	_, err := pm.nc.ChanSubscribe(data.PARTY, ch)
	if err != nil {
		fmt.Printf("error creating channel %v", err)
	}
	pm.partySubChan = ch
}

func (pm *partyManager) Listen() {
	log.Println("Listening")

	for {
		select {
		case m := <-pm.partySubChan:
			// process message
			log.Println("raw msg", m.Data)
			log.Println("Raw message (string)", string(m.Data[:len(m.Data)]))

			// decoded message
			var msg data.Message
			pm.decode(m.Data, &msg)
			log.Println("decoded msg", msg)

			if msg.Command.Action == data.JOIN {
				var joinMsg data.PartyMessage
				pm.decode(m.Data, &joinMsg)
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

				log.Println("rep msg", m.Reply, repMsg)

				err = pm.nc.Publish(m.Reply, repMsg)
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
				//pm.testCh <- []byte("TEST")
				pm.nc.Publish(subject, []byte("TEST BROADCAST TO PARTY"))
			}
		}
	}
}

func (pm *partyManager) decode(msg []byte, output interface{}) {
	err := json.Unmarshal(msg, output)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func NewPartyManager(address string) partyManager {
	pm := partyManager{
		address: address,
	}
	pm.bind()
	return pm
}
