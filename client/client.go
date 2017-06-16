package client

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/url"
	"reflect"

	"golang.org/x/net/websocket"
)

type Client interface {
	Connect()
	Close()
	Party()
	Join()
	Leave()
	Write([]byte)
}
type message struct {
	Command   string `json:"command"`
	PartyCode string `json:"party_code"`
}
type client struct {
	location  *url.URL
	origin    *url.URL
	ws        *websocket.Conn
	writeChan chan []byte
	doneChan  chan bool
	messageFn func(interface{})
}

func (c *client) Connect() {
	reset := func() { c.doneChan <- false }
	go reset()

	// websocket connection
	ws, err := websocket.DialConfig(&websocket.Config{
		Location: c.location,
		Origin:   c.origin,
		Version:  13,
		Protocol: []string{""},
	})
	if err != nil {
		log.Fatal(err)
	}

	c.ws = ws

	go c.listenRead()
	go c.listenWrite()
	log.Println("Client Connected")
}
func (c *client) Close() {
	c.doneChan <- true
	err := c.ws.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Client Closed")
}
func (c *client) Party() {
	// send API request to build party?
	// what is the downside to just using messages?
}
func (c *client) Join() {
	c.Write([]byte(`{ 
		"command": "CONNECT",
		"player": {
			"id": "1",
			"party_code": "test",
			"username": "test"
		}
	 }`))
}
func (c *client) Leave() {
	c.Write([]byte(`{ 
		"command": "DISCONNECT",
		"player": {
			"id": "1",
			"party_code": "test",
			"username": "test"
		}
	 }`))
}
func (c *client) Write(msg []byte) {
	c.writeChan <- msg
}
func (c *client) listenWrite() {
	for {
		select {
		case msg := <-c.writeChan:
			websocket.Message.Send(c.ws, msg)
		case <-c.doneChan:
			return
		}
	}
}
func (c *client) listenRead() {
	for {
		select {
		case <-c.doneChan:
			return
		default:
			var msg message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF || err == err.(*net.OpError) {
				return
			} else if err == err.(*json.SyntaxError) {
				log.Printf("CLIENT: bad message %v", err)
			} else if err != nil {
				log.Printf("unknown client error type %v", reflect.TypeOf(err).Elem())
				return
			}
			c.processMessage(&msg)
		}
	}
}
func (c *client) processMessage(msg *message) {
	//log.Printf("received: %v", msg)
	if msg.Command == "HEARTBEAT" {
		heartbeatMsg := []byte(`{"command": "HEARTBEAT"}`)
		c.writeChan <- heartbeatMsg
	} else {
		c.messageFn(msg)
	}
}
func NewClient(host string, o string, messageFn func(interface{})) Client {
	location, _ := url.ParseRequestURI(host)
	origin, _ := url.ParseRequestURI(o)
	writeChan := make(chan []byte)
	doneChan := make(chan bool)
	return &client{
		location,
		origin,
		nil,
		writeChan,
		doneChan,
		messageFn,
	}
}
