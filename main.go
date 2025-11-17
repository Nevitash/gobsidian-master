package main

import (
	"fmt"
	"log"
	"nevitash/gobsidain-master/internal/configuration"
	"nevitash/gobsidain-master/internal/file"
	"nevitash/gobsidain-master/internal/template"
	"os"
)

const (
	DEFAULT_PATH_CONFIG          = "./resources/config.yaml"
	DEFAULT_PATH_VAULT           = "./resources/vault"
	DEFAULT_PATH_OUTPUT          = "./resources/output/mega_vault.md"
	DEFAULT_TEMPLATE_PATH string = "./resources/template.md"
)

func main() {
	configPath := DEFAULT_PATH_CONFIG
	config, err := setupConfiguration(configPath)
	if err != nil {
		log.Fatalf("Failed to setup configuration: %v", err)
	}
	template, err := template.GetTemplate(DEFAULT_TEMPLATE_PATH)
	if err != nil {
		log.Fatalf("Failed to load template: %v", err)
	}
	config.CombineTemplate = *template
	configuration.SetConfig(config)
	vault, err := loadVault(DEFAULT_PATH_VAULT, configuration.GetConfig())
	if err != nil {
		log.Fatalf("Failed to load vault: %v", err)
	}
	combinedContent, err := file.CombineVault(vault, configuration.GetConfig())
	if err != nil {
		log.Fatalf("Failed to combine vault: %v", err)
	}
	err = os.WriteFile(DEFAULT_PATH_OUTPUT, []byte(combinedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}
	fmt.Printf("Combined vault written to %s\n", DEFAULT_PATH_OUTPUT)
}

func setupConfiguration(configPath string) (*configuration.Config, error) {
	config, err := configuration.LoadConfig(configPath)
	if err != nil {
		log.Printf("No config file found. Creating a new one: %v", configPath)
		config = configuration.NewDefaultConfig(configPath)
		err = configuration.SaveConfig(config)
	}
	if err != nil {
		log.Fatalf("Failed to create config file: %v", err)
		return nil, err
	}
	log.Printf("Config loaded: %+v", config)
	return config, nil
}

func loadVault(vaultPath string, config *configuration.Config) (*file.File, error) {
	vault, err := file.LoadVaultFile(vaultPath, config)
	if err != nil {
		log.Printf("Could not load vault from %s", DEFAULT_PATH_VAULT)
	}
	fmt.Printf("Vault loaded from %s: %+v\n", DEFAULT_PATH_VAULT, vault)
	return vault, nil
}
