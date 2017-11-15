package messenger

import (
	"log"
	"time"

	"strings"

	nats "github.com/nats-io/go-nats"
)

const (
	DISCONNECTED = Status(iota)
	CONNECTED
	CLOSED
	RECONNECTING
	CONNECTING
)

type natsDialect struct {
	id     string
	doneCh chan bool
	nc     *nats.Conn
}

func messageFn(fn CommandFn) func(*nats.Msg) {
	return func(m *nats.Msg) {
		cmd := Command{
			Service: m.Subject,
			Reply:   m.Reply,
			Data:    m.Data,
		}
		fn(cmd)
	}
}

func (n *natsDialect) reconnected(conn *nats.Conn) {
	log.Println("nats reconnected")
}

func (n *natsDialect) closed(conn *nats.Conn) {
	log.Println("nats closed")
}

func (n *natsDialect) disconnected(conn *nats.Conn) {
	log.Println("nats disconnected")
}

func (n *natsDialect) Connect(url string) error {
	opts := func(opt *nats.Options) error {
		opt.ClosedCB = n.closed
		opt.ReconnectedCB = n.reconnected
		opt.DisconnectedCB = n.disconnected
		opt.MaxReconnect = -1 // retry forever
		opt.AllowReconnect = true
		opt.ReconnectWait = 10 * time.Second
		opt.Timeout = 30 * time.Second
		opt.PingInterval = 5 * time.Second
		opt.MaxPingsOut = 5
		return nil
	}
	nc, err := nats.Connect(url, opts)
	if err != nil {
		return err
	}
	n.nc = nc

	log.Println("nats connected")
	return nil
}

func (n *natsDialect) Close() {
	n.nc.Close()
}

func (n *natsDialect) Key(params ...string) (key string) {
	return strings.Join(params, ".")
}

func (n *natsDialect) CommandKey(cmd Command, params ...string) (key string) {
	keys := []string{cmd.Service, cmd.Action}
	keys = append(keys, params...)
	return strings.Join(keys, ".")
}

func (n *natsDialect) Publish(cmd Command) error {
	return n.nc.Publish(cmd.Service, cmd.Data)
}

func (n *natsDialect) Request(cmd Command, res *Command) error {
	return n.RequestTimeout(cmd, res, DefaultTimeout)
}

func (n *natsDialect) RequestTimeout(cmd Command, res *Command, timeout time.Duration) (err error) {
	response, err := n.nc.Request(n.CommandKey(cmd), cmd.Data, timeout)
	if err != nil {
		return err
	}

	//log.Printf("messenger request response %v  cmd: %v", response, cmd)
	res.Reply = response.Reply
	res.Service = cmd.Service
	res.Data = response.Data

	//log.Println("messenger reply data", res)
	return
}

func (n *natsDialect) Reply(cmd Command) error {
	return n.nc.Publish(cmd.Reply, cmd.Data)
}

func (n *natsDialect) Subscribe(key string, cmdFn CommandFn) (Subscription, error) {
	sub, err := n.nc.Subscribe(key, messageFn(cmdFn))
	if err != nil {
		return nil, err
	}
	s := &natsSubscription{
		sub: sub,
	}
	return s, err
}

func (n *natsDialect) SubscribeChan(key string, cmdCh chan Command) (Subscription, error) {
	msgCh := make(chan *nats.Msg)
	sub, err := n.nc.ChanSubscribe(key, msgCh)
	if err != nil {
		return nil, err
	}
	s := &natsSubscription{
		sub:    sub,
		msgch:  msgCh,
		cmdCh:  cmdCh,
		stopCh: make(chan bool),
	}
	go s.relayChan()
	return s, err
}

func (n *natsDialect) Status() Status {
	if n.nc == nil {
		return DISCONNECTED
	}
	switch s := n.nc.Status(); s {
	case nats.DISCONNECTED:
		return DISCONNECTED
	case nats.CONNECTED:
		return CONNECTED
	case nats.CLOSED:
		return CLOSED
	case nats.RECONNECTING:
		return RECONNECTING
	case nats.CONNECTING:
		return CONNECTING
	default:
		return DISCONNECTED
	}
}

func (n *natsDialect) IsConnected() bool {
	return n.nc.IsConnected()
}

func (n *natsDialect) IsClosed() bool {
	return n.nc.IsClosed()
}

type natsSubscription struct {
	sub    *nats.Subscription
	msgch  chan *nats.Msg
	cmdCh  chan Command
	stopCh chan bool
}

// Converts messages to commands
func (s *natsSubscription) relayChan() {
	for {
		select {
		case msg := <-s.msgch:
			var cmd Command
			cmds := strings.Split(msg.Subject, ".")
			if len(cmds) > 1 {
				cmd = Command{
					Service: cmds[0],
					Action:  cmds[1],
					Data:    msg.Data,
					Reply:   msg.Reply,
				}
			} else {
				cmd = Command{
					Service: cmds[0],
					Action:  "",
					Data:    msg.Data,
					Reply:   msg.Reply,
				}
			}

			s.cmdCh <- cmd
		case _ = <-s.stopCh:
			return
		}
	}
}

func (s *natsSubscription) Subject() string {
	return s.sub.Subject
}

func (s *natsSubscription) Queue() string {
	return s.sub.Queue
}

func (s *natsSubscription) IsValid() bool {
	return s.sub.IsValid()
}

func (s *natsSubscription) AutoUnsubscribe(max int) error {
	return s.sub.AutoUnsubscribe(max)
}

func (s *natsSubscription) Unsubscribe() error {
	if s.stopCh != nil {
		s.stopCh <- true
	}
	close(s.cmdCh)
	return s.sub.Unsubscribe()
}