package main

import (
	"os"
	"tryhard-platform/src/config"
	svc "tryhard-platform/src/services/situation"
)

func main() {
	configFile := os.Args[1]
	cfg := config.ReadConfig(configFile)

	s := svc.GerSituationService(cfg)
	s.Execute()
}
