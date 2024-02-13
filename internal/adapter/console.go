package adapter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Consoler struct {
	input  *bufio.Scanner
	output io.Writer
	err    error
}

func NewConsole(r io.Reader, w io.Writer) *Consoler {
	return &Consoler{
		input:  bufio.NewScanner(r),
		output: w,
	}
}

func (t *Consoler) GetPrompt() string {
	t.input.Scan()

	return strings.TrimSpace(t.input.Text())
}

func (t *Consoler) Print(text string) {
	t.Printf("%s", text)
}

func (t *Consoler) Println(a ...string) {
	for argNum, arg := range a {
		if argNum > 0 {
			t.Print(" ")
		}

		t.Print(arg)
	}

	t.Print("\n")
}

func (t *Consoler) Printf(format string, args ...any) {
	if t.err == nil {
		_, t.err = t.output.Write([]byte(fmt.Sprintf(format, args...)))
	}
}
