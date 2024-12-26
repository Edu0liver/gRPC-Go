package main

import (
	"bidirectional-streaming/client"
	"bidirectional-streaming/server"
)

func main() {
	go server.Run()
	go client.Run()

	// Block forever
	select {}
}
