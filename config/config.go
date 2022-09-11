package config

import (
	"flag"

	"github.com/vharitonsky/iniflags"
)

var hostName = flag.String("hostName", "localhost", "Server host name. (Default- localhost)")
var hostPort = flag.Int("hostPort", 8080, "Server port number. (Default- 8080)")
var allowOrigins = flag.String("allowOrigins", "*", "List of allowed origins. (Default- *)")
var allowMethods = flag.String("allowMethods", "GET,POST,PUT,PATCH,DELETE", "List of allowed http methods. (Default- GET,POST,PUT,PATCH,DELETE)")

func init() {
	iniflags.Parse()
}

func GetHost() string {
	return *hostName
}
func GetPort() int {
	return *hostPort
}
func GetAllowOrigins() string {
	return *allowOrigins
}
func GetAllowMethods() string {
	return *allowMethods
}
