package obsidian

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileProperties(t *testing.T) {
	testContent, err := os.ReadFile("./resources/obsidian.property.test.md")
	assert.NoError(t, err, "Should read test markdown file without error")
	testString := string(testContent)
	properties, err := GetFileProperties(testString)

	assert.NoError(t, err, "GetFileProperties should not return an error")
	assert.NotNil(t, properties["tags"], "Properties should contain 'tags'")
	assert.NotNil(t, properties["aliases"], "Properties should contain 'aliases'")

	tags := properties["tags"].([]any)
	aliases := properties["aliases"].([]any)

	assert.Equal(t, 2, len(tags), "Tags should contain 2 items")
	assert.Equal(t, 3, len(aliases), "Aliases should contain 3 items")
	assert.Contains(t, tags, "Lore", "Tags should contain 'Lore'")
	assert.Contains(t, tags, "Beats", "Tags should contain 'Beats'")
	assert.Contains(t, aliases, "Test", "Aliases should contain 'Test'")
	assert.Contains(t, aliases, "Testcase", "Aliases should contain 'Testcase'")
	assert.Contains(t, aliases, "Unittest", "Aliases should contain 'Unittest'")
}
