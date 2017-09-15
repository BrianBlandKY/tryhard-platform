package main

import (
	"log"
	"tryhard-platform/config"
)

func main() {
	cfg := config.ReadConfig("config.yaml")

	log.Printf("\r\n %v \r\n", config.WriteConfig(cfg))

	h := newPartyService(cfg)
	h.Listen()
}
