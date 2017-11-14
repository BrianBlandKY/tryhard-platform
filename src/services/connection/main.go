package main

import (
	"log"
	"os"
	"tryhard-platform/src/config"
)

func main() {
	configFile := os.Args[1]

	cfg := config.ReadConfig(configFile)

	log.Printf("\r\n %v \r\n", config.WriteConfig(cfg))

	h := newConnectionService(cfg)
	h.Run()
}
