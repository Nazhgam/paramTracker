package main

import (
	"log"

	"github.com/Nazhgam/paramTracker/server"
)

func main() {
	server, err := server.ServerEnv(log.Default())
	if err != nil {
		return
	}
	server.Start()
}
