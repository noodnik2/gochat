package main

import (
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

	if errChat := controller.DoChat(cfg, adapter.NewConsole(os.Stdin, os.Stdout)); errChat != nil {
		panic(errChat)
	}
}
