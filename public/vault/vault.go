package vault

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/Nevitash/gobsidian-master/configuration"
	"github.com/Nevitash/gobsidian-master/internal/file"
)

func MergeToFile(
	vaultPath string,
	outputPath string,
	includeFilePatterns []string,
	excludeFilePatterns []string,
	includePathPatterns []string,
	excludePathPatterns []string,
	flags *configuration.Flags,
	outputTemplate *template.Template,
) (string, error) {
	config, err := CreateConfig(vaultPath, includeFilePatterns, excludeFilePatterns, includePathPatterns, excludePathPatterns, flags, outputTemplate)
	if err != nil {
		return "", err
	}
	return MergeToFileWithConfig(outputPath, config)
}

func MergeToFileWithConfig(outputPath string, config *configuration.Config) (string, error) {
	content, err := MergeToStringWithConfig(outputPath, config)
	if err != nil {
		return "", err
	}
	outputPath = filepath.Clean(outputPath)
	outputPath, err = filepath.Abs(outputPath)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return "", err
	}
	err = os.WriteFile(outputPath, []byte(content), 0o644)
	if err != nil {
		return "", nil
	}
	return content, nil
}

func MergeToString(
	vaultPath string,
	outputPath string,
	includeFilePatterns []string,
	excludeFilePatterns []string,
	includePathPatterns []string,
	excludePathPatterns []string,
	flags *configuration.Flags,
	outputTemplate *template.Template,
) (string, error) {
	config, err := CreateConfig(vaultPath, includeFilePatterns, excludeFilePatterns, includePathPatterns, excludePathPatterns, flags, outputTemplate)
	if err != nil {
		return "", err
	}
	return MergeToStringWithConfig(outputPath, config)
}

func MergeToStringWithConfig(outputPath string, config *configuration.Config) (string, error) {
	vault, err := file.LoadVaultFile(config.VaultPath, config)
	if err != nil {
		return "", err
	}
	content, err := file.CombineVault(vault, config)
	if err != nil {
		return "", err
	}
	return content, err
}

func CreateConfig(
	vaultPath string,
	includeFilePatterns []string,
	excludeFilePatterns []string,
	includePathPatterns []string,
	excludePathPatterns []string,
	flags *configuration.Flags,
	outputTemplate *template.Template,
) (*configuration.Config, error) {
	if includeFilePatterns == nil {
		includeFilePatterns = []string{"*.md", "*.txt"}
	}
	if excludeFilePatterns == nil {
		excludeFilePatterns = []string{"*.png", "*.jpg"}
	}
	if includePathPatterns == nil {
		includePathPatterns = []string{}
	}
	if excludePathPatterns == nil {
		excludePathPatterns = []string{"**/_*", "**/.*"}
	}
	if flags == nil {
		flags = &configuration.Flags{
			PrefixHeadings: true,
		}
	}
	if outputTemplate == nil {
		defaultTemplate, err := template.New("template").Parse(getDefaultOutputTemplateString())
		if err != nil {
			return nil, err
		}
		outputTemplate = defaultTemplate
	}
	return &configuration.Config{
		VaultPath:           vaultPath,
		IncludePathPatterns: includePathPatterns,
		ExcludePathPatterns: excludePathPatterns,
		IncludeFilePatterns: includeFilePatterns,
		ExcludeFilePatterns: excludeFilePatterns,
		Flags:               *flags,
		CombineTemplate:     *outputTemplate,
	}, nil
}

func getDefaultOutputTemplateString() string {
	return `{{range .Files}}---
# {{.Path}}
---
{{.GetContent -}}
---
{{end}}`
}
