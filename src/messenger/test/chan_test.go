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

	cmd := d.Command{
		Service: "TEST",
	}

	valObj := TestObj{
		Value: "Object Serialize",
	}

	cmd.Serialize(valObj)

	nc1.Publish(cmd)
	msg := <-nc1Chan

	t.Log("received message", msg)

	var testResult TestObj
	msg.Deserialize(&testResult)

	if testResult.Value != valObj.Value {
		t.Fatal("Invalid incomplete transmission")
	}

	if msg.Service != "TEST" {
		t.Fatal("Invalid command received")
	}
}
