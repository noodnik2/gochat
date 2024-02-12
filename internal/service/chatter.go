package service

import (
	"fmt"
	"io"

	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/noodnik2/gochat/internal/adapter/chatter"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/model"
)

type Chatter interface {
	MakeSynchronousTextQuery(input string, tw *adapter.Console) (string, error)
	io.Closer
}

type Chatterer struct {
	Model string
	Chatter
}

func NewChatterer(cfg config.ChatterConfig) (*Chatterer, error) {
	switch cfg.Adapter {
	case config.GeminiChatterAdapter:
		gemini, errGem := chatter.NewChatterGemini(cfg.Adapters.Gemini)
		if errGem != nil {
			return nil, errGem
		}
		return &Chatterer{
			Chatter: gemini,
			Model:   cfg.Adapters.Gemini.Model,
		}, nil
	case config.OpenAIChatterAdapter:
		openai, errOai := chatter.NewChatterOpenAI(cfg.Adapters.OpenAI)
		if errOai != nil {
			return nil, errOai
		}
		return &Chatterer{
			Chatter: openai,
			Model:   cfg.Adapters.OpenAI.Model,
		}, nil
	default:
		return nil, fmt.Errorf("%w: unrecognized 'Chatter' adapter %s",
			model.ErrConfig, cfg.Adapter)
	}
}

func (c *Chatterer) MakeSynchronousTextQuery(input string, tw *adapter.Console) (string, error) {
	return c.Chatter.MakeSynchronousTextQuery(input, tw)
}

func (c *Chatterer) Close() error {
	return c.Chatter.Close()
}
