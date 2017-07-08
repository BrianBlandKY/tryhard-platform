package test

import (
	"testing"
	d "tryhard-platform/messenger"
)

func TestChan(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc1 := NewDefaultConnection(t)
	defer nc1.Close()

	nc2 := NewDefaultConnection(t)
	defer nc2.Close()

	nc1Chan := make(chan d.Command)
	nc1.SubscribeChan("TEST", nc1Chan)

	nc1.Publish(d.Command{
		Service: "TEST",
	})

	msg := <-nc1Chan
	t.Log("received message", msg)

	if msg.Service != "TEST" {
		t.Fatal("Invalid command received")
	}
}
