package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/service"
)

// see: https://github.com/sashabaranov/go-openai

func main() {
	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		panic(cfgErr)
	}

	c, chErr := service.NewChatter(cfg)
	if chErr != nil {
		panic(chErr)
	}

	defer func() { _ = c.Close() }()
	s := bufio.NewScanner(os.Stdin)

	fmt.Printf("Using model: %s\n", c.Model)
	fmt.Println("Type 'exit' to quit")
	fmt.Println("Ask me anything: ")
	for {
		fmt.Print("> ")
		s.Scan()
		input := strings.TrimSpace(s.Text())

		if input == "" {
			continue
		}
		
		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if tqErr := c.MakeSynchronousTextQuery(input); tqErr != nil {
			panic(tqErr)
		}
	}

	fmt.Println("Bye bye!")
}
