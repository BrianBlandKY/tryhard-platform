package socket

import (
    "log"
    "io"
    "golang.org/x/net/websocket"
    "reflect"
    "encoding/json"
    "net"
    "time"
)
const (
    CLIENT_HEARTBEAT 	= "HEARTBEAT"
)
type client struct { 
    Id        		string			`json:"client_id"`
    ws          	*websocket.Conn		
    server      	*Server	
    party			*party
    writeMessageCh  chan interface{}		
    doneCh      	chan bool			
    heartbeat		heartbeat		
}
func (c *client) listen() {
    go c.monitor()
    go c.listenWrite()
    c.listenRead()
}
func (c *client) listenWrite() {
    for {
		select {
		case msg := <-c.writeMessageCh:
			websocket.JSON.Send(c.ws, msg)
		}
	}
}
func (c *client) listenRead() {
    for {
		select {
		default:
			msg := new(message)
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF || err == err.(*net.OpError) {
				log.Printf("client %v stopped \r\n", c.Id)
				if c.party != nil {
					c.party.drop(c)
				}
				c.server.drop(c)				
				c.doneCh <- true
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
func (c *client) monitor() {
	for {
		select {
		case <- c.heartbeat.ticker.C:
			c.heartbeat.latencyTime = time.Now()
			msg := message{
				client: c,
				Command: CLIENT_HEARTBEAT,
			}
			c.write(msg)
		case <-c.doneCh:
			return				
		}
	}
}
func (c *client) done() {
    c.doneCh <- true
}
func (c *client) write(msg interface{}) {
    c.writeMessageCh <- msg
}
func (c *client) processMessage(msg *message) {
	if msg.Command == CLIENT_HEARTBEAT {
		latencyTime := time.Now().Nanosecond() - c.heartbeat.latencyTime.Nanosecond()
		c.heartbeat.latency = (latencyTime / 1000000) / 2
		log.Printf("Heartbeat %vms", c.heartbeat.latency)
		return
	}
	msg.client = c
	c.server.write(msg)
}
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
				latencyTime: time.Now(),
				latency: 0,
				ticker: time.NewTicker(5 * time.Second),
			}}
}
