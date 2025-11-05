package configuration

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var config *Config

type Config struct {
	ConfigPath        string   `yaml:"config-path"`
	VaultPath         string   `yaml:"vault-path"`
	IncludePatterns   []string `yaml:"include-patterns"`
	ExcludePatterns   []string `yaml:"exclude-patterns"`
	IncludeExtentions []string `yaml:"include-extensions"`
	ExcludeExtensions []string `yaml:"exclude-extensions"`
}

func SetConfig(newConfig *Config) {
	config = newConfig
}

func GetConfig() *Config {
	return config
}

func LoadConfig(configPath string) (*Config, error) {
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("Can't read config file from %s\r\nError: %v", configPath, err)
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Printf("Can't parse config file from %s\r\nError: %v", configPath, err)
		return nil, err
	}
	config.ConfigPath = configPath
	return &config, nil
}

func SaveConfig(config *Config) error {
	ymlData, err := yaml.Marshal(config)
	if err != nil {
		log.Printf("Can't seriliaze config %v\r\nError: %v", config, err)
		return err
	}
	os.WriteFile(config.ConfigPath, ymlData, 0644)
	if err != nil {
		log.Printf("Can't save config to %w\r\nError: %v", config.ConfigPath, err)
		return err
	}
	return nil
}
