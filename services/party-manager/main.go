package main

import (
	"log"
)

//func basicUsage() {

// conn, err := nats.Connect(nickelodeonServer)
// if err != nil {
// 	log.Panic(err)
// }

// Party CREATE
//conn.Subscribe()

// Party JOIN

// Party DROP/LEAVE/DISCONNECT

// log.Println("Connected")
// conn.Close()
// log.Println("Disconnected")
//}

func main() {
	nickelodeonServer := "nats://10.0.0.111:4222"

	pm := NewPartyManager(nickelodeonServer)
	pm.Connect()
	go pm.ListenTest()
	pm.Listen()
	log.Println("Done.")
}
