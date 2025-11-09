package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGlobs(t *testing.T) {
	testConfig := &Config{
		IncludePathPatterns: []string{"*003 - Characters*", "*_test*"},
		ExcludePathPatterns: []string{"**/_images", "**/_images/*", "**/_config", "**/_config/*"},
	}
	includeGlob := testConfig.GetIncludePathGlob()
	excludeGlob := testConfig.GetExcludePathGlob()

	assert.NotNil(t, includeGlob)
	assert.NotNil(t, excludeGlob)

	assert.True(t, includeGlob.Match("003 - Characters/level2/note.md"), "Include glob should match '003 - Characters/level2/note.md'")
	assert.True(t, includeGlob.Match("./_test/haha.txt"), "Include glob should match './_test/haha.txt'")
	assert.False(t, includeGlob.Match("note.png"), "Include glob should not match 'note.png'")
	assert.True(t, excludeGlob.Match("./_config"), "Exclude glob should match './_config'")
	assert.True(t, excludeGlob.Match("003 - Characters/level2/_images"), "Exclude glob should match './003 - Characters/level2/_images'")
	assert.False(t, excludeGlob.Match("zzz - Info/__config"), "Exclude glob should not match 'zzz - Info/__config'")
}
