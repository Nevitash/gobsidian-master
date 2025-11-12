package filetype

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

const (
	//Texts
	MARKDOWN string = "md"
	TEXT     string = "txt"
	JSON     string = "json"
	XML      string = "xml"
	TOML     string = "toml"
	YAML     string = "yaml"
	//Images
	JPG  string = "jpg"
	JPEG string = "jpeg"
	PNG  string = "png"
	SVG  string = "svg"
	BMP  string = "bmp"
	//Applications
	PDF string = "pdf"
)

// func IsTextFile(text string) {
// 	var toCheck string
// 	toCheck = filepath.Ext(text)
// }

func extractExtensionToCheck(text string) (string, error) {
	extension := filepath.Ext(text)
	if extension != "" {
		extension = strings.TrimPrefix(extension, ".")
		return extension, nil
	}
	fileExtensions := GetAllKnownFileTypes()
	if slices.Contains(fileExtensions, text) {
		return text, nil
	}
	return "", fmt.Errorf("no file type could be extracted")
}

func GetAllKnownFileTypes() []string {
	return []string{
		MARKDOWN,
		TEXT,
		JSON,
		XML,
		TOML,
		YAML,
		JPG,
		JPEG,
		PNG,
		SVG,
		BMP,
		PDF,
	}
}
