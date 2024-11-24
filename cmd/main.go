package main

import (
	"fmt"
	"go-auto-sorting-files/internal/config"
	"go-auto-sorting-files/internal/scan_folder"
	"log"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	fmt.Println("Config read successfully!")

	for _, v := range config.Rules {
		fmt.Println(v.Directory + v.Extension)
	}

	fmt.Println(config.Path)

	scan_folder.ScanFolder(config.Path)

	fileList, err := scan_folder.ScanFolder(config.Path)
	if err != nil {
		log.Fatalf("Error reading source directory: %v", err)
	}

	for _, v := range fileList.Files {
		fmt.Println(v)
	}
}
