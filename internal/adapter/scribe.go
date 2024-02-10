package adapter

import "fmt"

type Scribe struct{}

func (s Scribe) Open() error {
	s.Printf("Starting Chat!\n")
	return nil
}

func (s Scribe) Close() error {
	s.Printf("Chat Closed\n")
	return nil
}

func (Scribe) Print(s string) {
	fmt.Print(s)
}

func (Scribe) Printf(s string, args ...any) {
	fmt.Printf(s, args...)
}

func (Scribe) Println() {
	fmt.Println()
}
