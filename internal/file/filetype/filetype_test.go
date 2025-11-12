package filetype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractExtensionToCheck(t *testing.T) {
	result, err := extractExtensionToCheck(".md")
	assert.NoError(t, err)
	assert.Equal(t, "md", result, "'md' should be extracted from '.md'")

	result, err = extractExtensionToCheck("md")
	assert.NoError(t, err)
	assert.Equal(t, "md", result, "'md' should be extracted from 'md'")

	result, err = extractExtensionToCheck("test.md")
	assert.NoError(t, err)
	assert.Equal(t, "md", result, "'md' should be extracted from 'test.md'")

	result, err = extractExtensionToCheck("/root/level1/test.md")
	assert.NoError(t, err)
	assert.Equal(t, "md", result, "'md' should be extracted from '/root/level1/test.md'")

	result, err = extractExtensionToCheck("./root/level1/test.md")
	assert.NoError(t, err)
	assert.Equal(t, "md", result, "'md' should be extracted from './root/level1/test.md'")

	result, err = extractExtensionToCheck("root/level1/test.md")
	assert.NoError(t, err)
	assert.Equal(t, "md", result, "'md' should be extracted from 'root/level1/test.md'")

	result, err = extractExtensionToCheck("/root/level1/test.jpg")
	assert.NoError(t, err)
	assert.Equal(t, "jpg", result, "'jpg' should be extracted from '/root/level1/test.jpg'")

	result, err = extractExtensionToCheck("./root/level1/test.pdf")
	assert.NoError(t, err)
	assert.Equal(t, "pdf", result, "'pdf' should be extracted from './root/level1/test.pdf'")

	result, err = extractExtensionToCheck("root/level1/test.json")
	assert.NoError(t, err)
	assert.Equal(t, "json", result, "'json' should be extracted from 'root/level1/test.json'")

	_, err = extractExtensionToCheck("root/level1/errorNoExtension")
	assert.Error(t, err, "Nothing to extract from 'root/level1/errorNoExtension'. Should throw an error")
}
