package server

import (
	"fmt"
	"log"
	"time"
)

type party struct {
	server  *Server
	players map[string]*player
	addCh   chan *message
	dropCh  chan *client
	doneCh  chan bool
	errCh   chan error
}

func (p *party) listen() {
	for {
		select {
		case msg := <-p.addCh:
			log.Printf("party %v added player %v \r\n", p.Code, msg.client.ID)
			p.setPlayer(msg, CONNECTED)
			log.Println("Now", len(p.players), "players connected.")

			playerList := ""

			for _, player := range p.players {
				playerList += fmt.Sprintf("player %v \r\n", player.client.ID)
			}

			log.Println(playerList)

			player := p.players[msg.client.ID]
			msg.client.party = p

			log.Println("sent PARTY_CONNECTED to client", msg.client.ID)
			msg.client.write(message{
				Command: CONNECTED,
				Players: p.list(),
			})

			p.broadcast(message{
				client:    msg.client,
				Command:   JOINED,
				ID:        msg.client.ID,
				PartyCode: p.Code,
				Username:  player.Username,
			})
		case c := <-p.dropCh:
			log.Printf("party %v dropped player %v", p.Code, c.ID)

			player := p.players[c.ID]
			player.Status = DISCONNECTED

			log.Println("Now", len(p.players), "players connected.")

			playerList := ""

			for _, player := range p.players {
				playerList += fmt.Sprintf("player %v \r\n", player.client.ID)
			}

			log.Println(playerList)

			c.write(message{
				client:  c,
				Command: DISCONNECTED,
			})

			p.broadcast(message{
				client:    c,
				ID:        c.ID,
				PartyCode: p.Code,
				Username:  player.Username,
			})
		case err := <-p.errCh:
			log.Printf("Party %v error %v", p.Code, err)
			p.server.err(err)
		case <-p.doneCh:
			log.Printf("Party %v stopped listening", p.Code)
			// send party disconnect to all connected clients
			p.broadcast(message{
				client:  nil,
				Command: DISCONNECTED,
			})
			return
		}
	}
}
func (p *party) broadcast(msg message) {
	for _, player := range p.players {
		player.client.write(msg)
	}
}
func (p *party) reset() {
	p.Status = OPENED
	p.players = make(map[string]*player)
}
func (p *party) write(msg *message) {
	p.process(msg)
}
func (p *party) process(msg *message) {
	switch {
	case msg.Command == CONNECT:
		p.add(msg)
		break
	case msg.Command == DISCONNECT:
		p.drop(msg.client)
		break
	default:
		break
	}
}
func (p *party) add(c *message) {
	p.addCh <- c
}
func (p *party) drop(c *client) {
	p.dropCh <- c
}
func (p *party) err(err error) {
	p.errCh <- err
}
func (p *party) list() []*player {
	list := make([]*player, len(p.players))
	i := 0
	for _, player := range p.players {
		list[i] = player
		i++
	}
	return list
}
func (p *party) setPlayer(msg *message, status string) {
	log.Println("party set player", status, msg.client.ID)
	if x, ok := p.players[msg.client.ID]; ok {
		x.client = msg.client
		x.Status = status
		x.ID = msg.client.ID
		x.Username = msg.Username
	} else {
		p.players[msg.client.ID] = &player{
			msg.client,
			msg.client.ID,
			msg.Username,
			status,
			time.Now(),
		}
	}
}
func newParty(code string, server *Server) *party {
	players := make(map[string]*player)
	addCh := make(chan *message)
	dropCh := make(chan *client)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &party{
		server,
		players,
		addCh,
		dropCh,
		doneCh,
		errCh,
		code,
		OPENED,
		time.Now(),
	}
}
