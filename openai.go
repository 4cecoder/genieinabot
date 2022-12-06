package main

import (
	"context"
	chick "github.com/byte-cats/filechick"
	gpt "github.com/sashabaranov/go-gpt3"
)

func OpenAICompletion(modelName string, prompt string) ([]string, error) {
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
		var empty []string
		return empty, err
	}

	var responses []string
	for _, response := range response.Choices {
		responses = append(responses, response.Text)
	}

	return responses, nil
}

func FindBestGpt3Responses(input string, n int) []string {
	var bestResponses []string

	for i := 0; i < n; i++ {

		responses, _ := OpenAICompletion(GetModel(), input)
		// TODO: Find the best response

		bestResponses = append(bestResponses, responses...)
	}

	return bestResponses
}
