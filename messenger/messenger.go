package messenger

import (
	"time"
)

const (
	// DefaultTimeout limit set to 10 seconds
	DefaultTimeout time.Duration = time.Second * 10
)

// Status for connection
type Status int

// Messenger Statuses
const (
	DISCONNECTED = Status(iota)
	CONNECTED
	CLOSED
	RECONNECTING
	CONNECTING
	HEARTBEAT
)

// Messenger interface
type Messenger interface {
	Connect(url, id string) error
	Publish(cmd Command) error
	Request(cmd Command, response *Command) error
	RequestTimeout(cmd Command, response *Command, timeout time.Duration) error
	Reply(cmd Command) error
	Subscribe(service string, cmdFn CommandFn) (Subscription, error)
	SubscribeChan(key string, ch chan Command) (Subscription, error)
	Close()
	ID() string
	Status() Status
	IsConnected() bool
	IsClosed() bool
}

// CommandFn Command Handler
type CommandFn func(cmd Command)

func NewMessenger() Messenger {
	return &natsDialect{}
}
