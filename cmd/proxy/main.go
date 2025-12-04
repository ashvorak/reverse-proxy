package main

import (
	"log"
	"os"
	"reverse-proxy/internal/config"
	"reverse-proxy/internal/router"
	"reverse-proxy/internal/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Configuration file required.")
	}

	config, err := config.Load(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to load config file %v", err)
	}

	router, err := router.New(config.Routes)
	if err != nil {
		log.Fatalf("Failed to create router: %v", err)
	}

	server, err := server.NewServer(router)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	server.Start()
}
