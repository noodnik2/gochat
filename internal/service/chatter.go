package service

import (
	"fmt"
	"io"

	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/model"
)

type Chatterer interface {
	io.Closer
	MakeSynchronousTextQuery(input string, tw adapter.Terminal) (string, error)
}

type Chatter struct {
	Model string
	Chatterer
}

func NewChatter(cfg config.Config) (*Chatter, error) {
	switch cfg.Adapter {
	case config.GeminiAdapter:
		gemini, errGem := adapter.NewChatterGemini(cfg.Gemini)
		if errGem != nil {
			return nil, errGem
		}
		return &Chatter{
			Chatterer: gemini,
			Model:     cfg.Gemini.Model,
		}, nil
	case config.OpenAIAdapter:
		openai, errOai := adapter.NewChatterOpenAI(cfg.OpenAI)
		if errOai != nil {
			return nil, errOai
		}
		return &Chatter{
			Chatterer: openai,
			Model:     cfg.OpenAI.Model,
		}, nil
	default:
		return nil, fmt.Errorf("%w: unrecognized 'Adapter' %s",
			model.ErrConfig, cfg.Scriber)
	}
}

func (c *Chatter) MakeSynchronousTextQuery(input string, tw adapter.Terminal) (string, error) {
	return c.Chatterer.MakeSynchronousTextQuery(input, tw)
}

func (c *Chatter) Close() error {
	return c.Chatterer.Close()
}
