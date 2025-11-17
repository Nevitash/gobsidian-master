package template

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"
)

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
		return "", fmt.Errorf("no template passed in")

	}
	var buffer bytes.Buffer
	err := template.Execute(&buffer, data)
	if err != nil {
		return "", errors.New("error rendering template: " + err.Error())
	}
	return buffer.String(), nil
}
