package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	newgenai "google.golang.org/genai"
)

const basePrompt = `You're a professional, experienced developer and open source contributor in different technologies. Create a first release README document for these files. The target audience is professional developers with one to three years of experience building full-stack projects.

- Include a description
- A list of interesting techniques the code uses in the files provided. When possible link to MDN documentation as part of the text of the technique.
- A list of technologies or libraries used in the code that would be of interest to professional developers with medium level experience, otherwise indicate if the project only uses the standard library for the respective programming language or framework.
- Make sure you add links to external libraries, including links to any specific fonts used.
- A breakdown of the project structure as a directory list code block: Include directories like any images directories or subfolders implied by the code, but not individual files unless they're in the root directory. Add a short description of any interesting directories underneath the code block
- If you mention a file or directory in the description, link to the file using relative links assuming you're in the root directory of the repo.
- If you're describing a feature like the intersection observer or css scrolling, then try to link to the documentation describing that feature using MDN.
- Include a How to Use section

Make sure you verify that the markdown is valid after you create it.

Avoid using verbose, indirect, or jargon-heavy phrases. Opt for straightforward, concise, and conversational language that is accessible and engaging to a broad audience. Strive for simplicity, clarity, and directness in your phrasing. It should directly engage the audience. Use a matter-of-fact tone, with fewer adjectives and a more straightforward approach. Please remain neutral.

Here is the project directory tree:`

const betterPrompt = `You are an expert technical writer and senior software engineer. Your task is to analyze the provided project layout and the full contents of its files to generate a comprehensive, accurate, and concise README.md file. Do not use generic placeholders; pull exact details from the code. Project Files & Code at the bottom.

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

func getResponse(prompt string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-flash-latest")
	prompt = basePrompt + "\n\n" + prompt
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	response := fmt.Sprintln(resp.Candidates[0].Content.Parts[0])

	return response
}

func newGetResponse(prompt string, files string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	ctx := context.Background()
	clientConfig := &newgenai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: newgenai.BackendGeminiAPI,
	}

	client, err := newgenai.NewClient(ctx, clientConfig)
	parts := []*newgenai.Part{
		{Text: prompt},
		{InlineData: &newgenai.Blob{Data: []byte(files), MIMEType: "text/plain"}},
	}
	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", []*newgenai.Content{{Parts: parts}}, nil)

	return result.Text()
}

func getFilesContents() string {
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

			// Append the file content to the buffer
			fileName := "// " + path + "\n"
			buffer.Write([]byte(fileName))
			buffer.Write(data)
			// Optional: add a newline or separator between files
			buffer.WriteString("\n")
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v", err)
	}

	return buffer.String()
}
