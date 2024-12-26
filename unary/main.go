package main

import (
	"unary/client"
	"unary/server"
)

func main() {
	go server.Run()
	client.Run()
}
