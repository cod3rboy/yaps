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

// Default values for configuration flags
const defaultHostName = "localhost"
const defaultHostPort = 8080
const defaultAllowOrigins = "*"
const defaultAllowMethods = "GET,POST,PUT,PATCH,DELETE"
const defaultPathPrefix = "/"

// Configuration variables for application
var (
	hostName     = flag.String("hostName", defaultHostName, "Server host name")
	hostPort     = flag.Int("hostPort", defaultHostPort, "Server port number")
	pathPrefix   = flag.String("pathPrefix", defaultPathPrefix, "Prefix path for all routes")
	allowOrigins = flag.String("allowOrigins", defaultAllowOrigins, "List of allowed origins")
	allowMethods = flag.String("allowMethods", defaultAllowMethods, "List of allowed http methods")
)

// Load parses the command-line flags
func Load() {
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

// PathPrefix returns configured prefix to use with all routes
func PathPrefix() string {
	return *pathPrefix
}

// AllowOrigins returns comma-separated list of configured cors origins.
func AllowOrigins() string {
	return *allowOrigins
}

// AllowMethods returns comma-separated list of configured cors methods.
func AllowMethods() string {
	return *allowMethods
}
