package main

import (
	"server-streaming/client"
	"server-streaming/server"
)

func main() {
	go server.Run()
	client.Run()
}
