package config

import (
	"encoding/json"
	"io"
	"os"
)

type AppConfig struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Path  string `json:"source_folder_path"`
	Rules []rule `json:"file_organization_rules"`
}

type rule struct {
	Extension string `json:"file_extension"`
	Directory string `json:"destination_folder"`
}

func Read() (AppConfig, error) {
	return readConfig()
}

func readConfig() (AppConfig, error) {
	var config AppConfig

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
