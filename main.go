package main

import (
	"github.com/edwardfernando/godiary/config"
	"github.com/edwardfernando/godiary/server"
)

var (
	echoserver *server.Server
)

func main() {
	config := config.LoadApplicationConfig()
	serverReady := make(chan bool)
	echoserver = &server.Server{
		Config:      config,
		ServerReady: serverReady,
	}

	echoserver.StartEchoServer()
}
