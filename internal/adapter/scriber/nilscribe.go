package scriber

import (
	"github.com/noodnik2/gochat/internal/model"
)

type NilScribe struct{}

func (t NilScribe) Header(model.ScribeHeader) {}

func (t NilScribe) Entry(model.ScribeEntry) {}

func (t NilScribe) Footer(model.ScribeFooter) {}

func (t NilScribe) Close() error { return nil }
