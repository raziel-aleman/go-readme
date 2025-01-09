package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var directoryTree string

func getDirectoryTree(path string, prefix string) error {
	// Open the directory
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the directory entries
	entries, err := file.Readdir(-1)
	if err != nil {
		return err
	}

	// Iterate over directory entries
	for _, entry := range entries {
		// Append the entry name with the appropriate prefix
		//fmt.Println(prefix + entry.Name())
		directoryTree += fmt.Sprintln(prefix + entry.Name())

		// Skip git directory contents
		if entry.Name() == ".git" {
			return nil
		}

		// If the entry is a directory, recursively print its contents
		if entry.IsDir() {

			prefixLen := len(prefix)

			// Add a new prefix for the next level
			newPrefix := "|" + strings.Repeat(" ", prefixLen) + "|- "

			// Recursively call getDirectoryTree for the subdirectory
			err := getDirectoryTree(filepath.Join(path, entry.Name()), newPrefix)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
