package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Route struct {
	Prefix   string `yaml:"prefix"`
	Upstream string `yaml:"upstream"`
}

type Config struct {
	Routes []Route `yaml:"routes"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error marshaling yaml config: %v", err)
	}

	return &config, nil
}
