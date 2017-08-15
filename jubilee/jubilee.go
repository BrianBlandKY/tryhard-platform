package main

import (
	"fmt"
	mess "tryhard-platform/messenger"

	"log"
	"tryhard-platform/model"
)

func main() {
	nickelodeonServer := "nats://localhost:4222"

	h := newJubilee(nickelodeonServer)
	h.Connect()
	h.Listen()
	log.Println("Done.")
}

// Remove references to NATS
type jubilee struct {
	address   string
	messenger mess.Messenger
	sub       mess.Subscription
	subCh     chan mess.Command
}

func (j *jubilee) Connect() {
	err := j.messenger.Connect(j.address, "CARNIVAL_ID")
	if err != nil {
		panic(err)
	}

	j.subscribe()
}

func (j *jubilee) subscribe() {
	ch := make(chan mess.Command)
	key := j.messenger.Key(mess.PARTY, "*")
	recvSub, err := j.messenger.SubscribeChan(key, ch)
	if err != nil {
		fmt.Printf("error creating channel %v", err)
	}

	j.sub = recvSub
	j.subCh = ch
}

func (j *jubilee) processAction(cmd mess.Command, party model.Party) (err error) {
	switch cmd.Action {
	case mess.JOIN:
		j.join(cmd, party)
	case mess.DISBAND:
		j.disband(cmd, party)
	default:
		err = fmt.Errorf("invalid action %v", cmd.Action)
	}
	return
}

// JOIN
func (j *jubilee) join(cmd mess.Command, party model.Party) {
	j.messenger.Reply(cmd)
}

// DISBAND
func (j *jubilee) disband(cmd mess.Command, party model.Party) {
	j.messenger.Reply(cmd)
}

func (j *jubilee) Listen() {
	log.Println("Listening")

	for {
		select {
		case m := <-j.subCh:
			log.Println("received message", m)

			var party model.Party
			err := m.Deserialize(&party)
			if err != nil {
				log.Println("failed to deserialize message", err)
				break
			}

			err = j.processAction(m, party)
			if err != nil {
				log.Printf("process error %v \n\r", err)
				break
			}
		}
	}
}

func newJubilee(address string) (j jubilee) {
	j = jubilee{
		address:   address,
		messenger: mess.NewMessenger(),
	}
	return
}
