package controller

import (
	"fmt"
	"os/user"
	"time"

	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/model"
	"github.com/noodnik2/gochat/internal/service"
)

type chatController struct {
	console     *adapter.Console
	scriber     *service.Scriber
	chatterer   *service.Chatterer
	chatterName string
}

func DoChat(cfg config.Config, console *adapter.Console) (err error) {
	cc := &chatController{console: console}
	if cc.chatterer, err = service.NewChatterer(cfg.Chatter); err != nil {
		return
	}

	defer func() { _ = cc.chatterer.Close() }()

	if cc.scriber, err = service.NewScriber(cfg.Scriber); err != nil {
		return
	}

	defer func() { _ = cc.scriber.Close() }()

	userName := getUsername()
	if userName == "" {
		userName = "You"
	}
	cc.chatterName = cc.chatterer.Model

	cc.scriber.Header(model.Context{
		Time:    time.Now(),
		User:    userName,
		Chatter: cc.chatterName,
	})

	defer func() {
		cc.scriber.Footer(model.Outcome{Time: time.Now()})
	}()

	console.Println("gochat started")
	console.Printf("Hello %s!\n", userName)
	console.Printf("Using model: %s\n", cc.chatterer.Model)

	defaultPromptName := cfg.Chatter.DefaultPrompt
	if defaultPromptName != "" {
		prompt := cfg.Chatter.Prompts[defaultPromptName]
		if prompt == "" {
			panic(fmt.Errorf("%w: default prompt(%s) not found", model.ErrConfig, defaultPromptName))
		}

		promptUserName := fmt.Sprintf("%s prompt", defaultPromptName)
		console.Printf("%s > %s", promptUserName, prompt)
		cc.doQuery(promptUserName, prompt)
	}

	console.Println("Type 'exit' to quit")
	console.Println("Ask me anything: ")

	cc.doDialog(userName)

	return
}

func (cc *chatController) doDialog(userName string) {
	for {
		cc.console.Print(fmt.Sprintf("%s > ", userName))
		prompt := cc.console.GetInput()

		if prompt == "" {
			continue
		}

		if prompt == "exit" {
			break
		}

		cc.doQuery(userName, prompt)
	}

	cc.console.Println("Goodbye!")
	return
}

func (cc *chatController) doQuery(userName, prompt string) {
	cc.scriber.Entry(model.Entry{
		Time: time.Now(),
		Who:  userName,
		What: prompt,
	})

	response, tqErr := cc.chatterer.MakeSynchronousTextQuery(prompt, cc.console)
	if tqErr != nil {
		panic(tqErr)
	}

	cc.scriber.Entry(model.Entry{
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
