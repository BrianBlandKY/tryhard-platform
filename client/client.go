package client

import (
    "io"
    "log"
    "golang.org/x/net/websocket"
    "net"
    "net/url"
    "reflect"
    "encoding/json"
)
type Client interface {
    Connect()
    Write([]byte)
}
type message struct {
    Command     string  `json:"command"`
    PartyCode   string  `json:"party_code"`
}
type client struct {
    ws          *websocket.Conn
    writeChan   chan []byte
    doneChan    chan bool
    messageFn   func(interface{})
}
func (c *client) Connect() {
    go c.listenRead()
    go c.listenWrite()
}
func (c *client) Write(msg []byte) {
    c.writeChan <- msg
}
func (c *client) listenWrite() {
	for {
		select {
		case msg := <-c.writeChan:
			websocket.Message.Send(c.ws, msg)
		}
	}
}
func (c *client) listenRead() {
	for {
		select {
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
			} else {
				c.processMessage(&msg)
			}
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
func NewClient(host string, o string, messageFn func(interface{})) (Client) {
    location, _ := url.ParseRequestURI(host)                 
    origin, _ := url.ParseRequestURI(o)                     

    ws, err := websocket.DialConfig(&websocket.Config {
        Location: location,
        Origin: origin,
        Version: 13,
        Protocol: []string{""},
    })
    if err != nil {
        log.Fatal(err)
    } 
    writeChan := make(chan []byte)
    doneChan := make(chan bool)
    return &client{
        ws,
        writeChan,
        doneChan,
        messageFn,
    }
}