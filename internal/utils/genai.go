package utils

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func GenerateReadme(prompt string, files string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env files")
	}

	ctx := context.Background()
	clientConfig := &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	}

	client, err := genai.NewClient(ctx, clientConfig)
	parts := []*genai.Part{
		{Text: prompt},
		{InlineData: &genai.Blob{Data: []byte(files), MIMEType: "text/plain"}},
	}

	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", []*genai.Content{{Parts: parts}}, nil)
	if err != nil {
		log.Fatalf("error generating content from Gemini: %v", err)
	}

	return result.Text()
}
