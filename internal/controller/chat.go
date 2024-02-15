package controller

import (
	"context"
	"fmt"
	"os/user"
	"strings"
	"time"

	"github.com/noodnik2/gochat/internal/adapter/chatter"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/model"
	"github.com/noodnik2/gochat/internal/service"
)

type chatController struct {
	console     chatter.Console
	scribe      service.Scribe
	chatter     service.Chatter
	chatterName string
}

func DoChat(ctx context.Context, cfg config.Config, console chatter.Console) error {
	ctrl := &chatController{console: console}

	var err error

	if ctrl.chatter, err = service.NewChatterer(ctx, cfg.Chatter); err != nil {
		return err
	}

	defer func() { _ = ctrl.chatter.Close() }()

	if ctrl.scribe, err = service.NewScriber(cfg.Scriber); err != nil {
		return err
	}

	defer func() { _ = ctrl.scribe.Close() }()

	userName := getUsername()
	if userName == "" {
		userName = "You"
	}

	ctrl.chatterName = ctrl.chatter.Model()

	ctrl.scribe.Header(model.ScribeHeader{
		Time:    time.Now(),
		User:    userName,
		Chatter: ctrl.chatterName,
	})

	defer func() { ctrl.scribe.Footer(model.ScribeFooter{Time: time.Now()}) }()

	console.Println("gochat started")
	console.Printf("Hello %s!\n", userName)
	console.Printf("Using model: %s\n", ctrl.chatter.Model())

	defaultPromptName := cfg.Chatter.DefaultPrompt
	if defaultPromptName != "" {
		prompt := cfg.Chatter.Prompts[defaultPromptName]
		if prompt == "" {
			panic(fmt.Errorf("%w: default prompt(%s) not found", model.ErrConfig, defaultPromptName))
		}

		promptUserName := fmt.Sprintf("%s prompt", defaultPromptName)

		console.Printf("%s > %s", promptUserName, prompt)

		if !strings.HasSuffix(prompt, "\n") {
			console.Println()
		}

		ctrl.doQuery(ctx, promptUserName, prompt)
	}

	console.Println("Type 'exit' to quit")
	console.Println("Ask me anything: ")

	ctrl.doDialog(ctx, userName)

	return err
}

func (cc *chatController) doDialog(ctx context.Context, userName string) {
	for {
		cc.console.Print(fmt.Sprintf("%s > ", userName))

		prompt := cc.console.GetPrompt()
		if prompt == "" {
			continue
		}

		if prompt == "exit" {
			break
		}

		cc.doQuery(ctx, userName, prompt)
	}

	cc.console.Println("Goodbye!")
}

func (cc *chatController) doQuery(ctx context.Context, userName, prompt string) {
	cc.scribe.Entry(model.ScribeEntry{
		Time: time.Now(),
		Who:  userName,
		What: prompt,
	})

	response, tqErr := cc.chatter.MakeSynchronousTextQuery(ctx, cc.console, prompt)
	if tqErr != nil {
		panic(tqErr)
	}

	cc.scribe.Entry(model.ScribeEntry{
		Time: time.Now(),
		Who:  cc.chatterName,
		What: response,
	})
}

func getUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		return ""
	}

	return currentUser.Username
}
