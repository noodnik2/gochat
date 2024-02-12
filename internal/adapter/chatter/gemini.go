package chatter

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/noodnik2/gochat/internal/adapter"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Gemini struct {
	ApiKey string
	Model  string
}

type ChatterGemini struct {
	gc    *genai.Client
	cs    *genai.ChatSession
	model string
	ctx   context.Context
}

func NewChatterGemini(gcfg Gemini) (*ChatterGemini, error) {
	ctx := context.Background()
	gc, gcErr := genai.NewClient(ctx, option.WithAPIKey(gcfg.ApiKey))
	if gcErr != nil {
		return nil, nil
	}

	cs := gc.GenerativeModel(gcfg.Model).StartChat()

	return &ChatterGemini{ctx: ctx, gc: gc, cs: cs, model: gcfg.Model}, nil
}

func (c *ChatterGemini) Close() error {
	return c.gc.Close()
}

func (c *ChatterGemini) MakeSynchronousTextQuery(input string, tw *adapter.Console) (string, error) {
	iter := c.cs.SendMessageStream(c.ctx, genai.Text(input))

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
				responseBuilder.Write([]byte(fmt.Sprintf("%s", part)))
				tw.Printf("%s", part)
			}
		}
	}

	tw.Println()
	return responseBuilder.String(), nil
}