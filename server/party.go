package socket

import (
	"time"
	"log"
	"fmt"
)

const (
	PARTY_STATUS_OPENED     = "OPENED"
	PARTY_STATUS_SUSPENDED  = "SUSPENDED"
	PARTY_STATUS_STALE      = "STALE"
	PARTY_STATUS_INGAME     = "IN-GAME"
	PARTY_STATUS_CLOSED     = "CLOSED"
	PARTY_NOT_AVAILABLE		= "PARTY_NOT_AVAILABLE"
	PARTY_CONNECT  	      	= "PARTY_CONNECT"
	PARTY_CONNECTED     	= "PARTY_CONNECTED"
	PARTY_DISCONNECT  		= "PARTY_DISCONNECT"
	PARTY_DISCONNECTED  	= "PARTY_DISCONNECTED"
	PARTY_PLAYER_JOIN   	= "PARTY_PLAYER_JOIN"
	PARTY_PLAYER_LEFT 		= "PARTY_PLAYER_LEFT"
)

type party struct {
	server          *Server              
	players         map[string]*player   
	addCh           chan *message         
	dropCh          chan *client         
	doneCh          chan bool           
	errCh           chan error        

	// exports
	Code       	    string      	`json:"code"`
	Status          string         	`json:"status"`
	DateCreated     time.Time      	`json:"date_created"`
}
func (p *party) listen() {
	for {
		select {
		case msg := <- p.addCh:
			log.Printf("party %v added player %v \r\n", p.Code, msg.client.Id)
			p.setPlayer(msg, PARTY_CONNECTED)
			log.Println("Now", len(p.players), "players connected.")

			playerList := ""

			for _, player := range p.players {
				playerList += fmt.Sprintf("player %v \r\n", player.client.Id)			
			}

			log.Println(playerList)
			
			player := p.players[msg.client.Id]
			msg.client.party = p

			log.Println("sent PARTY_CONNECTED to client", msg.client.Id)
			msg.client.write(playersMessage{
				Command: PARTY_CONNECTED,
				Players: p.list(),
			})
			
			p.broadcast(message{
				client: msg.client,
				Command: PARTY_PLAYER_JOIN,
				Player: &playerMessage {
					Id: msg.client.Id,
					PartyCode: p.Code,
					Username: player.Username,
				},
			})
		case c := <-p.dropCh:
			log.Printf("party %v dropped player %v", p.Code, c.Id)
			
			player := p.players[c.Id]
			player.Status = PARTY_DISCONNECTED

			log.Println("Now", len(p.players), "players connected.")

			playerList := ""

			for _, player := range p.players {
				playerList += fmt.Sprintf("player %v \r\n", player.client.Id)			
			}

			log.Println(playerList)			

			c.write(message{
				client: c,
				Command: PARTY_DISCONNECTED,
			})
			
			p.broadcast(message{
				client: c,
				Command: PARTY_PLAYER_LEFT,
				Player: &playerMessage {
					Id: c.Id,
					PartyCode: p.Code,
					Username: player.Username,
				},
			})
		case err := <- p.errCh:
			log.Printf("Party %v error %v", p.Code, err)
			p.server.err(err)
		case <-p.doneCh:
			log.Printf("Party %v stopped listening", p.Code)
			// send party disconnect to all connected clients
			p.broadcast(message{
				client: nil,
				Command: PARTY_DISCONNECTED,
			})
			return
		}
	}    
}
func (p *party) broadcast(msg message){
	for _, player := range p.players {
		player.client.write(msg)			
	}
}
func (p *party) reset() {
	p.Status = PARTY_STATUS_OPENED
	p.players = make(map[string]*player)
}
func (p *party) write(msg *message) {
	p.process(msg) 
}
func (p *party) process(msg *message) {
	switch {
		case msg.Command == PARTY_CONNECT:
			p.add(msg)
			break
		case msg.Command == PARTY_DISCONNECT:
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
func (p *party) list() ([]*player) {
	list := make([]*player, len(p.players))
	i := 0
	for _, player := range p.players {
		list[i] = player
		i += 1
	}
	return list
}
func (p *party) setPlayer(msg *message, status string) {
	log.Println("party set player", status, msg.client.Id)
	if x, ok := p.players[msg.client.Id]; ok {
		x.client = msg.client 
		x.Status = status
		x.Id = msg.client.Id
		x.Username = msg.Player.Username
	} else {
		p.players[msg.client.Id] = &player{
			msg.client,
			msg.client.Id,
			msg.Player.Username,
			status,
			time.Now(),
		}
	}
}
func newParty(code string, server *Server) *party {
	players     := make(map[string]*player)
	addCh       := make(chan *message)
	dropCh      := make(chan *client)
	doneCh      := make(chan bool)
	errCh       := make(chan error)

	return &party{
		server,
		players,
		addCh,
		dropCh,
		doneCh,
		errCh,
		code,
		PARTY_STATUS_OPENED,
		time.Now(),
	}
}
