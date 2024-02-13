package service

import (
	"fmt"

	"github.com/noodnik2/gochat/internal/adapter/scriber"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/model"
)

type Scribe interface {
	Header(context model.ScribeHeader)
	Entry(entry model.ScribeEntry)
	Footer(outcome model.ScribeFooter)
	Close() error
}

type Scriber struct {
	Scribe
}

func NewScriber(cfg config.ScriberConfig) (*Scriber, error) {
	switch cfg.Adapter {
	case config.NoScriberAdapter:
		return &Scriber{Scribe: &scriber.NilScribe{}}, nil

	case config.TemplateScriberAdapter:
		tScribe, errGem := scriber.NewTemplateScribe(cfg.Adapters.TemplateScribe)
		if errGem != nil {
			return nil, errGem
		}

		return &Scriber{Scribe: tScribe}, nil
	}

	return nil, fmt.Errorf("%w: unrecognized 'Scriber' adapter: %s",
		model.ErrConfig, cfg.Adapter)
}
