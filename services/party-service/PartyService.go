package main

import (
	"fmt"
	"tryhard-platform/config"
	mess "tryhard-platform/messenger"
	"tryhard-platform/services/common"

	"log"
	"tryhard-platform/model"
)

// Remove references to NATS
type partyService struct {
	common.BaseService
}

func (s *partyService) Listen() {
	log.Println("Connecting...")
	s.Connect()

	log.Println("Listening...")

	for {
		select {
		case m := <-s.SubCh:
			log.Println("received message", m)

			var party model.Party
			err := m.Deserialize(&party)
			if err != nil {
				log.Println("failed to deserialize message", err)
				break
			}

			err = s.processAction(m, party)
			if err != nil {
				log.Printf("process error %v \n\r", err)
				break
			}
		}
	}
}

func (s *partyService) processAction(cmd mess.Command, party model.Party) (err error) {
	switch cmd.Action {
	case mess.JOIN:
		s.join(cmd, party)
	case mess.DISBAND:
		s.disband(cmd, party)
	default:
		err = fmt.Errorf("invalid action %v", cmd.Action)
	}
	return
}

// JOIN
func (s *partyService) join(cmd mess.Command, party model.Party) {

	// TODO Data Storage
	// TODO Broadcast connected player to party

	s.Messenger.Reply(cmd)
}

// DISBAND
func (s *partyService) disband(cmd mess.Command, party model.Party) {

	// TODO Data Storage
	// TODO Broadcast connected player to party

	s.Messenger.Reply(cmd)
}

func newPartyService(cfg config.Config) (ps partyService) {
	ps = partyService{
		common.BaseService{
			ID:        cfg.Service.ID,
			Name:      cfg.Service.Name,
			Address:   cfg.Platform.Address,
			Messenger: mess.NewMessenger(),
		},
	}
	return
}
