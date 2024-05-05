package main

import (
	"log"

	"github.com/Nazhgam/paramTracker/server"
)

// @title			param_tracker
// @version		1.0
// @description	ТЗ.
// @host			localhost:8080
// @BasePath		/
func main() {
	svr, err := server.ServerEnv(log.Default())
	if err != nil {
		return
	}

	svr.Start()
}
