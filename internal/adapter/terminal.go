package adapter

import (
	"fmt"
	"io"
)

type Terminal struct {
	output io.Writer
	err    error
}

func NewTerminal(w io.Writer) Terminal {
	return Terminal{output: w}
}

func (t Terminal) Print(text string) {
	t.Printf("%s", text)
}

func (t Terminal) Println() {
	t.Print("\n")
}

func (t Terminal) Printf(format string, args ...any) {
	if t.err == nil {
		_, t.err = t.output.Write([]byte(fmt.Sprintf(format, args...)))
	}
}
