package file

import (
	"fmt"
	"io/fs"
	"log"
	"nevitash/gobsidain-master/internal/configuration"
	"os"
	"path/filepath"

	"github.com/gobwas/glob"
)

type FileProperty struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

type File struct {
	Parent        *File          `yaml:"parent"`
	Children      []*File        `yaml:"children"`
	Path          string         `yaml:"path"`
	FileExtension string         `yaml:"file-extension"`
	Properties    []FileProperty `yaml:"properties"`
}

func (f *File) GetContent() (string, error) {
	if exists, err := IsFile(f.Path); err == nil && exists {
		return "", fmt.Errorf("path %s is either not accessible, was deleted or is not a file.\r\nerror: %v", f.Path, err)
	}
	content, err := os.ReadFile(f.Path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func IsFile(path string) (bool, error) {
	if exists, err := FileExists(path); err == nil && exists {
		return false, fmt.Errorf("File %s is either not accessible or was deleted\r\nerror: %v", path, err)
	}
	fi, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return fi.Mode().IsRegular(), nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func LoadVaultFile(path string, config *configuration.Config) (*File, error) {
	if exists, err := FileExists(path); err != nil || !exists {
		return nil, err
	}
	var vault = &File{
		Path:          path,
		FileExtension: filepath.Ext(path),
	}
	var walkFunction = makeMappingWalkFunction(vault, config.GetIncludePathGlob(), config.GetExcludePathGlob())
	filepath.WalkDir(path, walkFunction)
	return vault, nil
}

func makeMappingWalkFunction(result *File, includeGlob glob.Glob, excludeGlob glob.Glob) func(string, fs.DirEntry, error) error {
	fileMap := map[string]*File{
		result.Path: result,
	}
	return func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			if dirEntry.IsDir() {
				return filepath.SkipDir
			}
			return err
		}
		if !shouldProcess(path, includeGlob, excludeGlob) {
			return nil
		}
		// Build the File node
		node := &File{
			Path:          filepath.Base(path),
			FileExtension: filepath.Ext(path),
		}

		// Determine parent path
		parentPath := filepath.Dir(path)
		parent, ok := fileMap[parentPath]
		if !ok {
			// If parent not found, default to root
			parent = result
		}

		// Link bidirectionally
		node.Parent = parent
		parent.Children = append(parent.Children, node)

		// Store node for future children
		if dirEntry.IsDir() {
			fileMap[path] = node
		}

		return nil
	}
}

func shouldProcess(path string, include glob.Glob, exclude glob.Glob) bool {
	if include != nil && !include.Match(path) {
		log.Printf("Path %s didn't match include patterns", path)
		return false
	}
	if exclude != nil && exclude.Match(path) {
		log.Printf("Path %s matched exclude patterns", path)
		return false
	}
	return true
}

func walkAndMapFiles(oath string, dirEntry fs.DirEntry, err error) error {
	//TODO: implement this function
	return nil
}
