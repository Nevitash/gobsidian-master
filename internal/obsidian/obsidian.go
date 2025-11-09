package obsidian

import (
	"regexp"

	"gopkg.in/yaml.v3"
)

func GetFileProperties(content string) (map[string]interface{}, error) {
	re := regexp.MustCompile(`(?s)^---\s(.*?)\s---\s`)
	properties := make(map[string]interface{})
	matches := re.FindStringSubmatch(content)
	if len(matches) < 2 {
		return properties, nil
	}
	match := matches[1]
	err := yaml.Unmarshal([]byte(match), properties)
	if err != nil {
		return nil, err
	}
	return properties, nil
}
