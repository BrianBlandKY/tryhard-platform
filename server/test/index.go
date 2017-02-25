package main

import (
    "log"
    std "github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo"
    "dimples-api/server"
)

func main() {
	defer func() {
		crash := recover()
		if crash != nil {
			log.Printf("Application Crash. %s", crash)
		}
	}()
    
    addr := "localhost:8080"
    
    e := echo.New()
	e.SetDebug(true)
    
    server := socket.NewServer("http://localhost/")
	go server.Listen(e)

    e.Run(std.New(addr))
}

// Client
// package main

// import (
//     "io"
//     "log"
//     "golang.org/x/net/websocket"
//     "bufio"
//     "net"
//     "reflect"
//     "encoding/json"
//     "os"
// )

// func listenWrite(ws *websocket.Conn, writeCh chan string) {
// 	for {
// 		select {
// 		case msg := <-writeCh:
// 			websocket.JSON.Send(ws, msg)
// 		}
// 	}
// }
// func listenRead(ws *websocket.Conn) {
// 	for {
// 		select {
// 		default:
// 			var msg string
// 			err := websocket.JSON.Receive(ws, &msg)
// 			if err == io.EOF || err == err.(*net.OpError) {
//                 return
// 			} else if err == err.(*json.SyntaxError) {
//                 log.Printf("CLIENT: bad message %v", err)
//             } else if err != nil {
//                 log.Printf("unknown client error type %v", reflect.TypeOf(err).Elem())
//                 return
// 			} else {
// 				//c.processMessage(&msg)
// 			}
// 		}
// 	}
// }

// func main() {
//     origin := "http://localhost/"
//     url := "ws://localhost:8080/ws"
//     ws, err := websocket.Dial(url, "", origin)
//     if err != nil {
//         log.Fatal(err)
//     }
   
//     writeCh := make(chan string)
//     go listenRead(ws)
//     go listenWrite(ws, writeCh)
    
//     // loop input
//     scanner := bufio.NewScanner(os.Stdin)
//     for scanner.Scan() {
//         writeCh <- scanner.Text()
//     }
// }
