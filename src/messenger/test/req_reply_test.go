package test

import (
	"log"
	"testing"
	d "tryhard-platform/messenger"
)

// func TestRequestReply(t *testing.T) {
// 	s := RunDefaultServer()
// 	defer s.Shutdown()

// 	nc := NewDefaultConnection(t)
// 	defer nc.Close()

// 	nc2 := NewDefaultConnection(t)
// 	defer nc2.Close()
// 	ch := make(chan bool)

// 	go func() {
// 		t.Logf("subscribing to %v", "test")
// 		_, err := nc.Subscribe("test", func(cmd d.Command) {
// 			t.Logf("command received on nc1 %v", cmd)
// 			t.Logf("sending reply to nc2 %v", cmd)
// 			nc.Reply(cmd)
// 			ch <- true
// 		})
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 	}()

// 	cmd := d.Command{
// 		Service: "test",
// 	}
// 	valObj := TestObj{
// 		Value: "Object Serialize",
// 	}

// 	t.Logf("serializing %v", valObj)
// 	// err = cmd.Serialize(valObj)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }

// 	time.Sleep(2 * time.Second)

// 	var res d.Command
// 	t.Logf("requesting %v", cmd)
// 	err := nc2.Request(cmd, &res)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// var resultObj TestObj
// 	// err = res.Deserialize(&resultObj)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }

// 	log.Printf("response: %v", res)

// 	// if valObj.Value != resultObj.Value {
// 	// 	t.Fatal("incomplete transmission", resultObj)
// 	// }

// 	if e := Wait(ch); e != nil {
// 		t.Fatalf("No command received\n")
// 	}
// }
func TestRequestReply(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc1 := NewDefaultConnection(t)
	defer nc1.Close()

	nc2 := NewDefaultConnection(t)
	defer nc2.Close()

	ch := make(chan bool)

	_, err := nc1.Subscribe("test", func(cmd d.Command) {
		log.Printf("command received on nc1 %v", cmd)
		cmd.Data = []byte("RANDOM DATA")
		log.Printf("sending reply to nc2 %v", cmd)
		nc1.Reply(cmd)
		ch <- true
	})
	if err != nil {
		t.Fatal(err)
	}

	cmd := d.Command{
		Service: "test",
	}
	var res d.Command
	nc2.Request(cmd, &res)

	log.Printf("response: %v", res)

	if e := Wait(ch); e != nil {
		t.Fatalf("No command received\n")
	}
}
