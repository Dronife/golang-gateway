package main

import (
	"fmt"
    "test/userbackend/config/server"
)


func main() {
	fmt.Println("MAIN")
    var server *server.HttpServer = server.Construct("8081")
    go server.StartListening()


	select {}
}
