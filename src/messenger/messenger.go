package messenger

import (
	"time"
)

const (
	// DefaultTimeout limit set to 10 seconds
	DefaultTimeout time.Duration = time.Second * 10
)

// CommandFn Command Handler
type CommandFn func(cmd Command)

// Status for connection
type Status int

// Messenger interface
type Messenger interface {
	CommandKey(cmd Command, params ...string) string
	Key(params ...string) string
	Connect(url string) error
	Close()
	Publish(cmd Command) error
	Request(cmd Command, response *Command) error
	RequestTimeout(cmd Command, response *Command, timeout time.Duration) error
	Reply(cmd Command) error
	Subscribe(service string, cmdFn CommandFn) (Subscription, error)
	SubscribeChan(key string, ch chan Command) (Subscription, error)
	Status() Status
	IsConnected() bool
	IsClosed() bool
}

func DefaultMessenger() Messenger {
	return &natsDialect{}
}
