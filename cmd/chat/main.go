package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/model"
	"github.com/noodnik2/gochat/internal/service"
)

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

	scribe, errScribe := service.NewScribe(cfg)
	if errScribe != nil {
		panic(errScribe)
	}

	userName := "You" // TODO use something more meaningful
	chatterName := c.Model

	scribe.Header(model.Context{
		Time:         time.Now(),
		Participants: []string{userName, chatterName},
	})

	defer scribe.Footer(model.Outcome{Time: time.Now()})

	terminal := adapter.NewTerminal(os.Stdout)

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

		scribe.Entry(model.Entry{
			Time: time.Now(),
			Who:  userName,
			What: input,
		})

		response, tqErr := c.MakeSynchronousTextQuery(input, terminal)
		if tqErr != nil {
			panic(tqErr)
		}

		scribe.Entry(model.Entry{
			Time: time.Now(),
			Who:  chatterName,
			What: response,
		})
	}

	fmt.Println("Bye bye!")
}
