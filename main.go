package main

import (
    _ "net/http/pprof"
    "runtime/pprof"
    "net/http"
    "flag"
    "os"
    "log"
    
    c "dimples-api/config"
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