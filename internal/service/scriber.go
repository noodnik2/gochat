package service

import (
	"fmt"

	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/model"
)

type Scriber struct {
	Scribe
}

type Scribe interface {
	Header(context model.Context)
	Entry(entry model.Entry)
	Footer(outcome model.Outcome)
	Close() error
}

func NewScribe(cfg config.Config) (Scribe, error) {
	switch cfg.Scriber {
	case config.NoScriber:
		return Scriber{Scribe: &adapter.NilScribe{}}, nil
	case config.TemplateScriber:
		tScribe, errGem := adapter.NewTemplateScribe(cfg.TemplateScribe)
		if errGem != nil {
			return Scriber{}, errGem
		}
		return Scriber{Scribe: tScribe}, nil
	}

	return Scriber{}, fmt.Errorf("%w: unrecognized 'Scriber' %s",
		model.ErrConfig, cfg.Scriber)
}
