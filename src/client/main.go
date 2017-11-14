package main

import (
	"log"
	config "tryhard-platform/src/config"
)

func main() {
	cfg := config.ReadConfig("config.yaml")

	log.Printf("\r\n %v \r\n", config.WriteConfig(cfg))
	c := newClient(cfg)

	log.Println("Scanning...")
	c.scan()
}

// log.Println("Raw message", rawMsg)
// log.Println("Raw message (string)", string(rawMsg[:len(rawMsg)]))
