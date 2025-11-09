package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFiles(t *testing.T) {
	rootFolder := &File{Path: "./"}
	level1Folder1 := &File{Path: "./level1-folder1", Parent: rootFolder}
	level1Folder2 := &File{Path: "./level1-folder2", Parent: rootFolder}
	rootFolder.Children = []*File{level1Folder1, level1Folder2}

	level1File1 := &File{Path: "./level1-folder1/level1-file1.md", Parent: level1Folder1, FileExtension: "md"}
	level2Folder1 := &File{Path: "./level1-folder1/level2-folder1", Parent: level1Folder1}
	level1Folder1.Children = []*File{level1File1, level2Folder1}

	level2File1 := &File{Path: "./level1-folder1/level2-folder1/level2-file1.md", Parent: level2Folder1, FileExtension: "md"}
	level2File2 := &File{Path: "./level1-folder1/level2-folder1/level2-file2.md", Parent: level2Folder1, FileExtension: "md"}
	level2Folder1.Children = []*File{level2File1, level2File2}

	expectedFiles := []*File{level1File1, level2File1, level2File2}
	foundFiles, err := rootFolder.GetFiles()

	assert.NoError(t, err, "GetFiles should not return an error")
	assert.Equal(t, expectedFiles, foundFiles, "GetFiles should return [./level1-folder1/level1-file1.md, "+
		"./level1-folder1/level2-folder1/level2-file1.md, "+
		" ./level1-folder1/level2-folder1/level2-file2.md]")
}
