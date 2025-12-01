package configuration

import (
	"text/template"

	"github.com/gobwas/glob"
)

type Flags struct {
	PrefixHeadings bool `yaml:"prefix-headings"`
}

var config *Config

type Config struct {
	ConfigPath          string            `yaml:"config-path"`
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
