package config

import (
	"os"
)

var GoschoolHome = ""

func init() {
	env := ""
	it, exists := os.LookupEnv("GOSCHOOL")
	if exists {
		env = it
	}
	it, exists = os.LookupEnv("GOSCHOOL_HOME")
	if exists {
		env = it
	}
	GoschoolHome = env
}
