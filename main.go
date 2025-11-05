package main

import (
	"log"
	"nevitash/gobsidain-master/internal/configuration"
)

var defaultConfigPath = "./config.yaml"
var config *configuration.Config

func main() {
	configPath := defaultConfigPath
	config, err := setupConfiguration(configPath)
	if err != nil {
		log.Fatalf("Failed to setup configuration: %v", err)
	}
	configuration.SetConfig(config)
}

func setupConfiguration(configPath string) (*configuration.Config, error) {
	config, err := configuration.LoadConfig(configPath)
	if err != nil {
		log.Printf("No config file found. Creating a new one: %v", configPath)
		config = &configuration.Config{
			ConfigPath:      configPath,
			IncludePatterns: []string{"*.md"},
		}
		err = configuration.SaveConfig(config)
	}
	if err != nil {
		log.Fatalf("Failed to create config file: %v", err)
		return nil, err
	}
	log.Printf("Config loaded: %+v", config)
	return config, nil
}
