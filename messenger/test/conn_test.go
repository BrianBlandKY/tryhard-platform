package test

import (
	"testing"
	d "tryhard-platform/messenger"
)

func TestDefaultConnection(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc := NewDefaultConnection(t)
	defer nc.Close()
}

func TestConnectionStatus(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc := NewDefaultConnection(t)
	defer nc.Close()

	if nc.Status() != d.CONNECTED {
		t.Fatal("Should have status set to CONNECTED")
	}
	if !nc.IsConnected() {
		t.Fatal("Should have status set to CONNECTED")
	}
	nc.Close()
	if nc.Status() != d.CLOSED {
		t.Fatal("Should have status set to CLOSED")
	}
	if !nc.IsClosed() {
		t.Fatal("Should have status set to CLOSED")
	}
}
