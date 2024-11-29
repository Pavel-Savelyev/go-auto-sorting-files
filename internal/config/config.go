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
	Rules []Rule `json:"file_organization_rules"`
}

type Rule struct {
	Directory      string   `json:"destination_folder"`
	FileExtensions []string `json:"file_extensions"`
	NameContains   []string `json:"name_contains"`
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
