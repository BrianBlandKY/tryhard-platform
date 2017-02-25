package messenger

import (
	"log"
	"time"

	nats "github.com/nats-io/go-nats"
)

const (
	// JSONEncoder for Json encoded connection
	JSONEncoder string = nats.JSON_ENCODER
	GOBEncoder  string = nats.GOB_ENCODER
)

type natsDialect struct {
	id  string
	enc *nats.EncodedConn
}

func (n *natsDialect) connect(url string, id string, encoding string) error {
	opts := func(opt *nats.Options) error {
		opt.ClosedCB = n.closed
		opt.ReconnectedCB = n.reconnected
		opt.DisconnectedCB = n.disconnected
		opt.MaxReconnect = 5
		opt.AllowReconnect = true
		opt.Name = id
		opt.ReconnectWait = 10 * time.Second
		opt.Timeout = 30 * time.Second
		opt.PingInterval = 5 * time.Second
		opt.MaxPingsOut = 5
		return nil
	}
	n.id = id
	nc, err := nats.Connect(url, opts)
	if err != nil {
		return err
	}

	if len(encoding) == 0 {
		panic("dialect encoder required.")
	}

	c, err := nats.NewEncodedConn(nc, encoding)
	if err != nil {
		return err
	}
	n.enc = c
	return nil
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

// dialect Connection interface
func (n *natsDialect) Publish(subject string, cmd Command) error {
	cmd.Subject = subject
	return n.PublishCommand(cmd)
}

func (n *natsDialect) PublishCommand(cmd Command) error {
	return n.enc.Publish(cmd.Subject, cmd)
}

func (n *natsDialect) PublishRequestCommand(reply string, cmd Command) error {
	cmd.Reply = reply
	return n.enc.Publish(cmd.Reply, cmd)
}

func (n *natsDialect) Subscribe(key string, cmdFn CommandFn) (Subscription, error) {
	sub, err := n.enc.Subscribe(key, cmdFn)
	if err != nil {
		return nil, err
	}
	s := &natsSubscription{
		sub: sub,
	}
	return s, err
}

func (n *natsDialect) RequestCommand(cmd Command, res *Command) {
	n.Request(cmd.Subject, cmd, res)
}

func (n *natsDialect) Request(subject string, cmd Command, res *Command) {
	n.RequestTimeout(subject, cmd, res, DefaultTimeout)
}

func (n *natsDialect) RequestTimeout(subject string, cmd Command, res *Command, timeout time.Duration) {
	n.enc.Request(subject, cmd, res, timeout)
}

func (n *natsDialect) Close() {
	n.enc.Close()
}

func (n *natsDialect) Status() Status {
	switch s := n.enc.Conn.Status(); s {
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

func (n *natsDialect) ID() string {
	return n.id
}

func (n *natsDialect) SetID(id string) {
	n.id = id
}

func (n *natsDialect) IsConnected() bool {
	return n.enc.Conn.IsConnected()
}

func (n *natsDialect) IsClosed() bool {
	return n.enc.Conn.IsClosed()
}

type natsSubscription struct {
	sub *nats.Subscription
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
	return s.sub.Unsubscribe()
}
