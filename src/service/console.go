package service

import (
	"bufio"
	"os"
	cfg "tryhard-platform/src/config"
)

func DefaultServiceScanner(svc Service) {
	// Show some cool output?
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Bytes()
		svc.Printf("command: %s \r\n", command)

		switch cmd := string(command[:len(command)]); cmd {
		case "print config":
			svc.Printf("\r\n %v \r\n", cfg.WriteConfig(svc.Config()))
		case "start":
			_ = svc.Start()
		case "stop":
			svc.Stop()
		default:
			svc.Println("unrecognized command", cmd)
		}
	}
}
