package messenger

import (
	"time"
)

// Status for connection
type Status int

const (
	DISCONNECTED = Status(iota)
	CONNECTED
	CLOSED
	RECONNECTING
	CONNECTING
	HEARTBEAT
)

// Connection interface
type Connection interface {
	Publish(key string, cmd Command) error
	PublishCommand(cmd Command) error
	PublishRequestCommand(reply string, cmd Command) error
	Subscribe(key string, cmdFn CommandFn) (Subscription, error)
	Request(key string, cmd Command, response *Command)
	RequestCommand(cmd Command, res *Command)
	RequestTimeout(key string, cmd Command, response *Command, timeout time.Duration)
	Close()
	ID() string
	SetID(string)
	Status() Status
	IsConnected() bool
	IsClosed() bool
}
