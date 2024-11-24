package config

import (
	"encoding/json"
	"io"
	"os"
)

type config struct {
	Path  string `json:"path"`
	Rules []rule `json:"rules"`
}

type rule struct {
	Extension string `json:"extension"`
	Directory string `json:"directory"`
}

func Read() (config, error) {
	return readConfig()
}

func readConfig() (config, error) {
	var config config

	jsonFile, err := os.Open("config.json")
	if err != nil {
		return config, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
