package test

import (
	"fmt"
	"time"

	"github.com/nats-io/gnatsd/server"
	gnatsd "github.com/nats-io/gnatsd/test"
	"github.com/nats-io/go-nats"
)

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

// NewConnection forms connection on a given port.
func NewConnection(port int) *nats.Conn {
	url := fmt.Sprintf("nats://localhost:%d", port)
	nc, err := nats.Connect(url)
	if err != nil {
		fmt.Printf("failed to connect %v", err)
		return nil
	}
	return nc
}

func NewDefaultConnection() *nats.Conn {
	return NewConnection(nats.DefaultPort)
}

func main() {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc := NewDefaultConnection()
	defer nc.Close()

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		fmt.Printf("Failed to create encoded conn %v", err)
	}
	defer c.Close()

	ch := make(chan bool)

	// Replying
	c.Subscribe("help", func(subj, reply string, msg string) {
		fmt.Printf("subsciber received: %v, %v, %v\n", subj, reply, msg)
		c.Publish(reply, "I can help!")
		ch <- true
	})

	// Requests
	var response string
	err = c.Request("help", "help me", &response, 10*time.Millisecond)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	}
	for {
		select {
		case <-ch:
			fmt.Printf("response %v\n", response)
			fmt.Printf("done")
			return
		}
	}
}
