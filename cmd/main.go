package main

import (
	"fmt"
	"go-auto-sorting-files/internal/config"
	"log"
)

func main() {
	config, err := config.Init()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	fmt.Println("Config read successfully!")

	for _, v := range config.Rules {
		fmt.Println(v.Directory + v.Extension)
	}
}
