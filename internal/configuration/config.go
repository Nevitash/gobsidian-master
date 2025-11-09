package configuration

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobwas/glob"
	"gopkg.in/yaml.v3"
)

var config *Config

type Config struct {
	ConfigPath          string   `yaml:"config-path"`
	VaultPath           string   `yaml:"vault-path"`
	IncludePathPatterns []string `yaml:"include-patterns"`
	ExcludePathPatterns []string `yaml:"exclude-patterns"`
	IncludeFilePatterns []string `yaml:"include-file-patterns"`
	ExcludeFilePatterns []string `yaml:"exclude-file-patterns"`
}

func (c *Config) GetIncludePathGlob() glob.Glob {
	return getPatternGlob(c.IncludePathPatterns)
}

func (c *Config) GetExcludePathGlob() glob.Glob {
	return getPatternGlob(c.ExcludePathPatterns)
}

func (c *Config) GetIncludeFileGlob() glob.Glob {
	return getPatternGlob(c.IncludeFilePatterns)
}

func (c *Config) GetExcludeFileGlob() glob.Glob {
	return getPatternGlob(c.ExcludeFilePatterns)
}

func getPatternGlob(patterns []string) glob.Glob {
	globPattern, err := prepareGlobPattern(patterns)
	if err != nil {
		return nil
	}
	compiledGlob, err := glob.Compile(globPattern)
	if err != nil {
		log.Printf("Error compiling exclude glob pattern: %v", err)
		return nil
	}
	return compiledGlob
}

func prepareGlobPattern(patterns []string) (string, error) {
	if len(patterns) == 0 {
		return "", fmt.Errorf("no patterns provided")
	}
	joinedPatterns := strings.Join(patterns, ",")
	return fmt.Sprintf("{%v}", joinedPatterns), nil
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
		log.Printf("Can't save config to %s\r\nError: %v", config.ConfigPath, err)
		return err
	}
	return nil
}
