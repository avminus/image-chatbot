package client

import (
	"fmt"

	"github.com/cohere-ai/cohere-go"
)

const (
	cohereAPIKey     = "En7bqXcgFqg6sjK9THezT9Btopd0VYRp9q9FpL88"
	endPointGenerate = "generate"
)

var (
	maxTokens   uint    = 300
	temperature float64 = 0.9
)

func GenerateText(prompt string) (string, error) {
	options := cohere.GenerateOptions{
		Model:             "command-light",
		Prompt:            prompt,
		MaxTokens:         &maxTokens,
		Temperature:       &temperature,
		StopSequences:     []string{},
		ReturnLikelihoods: "NONE",
	}
	cohereClient, err := cohere.CreateClient(cohereAPIKey)
	if err != nil {
		return "", fmt.Errorf("error in ai client creation: %v", err)
	}

	response, err := cohereClient.Generate(options)
	if err != nil {
		return "", err
	}

	return response.Generations[0].Text, nil
}
