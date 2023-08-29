package client

import (
	"fmt"
	"image-chatbot/utils"
	"io/ioutil"
	"log"

	"github.com/cohere-ai/cohere-go"
)

var (
	maxTokens   uint    = 300
	temperature float64 = 0.9
)

const KEYS_DIRECTORY = "/keys/"

func GenerateText(prompt string) (string, error) {
	options := cohere.GenerateOptions{
		Model:             "command-light",
		Prompt:            prompt,
		MaxTokens:         &maxTokens,
		Temperature:       &temperature,
		StopSequences:     []string{},
		ReturnLikelihoods: "NONE",
	}

	keysFilePath, err := utils.GetPath("cohere_key.txt", KEYS_DIRECTORY)
	if err != nil {
		return "", fmt.Errorf("error in getting the keys path: %v", err)
	}

	key, err := ioutil.ReadFile(keysFilePath)
	if err != nil {
		log.Fatal(err)
	}

	cohereClient, err := cohere.CreateClient(string(key))
	if err != nil {
		return "", fmt.Errorf("error in ai client creation: %v", err)
	}

	response, err := cohereClient.Generate(options)
	if err != nil {
		return "", err
	}

	return response.Generations[0].Text, nil
}
