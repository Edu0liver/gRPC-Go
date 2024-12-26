package main

import (
	"client-streaming/client"
	"client-streaming/server"
)

func main() {
	go server.Run()
	client.Run()
}
