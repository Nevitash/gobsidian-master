package file

import (
	"fmt"
	"nevitash/gobsidain-master/internal/configuration"
	"os"
	"regexp"
)

type FileProperty struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

const (
	PATTERN_HEADER_DETECTION    string = `(?m)^(#+)`
	PATTERN_HEADER_SUBSTITUTION string = `#$1`
)

type File struct {
	Parent        *File                 `yaml:"parent"`
	Children      []*File               `yaml:"children"`
	Path          string                `yaml:"path"`
	FileExtension string                `yaml:"file-extension"`
	Properties    []FileProperty        `yaml:"properties"`
	Config        *configuration.Config `yaml:"-"`
}

func (f *File) GetFiles() ([]*File, error) {
	var files []*File
	for _, child := range f.Children {
		isFile, err := child.IsFile()
		if err != nil {
			return nil, fmt.Errorf("error checking if path %s is file: %v", child.Path, err)
		}
		if !isFile {
			childFiles, err := child.GetFiles()
			if err != nil {
				return nil, fmt.Errorf("error getting files from path %s: %v", child.Path, err)
			}
			files = append(files, childFiles...)
		} else {
			files = append(files, child)
		}
	}
	return files, nil
}

func (f *File) IsFile() (bool, error) {
	return f.FileExtension != "", nil
}

func (f *File) GetContent() (string, error) {
	if exists, err := IsFile(f.Path); err == nil && exists {
		return "", fmt.Errorf("path %s is either not accessible, was deleted or is not a file.\r\nerror: %v", f.Path, err)
	}
	content, err := os.ReadFile(f.Path)
	if err != nil {
		return "", err
	}
	stringContent := string(content)
	if f.Config != nil && f.Config.Flags.PrefixHeadings {
		stringContent = prefixHeaders(stringContent)
	}
	return stringContent, nil
}

func prefixHeaders(content string) string {
	replacer := regexp.MustCompile(PATTERN_HEADER_DETECTION)
	return replacer.ReplaceAllString(content, PATTERN_HEADER_SUBSTITUTION)
}
