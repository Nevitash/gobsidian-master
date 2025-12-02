package configuration

import (
	"text/template"

	"github.com/Nevitash/gobsidian-master/internal/configuration"
	internalTemplate "github.com/Nevitash/gobsidian-master/internal/template"
)

func LoadConfigFromFile(path string) (*configuration.Config, error) {
	return configuration.LoadConfig(path)
}

func LoadTemplateFromFile(path string) (*template.Template, error) {
	return internalTemplate.GetTemplate(path)
}
