package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGlobs(t *testing.T) {
	testConfig := &Config{
		IncludePatterns: []string{"*.md", "*.txt"},
		ExcludePatterns: []string{"**/_images", "**/_images/*", "**/_config", "**/_config/*"},
	}
	includeGlob, err := testConfig.GetIncludeGlob()
	assert.NoError(t, err, "GetIncludeGlob should not return an error")
	excludeGlob, err := testConfig.GetExcludeGlob()

	assert.NoError(t, err, "GetExcludeGlob should not return an error")
	assert.True(t, includeGlob.Match("003 - Characters/note.md"), "Include glob should match 003 - Characters/note.md")
	assert.True(t, includeGlob.Match("./_test/haha.txt"), "Include glob should match ./_test/haha.txt")
	assert.False(t, includeGlob.Match("note.png"), "Include glob should not match note.png")
	assert.True(t, excludeGlob.Match("./_config"), "Exclude glob should match ./_config")
	assert.True(t, excludeGlob.Match("003 - Characters/_images"), "Exclude glob should match ./003 - Characters/_images")
	assert.False(t, excludeGlob.Match("zzz - Info/__config"), "Exclude glob should not match zzz - Info/__config")
}
