package chatter

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Gemini struct {
	APIKey string
	Model  string
}

type GeminiChatter struct {
	client  *genai.Client
	session *genai.ChatSession
	model   string
}

func NewGeminiChatter(ctx context.Context, gCfg Gemini) (*GeminiChatter, error) {
	gClient, gcErr := genai.NewClient(ctx, option.WithAPIKey(gCfg.APIKey))
	if gcErr != nil {
		return nil, gcErr
	}

	session := gClient.GenerativeModel(gCfg.Model).StartChat()

	return &GeminiChatter{
		client:  gClient,
		session: session,
		model:   gCfg.Model,
	}, nil
}

func (c *GeminiChatter) Model() string {
	return c.model
}

func (c *GeminiChatter) Close() error {
	return c.client.Close()
}

func (c *GeminiChatter) MakeSynchronousTextQuery(ctx context.Context, console Console, prompt string) (string, error) {
	iter := c.session.SendMessageStream(ctx, genai.Text(prompt))

	var responseBuilder strings.Builder

	for {
		resp, errNext := iter.Next()

		if errors.Is(errNext, iterator.Done) {
			break
		}

		if errNext != nil {
			return "", errNext
		}

		for _, candidate := range resp.Candidates {
			for _, part := range candidate.Content.Parts {
				responseBuilder.WriteString(fmt.Sprintf("%s", part))
				console.Printf("%s", part)
			}
		}
	}

	console.Println()

	return responseBuilder.String(), nil
}
