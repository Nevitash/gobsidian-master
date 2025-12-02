package configuration

import (
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/Nevitash/gobsidian-master/internal/configuration"
	internalTemplate "github.com/Nevitash/gobsidian-master/internal/template"
	"github.com/gobwas/glob"
)

func LoadConfigFromFile(path string) (*configuration.Config, error) {
	return configuration.LoadConfig(path)
}

func LoadTemplateFromFile(path string) (*template.Template, error) {
	return internalTemplate.GetTemplate(path)
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

type Flags struct {
	PrefixHeadings bool `yaml:"prefix-headings"`
}

type Config struct {
	VaultPath           string            `yaml:"vault-path"`
	IncludePathPatterns []string          `yaml:"include-patterns"`
	ExcludePathPatterns []string          `yaml:"exclude-patterns"`
	IncludeFilePatterns []string          `yaml:"include-file-patterns"`
	ExcludeFilePatterns []string          `yaml:"exclude-file-patterns"`
	Flags               Flags             `yaml:"flags"`
	CombineTemplate     template.Template `yaml:"-"`
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
