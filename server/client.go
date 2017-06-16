package server

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"reflect"
	"time"
	d "tryhard-platform/data"

	"golang.org/x/net/websocket"
)

/*
	Manages the client connection.
	- Client Properties
	- WebSocket Connection
	- Server Instance
	- Message Handling
	- Heartbeating
*/
type clientConn struct {
	ws             *websocket.Conn
	client         *d.Client
	server         *Server
	party          *party
	writeMessageCh chan interface{}
	closeCh        chan bool
	heartbeat      heartbeat
}

/*
	Primary listener
	- Heartbeat Monitor
	- Listens for messages from the server and sends to the client
	- Listens for messages from the client and sends to the server
*/
func (c *client) listen() {
	go c.monitor()
	go c.listenWrite()
	c.listenRead()
}

/*
	Listens for messages on the server
	and sends them to the client.
*/
func (c *client) listenWrite() {
	for {
		select {
		case msg := <-c.writeMessageCh:
			websocket.JSON.Send(c.ws, msg)
		}
	}
}

/*
	Listens for message from the client
	and sends them to the server.

	Also manages connection status by handling
	operation errors. When a disconnection occurs
	we disconnect from the party and drop the client
	from the server.
*/
func (c *client) listenRead() {
	for {
		select {
		default:
			msg := new(message)
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF || err == err.(*net.OpError) {
				log.Printf("client %v stopped \r\n", c.ID)
				if c.party != nil {
					c.party.drop(c)
				}
				c.server.drop(c)
				c.closeCh <- true
			} else if err == err.(*json.SyntaxError) || err == err.(*json.UnmarshalTypeError) {
				log.Printf("CLIENT: bad message %v", err)
			} else if err != nil {
				log.Printf("unknown client error type %v \r\n", reflect.TypeOf(err).Elem())
				c.server.err(err)
			} else {
				msg.client = c
				c.processMessage(msg)
			}
		}
	}
}

/*
	Monitors the connection by
	sending heartbeat messages
	to the client.
*/
func (c *client) monitor() {
	for {
		select {
		case <-c.heartbeat.interval.C:
			c.heartbeat.latencyTime = time.Now()
			msg := message{
				client:  c,
				Command: HEARTBEAT,
			}
			c.write(msg)
		case <-c.closeCh:
			return
		}
	}
}

/*
	Close the WebSocket connection
	and stop all running go functions.
*/
func (c *client) close() {
	c.closeCh <- true
}

/*
	Writes a message to the channel
	and sent to the server.
*/
func (c *client) write(msg interface{}) {
	c.writeMessageCh <- msg
}

/*
	Processes messages

	Updates HEARTBEAT message
	or
	Sends message to server for other clients.
*/
func (c *client) processMessage(msg *message) {
	if msg.Command == HEARTBEAT {
		latencyTime := time.Now().Nanosecond() - c.heartbeat.latencyTime.Nanosecond()
		c.heartbeat.latencySeconds = (latencyTime / 1000000) / 2
		log.Printf("Heartbeat %vms", c.heartbeat.latencySeconds)
		return
	}
	msg.client = c
	c.server.write(msg)
}

/*
	Spins up a new client
*/
func newClient(id string, ws *websocket.Conn, server *Server) *client {
	if ws == nil {
		panic("ws cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	messageCh := make(chan interface{})
	doneCh := make(chan bool)

	return &client{
		id,
		ws,
		server,
		nil,
		messageCh,
		doneCh,
		heartbeat{
			latencyTime:    time.Now(),
			latencySeconds: 0,
			interval:       time.NewTicker(5 * time.Second),
		}}
}
