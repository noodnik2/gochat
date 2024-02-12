package adapter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Console struct {
	input  *bufio.Scanner
	output io.Writer
	err    error
}

func NewConsole(r io.Reader, w io.Writer) *Console {
	return &Console{
		input:  bufio.NewScanner(r),
		output: w,
	}
}

func (t *Console) GetInput() string {
	t.input.Scan()
	return strings.TrimSpace(t.input.Text())
}

func (t *Console) Print(text string) {
	t.Printf("%s", text)
}

func (t *Console) Println(a ...string) {
	for argNum, arg := range a {
		if argNum > 0 {
			t.Print(" ")
		}
		t.Print(arg)
	}
	t.Print("\n")
}

func (t *Console) Printf(format string, args ...any) {
	if t.err == nil {
		_, t.err = t.output.Write([]byte(fmt.Sprintf(format, args...)))
	}
}
