package server

import (
	"fmt"
	"math/rand"
	"time"
)

type partySource struct {
	parties map[string]*party
}

func (ps *partySource) get(code string) *party {
	return ps.parties[code]
}
func (ps *partySource) host(server *Server, code string) *party {
	return ps.getParty(server, code)
}
func (ps *partySource) getParty(server *Server, code string) *party {
	for {
		if len(code) == 0 {
			code = ps.generateCode()
		}

		if party, ok := ps.parties[code]; ok {
			//if party.Status == PARTY_STATUS_STALE {
			//party.reset()
			return party
			//}
		}
		party := newParty(code, server)
		ps.parties[code] = party
		return party
	}
}
func (ps *partySource) generateCode() string {
	return fmt.Sprintf("%s%s%s%s",
		ps.getRandomChar(),
		ps.getRandomChar(),
		ps.getRandomChar(),
		ps.getRandomChar())
}
func (ps *partySource) getRandomChar() string {
	// 65 - 90 A-Z
	startIndex := 65
	randomRange := 25
	randomValue := rand.Intn(randomRange)
	return string(startIndex + randomValue)
}
func (ps *partySource) getRandomDigit() string {
	// 48 - 57 0-9
	startIndex := 48
	randomRange := 10
	randomValue := rand.Intn(randomRange)
	return string(startIndex + randomValue)
}
func newPartySource() *partySource {
	rand.Seed(time.Now().UnixNano())
	parties := make(map[string]*party)
	return &partySource{
		parties,
	}
}
