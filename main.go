package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"

	c "tryhard-platform/config"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	cfg := c.ParseConfig("./app.config")

	if cfg.Env == "development" {
		flag.Parse()
		if *cpuprofile != "" {
			f, err := os.Create(*cpuprofile)
			if err != nil {
				log.Fatal(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}

		go func() {
			log.Println(http.ListenAndServe(":6060", nil))
		}()
	}

	StartServer(cfg)
}
