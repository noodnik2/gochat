package chatter

import (
	"context"
	"errors"
	"io"

	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	APIKey string
	Model  string
}

type OpenAIChatter struct {
	client *openai.Client
	model  string
	dialog []openai.ChatCompletionMessage
}

func NewOpenAIChatter(cCfg OpenAI) (*OpenAIChatter, error) {
	client := openai.NewClient(cCfg.APIKey)

	return &OpenAIChatter{client: client, model: cCfg.Model}, nil
}

func (c *OpenAIChatter) Model() string {
	return c.model
}

func (c *OpenAIChatter) Close() error {
	return nil
}

func (c *OpenAIChatter) MakeSynchronousTextQuery(ctx context.Context, console Console, prompt string) (string, error) {
	message := openai.ChatCompletionMessage{ //nolint:exhaustruct
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	}

	c.dialog = append(c.dialog, message)

	resp, errCc := c.client.CreateChatCompletionStream(
		ctx,
		openai.ChatCompletionRequest{ //nolint:exhaustruct
			Model:    c.model,
			Messages: c.dialog,
		},
	)

	if errCc != nil {
		return "", errCc
	}

	var responseText string

	for {
		response, errRecv := resp.Recv()

		if errRecv != nil {
			if !errors.Is(errRecv, io.EOF) {
				return "", errRecv
			}

			break
		}

		responseChunk := response.Choices[0].Delta.Content
		responseText += responseChunk

		console.Print(responseChunk)
	}

	c.dialog = append(c.dialog, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: responseText,
	})

	console.Println()

	return responseText, nil
}
