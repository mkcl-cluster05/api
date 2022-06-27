package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	Mode string `json:"mode"`
}

var Config config

func SetupConfig(filename string) error {

	jsFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jsFile.Close()

	bytes, err := ioutil.ReadAll(jsFile)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &Config)
	return nil
}
