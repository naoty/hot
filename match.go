package main

import (
	"path/filepath"
	"strings"
)

func Match(pattern string, filenames []string) []string {
	var matched []string

	pattern, _ = filepath.Abs(pattern)
	globbed, _ := filepath.Glob(pattern)

	for _, filename := range filenames {
		abs, _ := filepath.Abs(filename)
		if contains(abs, globbed) {
			matched = append(matched, filename)
		}
	}

	return matched
}

func contains(filename string, filenames []string) bool {
	for _, f := range filenames {
		if strings.EqualFold(filename, f) {
			return true
		}
	}
	return false
}
