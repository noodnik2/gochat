package adapter

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/noodnik2/gochat/internal/model"
)

type Templates struct {
	Header string
	Entry  string
	Footer string
}

type TemplateScribe struct {
	SaveDir   string
	SaveFile  string
	Templates Templates
}

type TemplateScribeImpl struct {
	writer     *os.File
	headerTmpl *template.Template
	entryTmpl  *template.Template
	footerTmpl *template.Template
}

func NewTemplateScribe(cfg TemplateScribe) (*TemplateScribeImpl, error) {

	tFuncs := template.FuncMap{
		"split": func(s, sep string) []string {
			return strings.Split(s, sep)
		},
	}

	var (
		tsi TemplateScribeImpl
		err error
	)

	if tsi.headerTmpl, err = template.New("header").Funcs(tFuncs).Parse(cfg.Templates.Header); err != nil {
		return nil, err
	}

	if tsi.entryTmpl, err = template.New("entry").Funcs(tFuncs).Parse(cfg.Templates.Entry); err != nil {
		return nil, err
	}

	if tsi.footerTmpl, err = template.New("footer").Funcs(tFuncs).Parse(cfg.Templates.Footer); err != nil {
		return nil, err
	}

	ofn := fmt.Sprintf("%s-%s", time.Now().Format("060102150405"), cfg.SaveFile)

	if cfg.SaveDir != "" {
		ofn = path.Join(cfg.SaveDir, ofn)
	}

	if tsi.writer, err = os.Create(ofn); err != nil {
		return nil, err
	}

	return &tsi, nil
}

func (t TemplateScribeImpl) Header(context model.Context) {
	execTemplate(t.headerTmpl, t.writer, context)
}

func (t TemplateScribeImpl) Entry(entry model.Entry) {
	execTemplate(t.entryTmpl, t.writer, entry)
}

func (t TemplateScribeImpl) Footer(outcome model.Outcome) {
	execTemplate(t.footerTmpl, t.writer, outcome)
}

func (t TemplateScribeImpl) Close() error {
	return t.writer.Close()
}

func execTemplate(tmpl *template.Template, w io.Writer, data any) {
	err := tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
