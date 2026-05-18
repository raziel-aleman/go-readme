package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const prompt = `You are an expert technical writer and senior software engineer. Your task is to analyze the provided project layout and the full contents of its files to generate a comprehensive, accurate, and concise README.md file. Do not use generic placeholders; pull exact details from the code except for anything that looks like environment variables or .env files, use placeholders for those. Project Files & Code at the bottom.

README Requirements
Generate a highly accurate README.md containing:

Project Overview: A concise explanation of the project's exact purpose, business logic, and architecture.

Key Features: A precise list of functionalities mapping directly to the implemented code logic.

Tech Stack & Dependencies: List the exact languages, frameworks, and critical third-party libraries found in the configuration files.

Environment Variables & Configuration: Document every single .env variable or configuration key found in the codebase, along with its expected data type or purpose.

Installation & Setup: The exact, sequential terminal commands required to clone, install dependencies, seed databases (if applicable), and run the project locally.

Usage Examples: Provide 1-2 realistic code snippets or curl commands demonstrating how to interact with the primary entry points or APIs, derived directly from the source code.

Testing: The exact command used to run the test suites found in the repository.

Make it highly scannable using tables, code blocks, and markdown checkboxes where appropriate. Ensure there is zero hallucination—if a detail isn't in the code, do not invent it.`

var DirectoryTree string

func GetDirectoryTree(path string, prefix string) error {
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
		DirectoryTree += fmt.Sprintln(prefix + entry.Name())

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
			err := GetDirectoryTree(filepath.Join(path, entry.Name()), newPrefix)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetFilesContents() string {
	var buffer bytes.Buffer

	// Traverse the directory recursively
	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, process only files
		if !d.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Append the file content to the buffer and add file name header
			fileName := "// " + path + "\n"
			buffer.Write([]byte(fileName))
			buffer.Write(data)
			// Add a newline or separator between files
			buffer.WriteString("\n")
		}
		return nil
	})

	if err != nil {
		log.Fatalf("error walking the path: %v", err)
	}

	return buffer.String()
}

func GenerateLicense(author string) string {
	return `MIT License

Copyright (c) ` + strconv.Itoa(time.Now().Year()) + " " + author + `

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`
}
