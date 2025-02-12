package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

var errNoResponse = errors.New("OpenAI API did not return a response")

func prompt(gitDiff string) string {
	return `
		Please generate a concise commit message based on the following changes:

		- The commit message should begin in lowercase.
		- The commit message should begin with an imperative verb in the present tense.
		- Limit the message to no more than 10 words.
		- Do not include a period at the end of the message.
		- Focus on summarizing what was changed in the code (e.g., added, removed, fixed).

		Git Changes:
	` + gitDiff
}

func generateCommitMessage(gitDiff string, apiKey string) (string, error) {
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "developer",
				Content: "You are a helpful AI assistant that helps developers write commit messages.",
			},
			{
				Role:    "user",
				Content: prompt(gitDiff),
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", errNoResponse
	}

	return resp.Choices[0].Message.Content, nil
}
