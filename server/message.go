package socket

import ()

type message struct {
    client          *client    
    Command         string      	`json:"command"`
    Player 			*playerMessage	`json:"player"`
}

type playerMessage struct {
	Id				string			`json:"id"`
	PartyCode 		string			`json:"party_code"`
	Username		string			`json:"username"`
}

type playersMessage struct {
	Command 	string				`json:"command"`
	Players 	[]*player 			`json:"players"`
}