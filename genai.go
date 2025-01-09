package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

const basePrompt = `You're a professional and experienced developer and open source contributor in different technologies. Create a first release README document for these files. The target audience is professional developers with one to three years of experience building full-stack projects.

- Include a description
- A list of interesting techniques the code uses in the files provided. When possible link to MDN documentation as part of the text of the technique.
- A list of non-obvious technologies or libraries used in the code that would be of interest to professional developers with medium level experience.
- Make sure you add links to external libraries, including links to any specific fonts used.
- A breakdown of the project structure as a directory list code block: Include directories like any images directories or subfolders implied by the code, but not individual files unless they're in the root directory. Add a short description of any interesting directories underneath the code block
- If you mention a file or directory in the description, link to the file using relative links assuming you're in the root directory of the repo.
- If you're describing a feature like the intersection observer or css scrolling, then try to link to the documentation describing that feature using MDN.
- Include a How to Use section

Make sure you verify that the markdown is valid after you create it.

Avoid using verbose, indirect, or jargon-heavy phrases. Opt for straightforward, concise, and conversational language that is accessible and engaging to a broad audience. Strive for simplicity, clarity, and directness in your phrasing. It should directly engage the audience. Use a matter-of-fact tone, with fewer adjectives and a more straightforward approach. Please remain neutral.

Here is the project directory tree:`

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

	model := client.GenerativeModel("gemini-1.5-flash")
	prompt = basePrompt + "\n\n" + prompt
	//log.Println(prompt)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	response := fmt.Sprintln(resp.Candidates[0].Content.Parts[0])
	//fmt.Println(response)

	//fmt.Println(resp.Candidates[0].Content.Parts[0])
	return response
}
