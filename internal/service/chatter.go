package service

import (
	"io"

	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/noodnik2/gochat/internal/config"
)

type Chatterer interface {
	io.Closer
	MakeSynchronousTextQuery(input string, scribe adapter.Scribe) error
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
		panic("configuration 'Adapter' property not set")
	}
}

func (c *Chatter) MakeSynchronousTextQuery(input string, scribe adapter.Scribe) error {
	return c.Chatterer.MakeSynchronousTextQuery(input, scribe)
}

func (c *Chatter) Close() error {
	return c.Chatterer.Close()
}
