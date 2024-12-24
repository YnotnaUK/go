package dotenv

import (
	"bytes"
	"fmt"
	"os"
)

var filename string = ".env"
var debug bool = false

type DotEnvConfig struct {
	Filename string
	Debug    bool
}

func Load() error {
	// Attempt to load file
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	// split file into lines
	fileLines := bytes.Split(fileContents, []byte("\n"))
	// use each line to set variables
	for _, fileLine := range fileLines {
		// Split each line into a key and value
		keyValues := bytes.SplitN(fileLine, []byte("="), 2)
		key := string(keyValues[0])
		value := string(keyValues[1])
		// Set each key value as env
		os.Setenv(key, value)
		if debug {
			fmt.Printf("Set key %s to value %s\n", key, value)
		}
	}
	return nil
}

func LoadWithConfig(config DotEnvConfig) {
	filename = config.Filename
	debug = config.Debug
	Load()
}
