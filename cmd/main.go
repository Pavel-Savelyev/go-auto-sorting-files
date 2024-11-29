package main

import (
	"go-auto-sorting-files/internal/config"
	"go-auto-sorting-files/internal/scan_folder"
	"go-auto-sorting-files/internal/sort_file"
	"log"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	for _, folder := range config.Folders {
		fileList, err := scan_folder.ScanFolder(folder.Path)
		if err != nil {
			log.Fatalf("Error reading source directory: %v", err)
		}

		if err := sort_file.Sort(folder, fileList); err != nil {
			log.Fatalf("Error when transferring file: %v", err)
		}
	}
}
