package main

import (
	"flag"
	"golang-test-6-socket-chat/chat"
)

var (
	port = flag.String("p", ":8080", "set port")
)

func init() {
	flag.Parse()
}

func main() {
	chat.Start(*port)
}
