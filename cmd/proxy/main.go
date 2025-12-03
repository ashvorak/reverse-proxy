package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reverse-proxy/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Configuration file required.")
	}

	config, err := config.Load(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to load config file %v", err)
	}

	log.Println("Service started with configuration:")
	for _, r := range config.Routes {
		log.Printf("Prefix %s, Upstream %s\n", r.Prefix, r.Upstream)
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
