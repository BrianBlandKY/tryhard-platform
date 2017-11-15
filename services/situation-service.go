package main

import (
	"os"
	"tryhard-platform/src/config"
	"tryhard-platform/src/services/situation"
)

func main() {
	configFile := os.Args[1]
	cfg := config.ReadConfig(configFile)

	s := situation.GetSituationService(cfg)
	s.Execute()
}
