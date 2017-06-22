package main

import (
	"log"

	nats "github.com/nats-io/go-nats"
)

func basicUsage() {
	nickelodeonServer := "nats://10.0.0.111:4222"
	conn, err := nats.Connect(nickelodeonServer)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected")
	conn.Close()
	log.Println("Disconnected")
}

func main() {
	basicUsage()
}
