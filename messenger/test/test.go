package test

import (
	"errors"
	"fmt"
	"time"
	d "tryhard-platform/messenger"

	"github.com/nats-io/gnatsd/server"
	gnatsd "github.com/nats-io/gnatsd/test"
	nats "github.com/nats-io/go-nats"
)

///
/// Borrowed functionality from
/// https://github.com/nats-io/go-nats/blob/master/test/test.go
///

// So that we can pass tests and benchmarks...
type tLogger interface {
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// TestLogger
type TestLogger tLogger

// Dumb wait program to sync on callbacks, etc... Will timeout
func Wait(ch chan bool) error {
	return WaitTime(ch, 5*time.Second)
}

// Wait for a chan with a timeout.
func WaitTime(ch chan bool, timeout time.Duration) error {
	select {
	case <-ch:
		return nil
	case <-time.After(timeout):
	}
	return errors.New("timeout")
}

// NewDefaultConnection
func NewDefaultConnection(t tLogger) d.Messenger {
	return NewConnection(t, nats.DefaultPort)
}

// NewConnection forms connection on a given port.
func NewConnection(t tLogger, port int) d.Messenger {
	url := fmt.Sprintf("nats://localhost:%d", port)
	messenger := d.NewMessenger()
	err := messenger.Connect(url, "test_client")
	if err != nil {
		t.Fatalf("Failed to create default connection: %v\n", err)
		return nil
	}
	return messenger
}

// RunDefaultServer will run a server on the default port.
func RunDefaultServer() *server.Server {
	return RunServerOnPort(nats.DefaultPort)
}

// RunServerOnPort will run a server on the given port.
func RunServerOnPort(port int) *server.Server {
	opts := gnatsd.DefaultTestOptions
	opts.Port = port
	return RunServerWithOptions(opts)
}

// RunServerWithOptions will run a server with the given options.
func RunServerWithOptions(opts server.Options) *server.Server {
	return gnatsd.RunServer(&opts)
}
