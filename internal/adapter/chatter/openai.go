package chatter

import (
	"context"
	"errors"
	"io"

	"github.com/noodnik2/gochat/internal/adapter"
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

func (c *ChatterOpenAI) MakeSynchronousTextQuery(input string, tw *adapter.Console) (string, error) {

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
		tw.Print(responseChunk)
	}

	c.dialog = append(c.dialog, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: responseText,
	})

	tw.Println()
	return responseText, nil
}
