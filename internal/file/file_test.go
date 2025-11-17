package file

import (
	"nevitash/gobsidain-master/internal/configuration"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func Test_makeMappingWalkFunction_withFakeEntries(t *testing.T) {
	rootDir := t.TempDir()
	level1Dir := filepath.Join(rootDir, "level1-folder")
	level2Dir := filepath.Join(level1Dir, "level2-folder")
	err := os.MkdirAll(level2Dir, os.ModePerm)
	require.NoError(t, err, "should create test directories without error")

	level1File1Path := filepath.Join(level1Dir, "level1-file1.md")
	level2File1Path := filepath.Join(level2Dir, "level2-file1.txt")
	level2File2Path := filepath.Join(level2Dir, "level2-file2.jpg")

	err = os.WriteFile(level1File1Path, []byte("Level 1 File 1 Content"), os.ModePerm)
	require.NoError(t, err, "should create level1-file1.md without error")
	err = os.WriteFile(level2File1Path, []byte("Level 2 File 1 Content"), os.ModePerm)
	require.NoError(t, err, "should create level2-file1.txt without error")
	err = os.WriteFile(level2File2Path, []byte("Level 2 File 2 Content"), os.ModePerm)
	require.NoError(t, err, "should create level2-file2.md without error")
	config := &configuration.Config{
		IncludeFilePatterns: []string{"*.md,*.txt"},
		ExcludeFilePatterns: []string{"*.png", "*.jpg"},
		ExcludePathPatterns: []string{"**/_*"},
	}
	vault, err := LoadVaultFile(rootDir, config)
	require.NoError(t, err, "LoadVaultFile should not return an error")
	assert.Equal(t, rootDir, vault.Path)
	assert.Equal(t, 1, len(vault.Children))
	assert.Equal(t, level1Dir, vault.Children[0].Path)
	assert.Equal(t, 2, len(vault.Children[0].Children))
	assert.True(t, slices.ContainsFunc(
		vault.Children[0].Children,
		func(f *File) bool {
			return f.Path == level1File1Path
		}),
		"level1-file1.md should be included",
	)
	assert.True(t, slices.ContainsFunc(
		vault.Children[0].Children,
		func(f *File) bool {
			return f.Path == level2Dir
		}),
		"level2-folder should be included",
	)
	indexLevel2Dir := slices.IndexFunc(vault.Children[0].Children, func(f *File) bool {
		return f.Path == level2Dir
	})
	require.Greater(t, indexLevel2Dir, -1, "level2-folder should be a child of level1-folder")
	assert.True(t, slices.ContainsFunc(
		vault.Children[0].Children[indexLevel2Dir].Children,
		func(f *File) bool {
			return f.Path == level2File1Path
		}),
		"level2-file1.txt should be included",
	)
	assert.False(t, slices.ContainsFunc(
		vault.Children[0].Children[indexLevel2Dir].Children,
		func(f *File) bool {
			return f.Path == level2File2Path
		}),
		"level2-file2.jpg should be excluded",
	)
}

func TestPrefixHeadings(t *testing.T) {
	var content string = `---
tags:
  - Sample
  - Tag
---
# Main Title

# Section One
## Subsection One
- First point about the topic
- Second point about the topic

# Section Two
- Another important note
- Related information here
- More details with reference [[Link]]
![Sample](image.svg)
- **Category One (Description)**
  - **Item One** (🏷️ 📝): A description of the first item with relevant details.
  - **Item Two** (📍 🔍): A description of the second item with additional context.
  - **Item Three** (⚙️ 🎯): A description of the third item with supporting information.`
	result := prefixHeaders(content)
	var expected string = `---
tags:
  - Sample
  - Tag
---
## Main Title

## Section One
### Subsection One
- First point about the topic
- Second point about the topic

## Section Two
- Another important note
- Related information here
- More details with reference [[Link]]
![Sample](image.svg)
- **Category One (Description)**
  - **Item One** (🏷️ 📝): A description of the first item with relevant details.
  - **Item Two** (📍 🔍): A description of the second item with additional context.
  - **Item Three** (⚙️ 🎯): A description of the third item with supporting information.`
	assert.Equal(t, expected, result, "Headings should be prefixed correctly")
}

func TestCombineVault(t *testing.T) {
	rootDir := t.TempDir()
	level1Dir := filepath.Join(rootDir, "level1-folder")
	level2Dir := filepath.Join(level1Dir, "level2-folder")
	err := os.MkdirAll(level2Dir, os.ModePerm)
	require.NoError(t, err, "should create test directories without error")

	level1File1Path := filepath.Join(level1Dir, "level1-file1.md")
	level2File1Path := filepath.Join(level2Dir, "level2-file1.md")

	err = os.WriteFile(level1File1Path, []byte("# Level 1 File 1\nContent of level 1 file."), os.ModePerm)
	require.NoError(t, err, "should create level1-file1.md without error")
	err = os.WriteFile(level2File1Path, []byte("# Level 2 File 1\nContent of level 2 file."), os.ModePerm)
	require.NoError(t, err, "should create level2-file1.md without error")

	template, err := getCombineTemplate()
	require.NoError(t, err, "getCombineTemplate should not return an error")

	config := &configuration.Config{
		IncludeFilePatterns: []string{"*.md"},
		ExcludeFilePatterns: []string{},
		ExcludePathPatterns: []string{},
		CombineTemplate:     *template,
		Flags: configuration.Flags{
			PrefixHeadings: true,
		},
	}
	vault, err := LoadVaultFile(rootDir, config)
	require.NoError(t, err, "LoadVaultFile should not return an error")

	combinedContent, err := CombineVault(vault, config)
	require.NoError(t, err, "CombineVault should not return an error")

	var expectedContent string = "---\n" +
		"# " + level1File1Path + "\n" +
		"---\n" +
		"## Level 1 File 1\nContent of level 1 file.\n" +
		"---\n" +
		"---\n" +
		"# " + level2File1Path + "\n" +
		"---\n" +
		"## Level 2 File 1\nContent of level 2 file.\n" +
		"---\n"
	assert.Equal(t, expectedContent, combinedContent, "Combined content should match expected output")
}

func getCombineTemplate() (*template.Template, error) {
	const templateContent = `{{range .Files}}---
# {{.Path}}
---
{{.GetContent}}
---
{{end}}`
	return template.New("template").Parse(templateContent)
}
