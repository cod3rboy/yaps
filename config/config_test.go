package config

import (
	"flag"
	"strconv"
	"testing"
)

// LoadFlags parses configuration flags if not already parsed
func LoadFlags() {
	if !flag.Parsed() {
		Load()
	}
}

type TestData struct {
	FlagArg  string
	Expected string
}

var testDataHost = []TestData{
	{FlagArg: "", Expected: defaultHostName},
	{FlagArg: "127.0.0.1", Expected: "127.0.0.1"},
	{FlagArg: "0.0.0.0", Expected: "0.0.0.0"},
}

func TestHost(t *testing.T) {
	LoadFlags()
	for _, data := range testDataHost {
		if data.FlagArg != "" {
			flag.Set("hostName", data.FlagArg)
		}
		actual := Host()
		if actual != data.Expected {
			t.Errorf("expected = %s, actual = %s\n", data.Expected, actual)
		}
	}
}

var testDataPort = []TestData{
	{FlagArg: "", Expected: strconv.Itoa(defaultHostPort)},
	{FlagArg: "3000", Expected: "3000"},
}

func TestPort(t *testing.T) {
	LoadFlags()
	for _, data := range testDataPort {
		if data.FlagArg != "" {
			flag.Set("hostPort", data.FlagArg)
		}
		actual := Port()
		expected, err := strconv.Atoi(data.Expected)
		if err != nil {
			t.Fatal(err)
		}
		if actual != expected {
			t.Errorf("expected = %d, actual = %d\n", expected, actual)
		}
	}
}

var testAllowOriginsData = []TestData{
	{FlagArg: "", Expected: defaultAllowOrigins},
	{FlagArg: "localhost,example.com", Expected: "localhost,example.com"},
}

func TestAllowOrigins(t *testing.T) {
	LoadFlags()
	for _, data := range testAllowOriginsData {
		if data.FlagArg != "" {
			flag.Set("allowOrigins", data.FlagArg)
		}
		actual := AllowOrigins()
		if actual != data.Expected {
			t.Errorf("expected = %s, actual = %s\n", data.Expected, actual)
		}
	}
}

var testAllowMethodsData = []TestData{
	{FlagArg: "", Expected: defaultAllowMethods},
	{FlagArg: "GET", Expected: "GET"},
	{FlagArg: "POST,PUT", Expected: "POST,PUT"},
	{FlagArg: "DELETE,PATCH,GET", Expected: "DELETE,PATCH,GET"},
}

func TestAllowMethods(t *testing.T) {
	LoadFlags()
	for _, data := range testAllowMethodsData {
		if data.FlagArg != "" {
			flag.Set("allowMethods", data.FlagArg)
		}
		actual := AllowMethods()
		if actual != data.Expected {
			t.Errorf("expected = %s, actual = %s\n", data.Expected, actual)
		}
	}
}
