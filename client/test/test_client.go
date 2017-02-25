package main

import (
    "log"
    "bufio"
    "os"
    c "API/client"
)

// func getLobbyCode() []interface{} {
//     resp, err := http.Get("http://localhost:8080/lobby/gen")
//     if err != nil {
//         log.Fatal(err)
//     }
//     data, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         log.Fatal(err)
//     }
//     resp.Body.Close()
//     fmt.Printf("%s\r\n", data)
//     return data
// }
func messageHandler() func(interface{}) {
    return func(msg interface{}){
        log.Printf("message handler %v \r\n", msg)
    }
}
func main() {
    //_ = getLobbyCode()

    client := c.NewClient(
        "ws://localhost:8080/connect",
        "http://localhost/",
        messageHandler(),
    )
    client.Connect()
    
    // loop input
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        data := scanner.Bytes()
        log.Printf("data: %v", data)
        client.Write(data)
    }
}