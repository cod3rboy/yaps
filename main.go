package main

import (
	"github.com/cod3rboy/yaps/config"
	"github.com/cod3rboy/yaps/server"
)

func main() {
	config.Load()
	server.SetupAndListen()
}
