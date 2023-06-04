package main

import (
	"github.com/web_server_golang/src"
	"log"
)

func main() {
	server := src.NewServer(":3000")
	log.Fatal(server.Listen())
	
}
