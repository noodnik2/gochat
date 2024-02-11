package adapter

import (
	"github.com/noodnik2/gochat/internal/model"
)

type NilScribe struct{}

func (t NilScribe) Header(model.Context) {}

func (t NilScribe) Entry(model.Entry) {}

func (t NilScribe) Footer(model.Outcome) {}

func (t NilScribe) Close() error { return nil }
