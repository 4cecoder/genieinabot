package main

import (
	"context"
	chick "github.com/byte-cats/filechick"
	gpt "github.com/sashabaranov/go-gpt3"
	"strings"
)

func OpenAICompletion(modelName string, prompt string) (string, error) {
	max := chick.StringToInt(GetMax())
	// Create a new OpenAI client
	// Replace "apiKey" with your OpenAI API key
	client := gpt.NewClient(GetAIKeyEnv())

	ctx := context.Background()
	request := gpt.CompletionRequest{
		Model:     modelName,
		MaxTokens: max,
		Prompt:    prompt,
	}
	response, err := client.CreateCompletion(ctx, request)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Text, nil
}

func refineGpt3Response(response string) string {
	response = strings.Replace(response, "Cyborg Genie", "", -1)
	//remove  Cyber Genie
	response = strings.Replace(response, "Cyber Genie", "", -1)

	// remove Cybernetic Genie
	response = strings.Replace(response, "Cybernetic Genie", "", -1)

	// Replace "Human:" with an empty string
	response = strings.Replace(response, "Human:", "", -1)

	response = strings.Split(response, ":")[1]

	// Remove any leading or trailing whitespace from the response
	response = strings.TrimSpace(response)

	// Remove any leading or trailing whitespace from the response
	response = strings.TrimSpace(response)

	// Return the refined response
	return response
}
