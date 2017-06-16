package main

import (
	"bufio"
	"log"
	"os"
	c "tryhard-platform/client"
)

func messageHandler() func(interface{}) {
	return func(msg interface{}) {
		log.Printf("message handler %v \r\n", msg)
	}
}
func main() {
	//_ = getLobbyCode()

	log.Println("Client Tester Started")
	log.Println("Enter command to get started...")

	client := c.NewClient(
		"ws://localhost:8181/connect",
		"http://localhost/",
		messageHandler(),
	)

	// loop input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Bytes()
		log.Printf("command: %s", command)

		switch cmd := string(command[:len(command)]); cmd {
		case "connect":
			log.Println("Connecting client...")
			client.Connect()
		case "close":
			log.Println("Closing client...")
			client.Close()
		case "party":
			log.Println("Generating a party...")
			client.Party()
		case "join":
			log.Println("Joining a party...")
			client.Join()
		case "leave":
			log.Println("Leaving party...")
			client.Leave()
		default:
			client.Write(command)
		}
	}
}
