package test

import (
	"log"
	"testing"
	d "try-hard-platform/messenger"
)

func TestPubSub(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc := NewDefaultConnection(t)
	defer nc.Close()

	ch := make(chan bool)
	_, err := nc.Subscribe("test", func(subj, reply string, cmd d.Command) {
		log.Printf("command received %v", cmd)
		ch <- true
	})

	if err != nil {
		t.Fatal(err)
	}

	cmd := d.Command{
		Subject: "test",
	}

	err = nc.PublishCommand(cmd)
	if err != nil {
		t.Fatal(err)
	}

	if e := Wait(ch); e != nil {
		t.Fatalf("No command received\n")
	}
}
