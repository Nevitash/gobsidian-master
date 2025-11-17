package template

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"
)

const (
	DEFAULT_TEMPLATE_PATH string = "./resources/template.md"
)

func GetDefaultTemplate() (*template.Template, error) {
	return GetTemplate(DEFAULT_TEMPLATE_PATH)
}

func GetTemplate(path string) (*template.Template, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("no template file found with path %s: %v", path, err)
		} else {
			return nil, fmt.Errorf("error reading template file %s: %v", path, err)
		}
	}
	return template.New("template").Parse(string(content))
}

func RenderTemplate(template *template.Template, data any) (string, error) {
	if template == nil {
		var err error
		template, err = GetDefaultTemplate()
		if err != nil {
			return "", fmt.Errorf("no template passed in and couldn't load default template: %s", DEFAULT_TEMPLATE_PATH)
		}
	}
	var buffer bytes.Buffer
	err := template.Execute(&buffer, data)
	if err != nil {
		return "", errors.New("error rendering template: " + err.Error())
	}
	return buffer.String(), nil
}
