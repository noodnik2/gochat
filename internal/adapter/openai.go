package adapter

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	ApiKey string
	Model  string
}

type ChatterOpenAI struct {
	ctx    context.Context
	client *openai.Client
	model  string
	dialog []openai.ChatCompletionMessage
}

func NewChatterOpenAI(ccfg OpenAI) (*ChatterOpenAI, error) {
	ctx := context.Background()
	client := openai.NewClient(ccfg.ApiKey)
	return &ChatterOpenAI{ctx: ctx, client: client, model: ccfg.Model}, nil
}

func (c *ChatterOpenAI) Close() error {
	return nil
}

func (c *ChatterOpenAI) MakeSynchronousTextQuery(input string) error {

	message := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: input,
	}

	c.dialog = append(c.dialog, message)

	resp, errCc := c.client.CreateChatCompletionStream(
		c.ctx,
		openai.ChatCompletionRequest{
			Model:    c.model,
			Messages: c.dialog,
		},
	)
	if errCc != nil {
		return errCc
	}

	var responseText string
	for {
		response, errRecv := resp.Recv()
		if errRecv != nil {
			if !errors.Is(errRecv, io.EOF) {
				return errRecv
			}
			break
		}
		responseChunk := response.Choices[0].Delta.Content
		responseText += responseChunk
		fmt.Print(responseChunk)
	}

	c.dialog = append(c.dialog, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: responseText,
	})

	fmt.Println()
	return nil
}
