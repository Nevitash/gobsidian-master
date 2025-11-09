package main

import (
	"log"
	"nevitash/gobsidain-master/internal/configuration"
	"nevitash/gobsidain-master/internal/file"
)

var DEFAULT_PATH_CONFIG = "./resources/config.yaml"
var DEFAULT_PATH_VAULT = "./resources/vault"
var config *configuration.Config

func main() {
	configPath := DEFAULT_PATH_CONFIG
	config, err := setupConfiguration(configPath)
	if err != nil {
		log.Fatalf("Failed to setup configuration: %v", err)
	}
	configuration.SetConfig(config)
	file.LoadVaultFile(DEFAULT_PATH_VAULT, configuration.GetConfig())
}

func setupConfiguration(configPath string) (*configuration.Config, error) {
	config, err := configuration.LoadConfig(configPath)
	if err != nil {
		log.Printf("No config file found. Creating a new one: %v", configPath)
		config = &configuration.Config{
			ConfigPath:          configPath,
			IncludeFilePatterns: []string{"*.md"},
			ExcludeFilePatterns: []string{"*.png", "*.jpg"},
			ExcludePathPatterns: []string{"**/_*"},
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
