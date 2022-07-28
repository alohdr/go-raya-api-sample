package main

import (
	"go-rest-api/config"
	"go-rest-api/server"
)

func main() {
	config.Init()
	server.NewServer()
}
