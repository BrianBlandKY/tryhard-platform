package test

import (
	"encoding/gob"
	"log"
	"testing"
	d "try-hard-platform/messenger"
)

func TestRequestReply(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc1 := NewDefaultConnection(t)
	defer nc1.Close()

	nc2 := NewDefaultConnection(t)
	defer nc2.Close()

	ch := make(chan bool)

	type RndObject struct {
		Random string
	}

	gob.Register(RndObject{})

	_, err := nc1.Subscribe("test", func(subj, reply string, cmd d.Command) {
		log.Printf("command received on nc1 %v", cmd)
		cmd.Data = &RndObject{
			"random object",
		}
		log.Printf("sending reply to nc2 %v", cmd)
		nc1.PublishRequestCommand(reply, cmd)
		ch <- true
	})
	if err != nil {
		t.Fatal(err)
	}

	cmd := d.Command{
		Subject: "test",
	}
	var res d.Command
	nc2.RequestCommand(cmd, &res)

	log.Printf("response: %v", res)

	if e := Wait(ch); e != nil {
		t.Fatalf("No command received\n")
	}
}
