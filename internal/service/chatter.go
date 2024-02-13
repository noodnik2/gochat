package service

import (
	"context"
	"fmt"
	"io"

	"github.com/noodnik2/gochat/internal/adapter/chatter"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/model"
)

type Chatter interface {
	Model() string
	MakeSynchronousTextQuery(ctx context.Context, console chatter.Console, prompt string) (string, error)
	io.Closer
}

type Chatterer struct {
	Chatter
}

func NewChatterer(ctx context.Context, cfg config.ChatterConfig) (*Chatterer, error) {
	switch cfg.Adapter {
	case config.GeminiChatterAdapter:
		gemini, errGem := chatter.NewGeminiChatter(ctx, cfg.Adapters.Gemini)
		if errGem != nil {
			return nil, errGem
		}

		return &Chatterer{Chatter: gemini}, nil

	case config.OpenAIChatterAdapter:
		openai, errOai := chatter.NewOpenAIChatter(cfg.Adapters.OpenAI)
		if errOai != nil {
			return nil, errOai
		}

		return &Chatterer{Chatter: openai}, nil

	default:
		return nil, fmt.Errorf("%w: unrecognized 'Chatter' adapter %s",
			model.ErrConfig, cfg.Adapter)
	}
}

func (c *Chatterer) MakeSynchronousTextQuery(ctx context.Context, console chatter.Console, prompt string) (
	string, error,
) {
	return c.Chatter.MakeSynchronousTextQuery(ctx, console, prompt)
}

func (c *Chatterer) Close() error {
	return c.Chatter.Close()
}
