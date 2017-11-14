package test

import (
	"fmt"
	"log"
	"testing"
	messenger "try-hard-platform/messenger"
	mock "try-hard-platform/messenger/test"
	realm "try-hard-platform/realm"

	nats "github.com/nats-io/go-nats"
)

func TestConnectionSubscription(t *testing.T) {
	s := mock.RunDefaultServer()
	defer s.Shutdown()

	url := fmt.Sprintf("nats://localhost:%d", nats.DefaultPort)

	r := realm.Realm{}
	r.Connect(url, "test_realm")

	nc := mock.NewDefaultConnection(t)
	log.Println("CLIENT - sending connection message")
	nc.Publish("connection", messenger.Command{
		Data: nc.ID(),
	})

	ch := make(chan bool)

	nc.Subscribe("connected", func(subject, reply string, cmd messenger.Command) {
		log.Println("CLIENT - received connected message")
		nc.Close()
		log.Println("CLIENT - IsConnected", nc.IsConnected())
		log.Println("CLIENT - IsClosed", nc.IsClosed())
		ch <- true
	})

	if e := mock.Wait(ch); e != nil {
		t.Fatalf("No command received\n")
	}
}

// func TestManualDisconnect(t *testing.T) {
// 	s := mock.RunDefaultServer()
// 	defer s.Shutdown()

// 	url := fmt.Sprintf("nats://localhost:%d", nats.DefaultPort)

// 	r := realm.Realm{}
// 	r.Connect(url, "test_realm")

// 	nc := mock.NewDefaultConnection(t)
// 	nc.Publish("connection", messenger.Command{
// 		Data: "test",
// 	})

// 	//heartbeatCh := make(chan bool)

// 	nc.Subscribe("connected", func(subject, reply string, cmd messenger.Command) {
// 		log.Println("CLIENT - received [connected]")
// 	})

// 	nc.Subscribe("heartbeat", func(subject, reply string, cmd messenger.Command) {
// 		log.Println("CLIENT - received [heartbeat]")
// 	})

// 	heartbeatCnt := 0
// 	heartbeatFn := func(counter int) {
// 		cmd := messenger.Command{
// 			Subject: "heartbeat",
// 		}
// 		var res messenger.Command
// 		nc.RequestCommand(cmd, &res)
// 		log.Println("CLIENT - response heartbeat", counter, res)
// 	}
// 	for heartbeatCnt < 3 {
// 		log.Println("send heartbeat", heartbeatCnt)
// 		heartbeatFn(heartbeatCnt)
// 		heartbeatCnt++
// 	}
// 	// heartbeatCh <- true
// 	// if e := mock.WaitTime(heartbeatCh, 40*time.Second); e != nil {
// 	// 	t.Fatalf("No command received\n")
// 	// }
// }

// func TestAutoDisconnect(t *testing.T) {
// 	s := mock.RunDefaultServer()
// 	defer s.Shutdown()

// 	url := fmt.Sprintf("nats://localhost:%d", nats.DefaultPort)

// 	r := realm.Realm{}
// 	r.Connect(url, "test_realm")

// 	nc := mock.NewDefaultConnection(t)
// 	nc.Publish("connection", messenger.Command{
// 		Data: "test",
// 	})

// 	//heartbeatCh := make(chan bool)

// 	nc.Subscribe("connected", func(subject, reply string, cmd messenger.Command) {
// 		log.Println("CLIENT - received [connected]")
// 	})

// 	nc.Subscribe("heartbeat", func(subject, reply string, cmd messenger.Command) {
// 		log.Println("CLIENT - received [heartbeat]")
// 	})

// 	heartbeatCnt := 0
// 	heartbeatFn := func(counter int) {
// 		cmd := messenger.Command{
// 			Subject: "heartbeat",
// 		}
// 		var res messenger.Command
// 		nc.RequestCommand(cmd, &res)
// 		log.Println("CLIENT - response heartbeat", counter, res)
// 	}
// 	for heartbeatCnt < 3 {
// 		log.Println("send heartbeat", heartbeatCnt)
// 		heartbeatFn(heartbeatCnt)
// 		heartbeatCnt++
// 	}
// 	// heartbeatCh <- true
// 	// if e := mock.WaitTime(heartbeatCh, 40*time.Second); e != nil {
// 	// 	t.Fatalf("No command received\n")
// 	// }
// }
