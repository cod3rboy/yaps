package main

import (
	_ "github.com/cod3rboy/yaps/config"
	"github.com/cod3rboy/yaps/server"
)

func main() {
	server.SetupAndListen()
}
