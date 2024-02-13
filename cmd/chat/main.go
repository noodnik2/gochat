package main

import (
	"context"
	"os"

	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/noodnik2/gochat/internal/config"
	"github.com/noodnik2/gochat/internal/controller"
)

func main() {
	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		panic(cfgErr)
	}

	console := adapter.NewConsole(os.Stdin, os.Stdout)

	if errChat := controller.DoChat(context.Background(), cfg, console); errChat != nil {
		panic(errChat)
	}
}
