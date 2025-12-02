package file

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/Nevitash/gobsidian-master/configuration"

	"github.com/Nevitash/gobsidian-master/internal/template"
	"github.com/gobwas/glob"
)

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
	config.VaultPath = path
	incluidePathGlob := config.GetIncludePathGlob()
	excludePathGlob := config.GetExcludePathGlob()
	includeFileGlob := config.GetIncludeFileGlob()
	excludeFileGlob := config.GetExcludeFileGlob()
	var walkFunction = makeMappingWalkFunction(
		vault,
		incluidePathGlob,
		excludePathGlob,
		includeFileGlob,
		excludeFileGlob,
		config,
	)
	filepath.WalkDir(path, walkFunction)
	return vault, nil
}

func makeMappingWalkFunction(
	result *File,
	includePathGlob glob.Glob,
	excludePathGlob glob.Glob,
	includeFileGlob glob.Glob,
	excludeFileGlob glob.Glob,
	configuration *configuration.Config,
) func(string, fs.DirEntry, error) error {
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
		if !ShouldBeProcessed(path, includePathGlob, excludePathGlob) {
			return filepath.SkipDir
		}
		if !dirEntry.IsDir() && !ShouldBeProcessed(path, includeFileGlob, excludeFileGlob) {
			return nil
		}
		if result.Path == path {
			return nil
		}
		// Build the File node
		node := &File{
			Path:          path,
			FileExtension: filepath.Ext(path),
			Config:        configuration,
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

func ShouldBeProcessed(path string, include glob.Glob, exclude glob.Glob) bool {
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

func CombineVault(vault *File, config *configuration.Config) (string, error) {
	if vault == nil {
		return "", fmt.Errorf("vault is nil")
	}
	files, err := vault.GetFiles()
	if err != nil {
		return "", fmt.Errorf("no files matching criteria in vault: %v", err)
	}
	var templateData = &TemplateData{
		Files: files,
	}
	content, err := template.RenderTemplate(&config.CombineTemplate, templateData)
	if err != nil {
		return "", fmt.Errorf("failed to render template: %v", err)
	}
	return content, nil
}
