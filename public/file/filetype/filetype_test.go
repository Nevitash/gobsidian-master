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

func TestIsTextFile(t *testing.T) {
	assert.True(t, IsTextFile("/root/test."+MARKDOWN), "MARKDOWN should be recognized as a text file")
	assert.True(t, IsTextFile("/root/test."+TEXT), "TEXT should be recognized as a text file")
	assert.True(t, IsTextFile("/root/test."+JSON), "JSON should be recognized as a text file")
	assert.True(t, IsTextFile("/root/test."+XML), "XML should be recognized as a text file")
	assert.True(t, IsTextFile("/root/test."+TOML), "TOML should be recognized as a text file")
	assert.True(t, IsTextFile("/root/test."+YAML), "YAML should be recognized as a text file")
	assert.False(t, IsTextFile("/root/test."+PNG), "PNG should not be recognized as a text file")
}

func TestIsImageFile(t *testing.T) {
	assert.True(t, IsImageFile("/root/test."+PNG), "PNG should be recognized as an image file")
	assert.True(t, IsImageFile("/root/test."+JPEG), "JPEG should be recognized as an image file")
	assert.True(t, IsImageFile("/root/test."+JPG), "JPG should be recognized as an image file")
	assert.True(t, IsImageFile("/root/test."+BMP), "BMP should be recognized as an image file")
	assert.True(t, IsImageFile("/root/test."+SVG), "SVG should be recognized as an image file")
	assert.False(t, IsImageFile("/root/test."+MARKDOWN), "MARKDOWN should not be recognized as an image file")
}

func TestIsApplicationFile(t *testing.T) {
	assert.True(t, IsApplicationFile("/root/test."+PDF), "PDF should be recognized as an application file")
	assert.False(t, IsApplicationFile("/root/test."+JSON), "JSON should not be recognized as an application file")
}
