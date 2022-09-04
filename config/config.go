package config

import (
	"os"
	"path"

	"github.com/vharitonsky/iniflags"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	iniflags.SetConfigFile(path.Join(wd, "config/config.ini"))
	iniflags.Parse()
}
