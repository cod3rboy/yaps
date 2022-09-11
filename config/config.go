// Package config provides the functions to access the configuration values.
//
// Configuration values are either loaded from config.ini file or from command line arguments.
// The command line arguments have priority over config.ini file.
// In case no values are configured, the default value for each configuration variable is used.
package config

import (
	"flag"

	"github.com/vharitonsky/iniflags"
)

// Configuration variables for application
var (
	hostName     = flag.String("hostName", "localhost", "Server host name. (Default- localhost)")
	hostPort     = flag.Int("hostPort", 8080, "Server port number. (Default- 8080)")
	allowOrigins = flag.String("allowOrigins", "*", "List of allowed origins. (Default- *)")
	allowMethods = flag.String("allowMethods", "GET,POST,PUT,PATCH,DELETE", "List of allowed http methods. (Default- GET,POST,PUT,PATCH,DELETE)")
)

func init() {
	// Parse flag arguments/ini file
	iniflags.Parse()
}

// Host returns configured server hostname.
func Host() string {
	return *hostName
}

// Port returns configured server port number.
func Port() int {
	return *hostPort
}

// AllowOrigins returns comma-separated list of configured cors origins.
func AllowOrigins() string {
	return *allowOrigins
}

// AllowMethods returns comma-separated list of configured cors methods.
func AllowMethods() string {
	return *allowMethods
}
