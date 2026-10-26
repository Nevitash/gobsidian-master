# Gobsidian Master _(gobsidian-master)_

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

Go core library for parsing, filtering, and merging Obsidian knowledge vaults.

Gobsidian Master provides the underlying programmatic Go API to crawl Obsidian notes, resolve links, apply glob filters, and merge markdown vaults into structured single-document outputs.

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [API](#api)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Background

Obsidian notes use interconnected markdown files and links. Consolidating or preprocessing these vaults programmatically requires robust path resolution, glob pattern filtering, and template rendering. Gobsidian Master extracts this core functionality into a reusable Go module utilized by CLI tools like gobsidian-master-dungeon.

## Install

Requires Go 1.25+:

`ash
go get github.com/Nevitash/gobsidian-master@latest
`

## Usage

Import and call vault merge methods directly in your Go code:

`go
package main

import (
	"log"
	"github.com/Nevitash/gobsidian-master/public/vault"
)

func main() {
	output, err := vault.MergeToFile(
		"/path/to/vault",
		"/path/to/output.md",
		[]string{"*.md"},
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to merge vault: %v", err)
	}
	log.Printf("Merged output written: %s", output)
}
`

## API

### ault.MergeToFile(vaultPath, outputPath, includeFilePatterns, excludeFilePatterns, includePathPatterns, excludePathPatterns, flags, outputTemplate)

- Merges all matching markdown files in aultPath into outputPath.
- Returns: (string, error)

### ault.MergeToStringWithConfig(outputPath, config)

- Compiles vault content into an in-memory formatted markdown string.
- Returns: (string, error)

## Maintainers

[@Nevitash](https://github.com/Nevitash)

## Contributing

PRs accepted. Please open an issue first to discuss public API adjustments or new configuration options.

## License

[MIT](LICENSE)
