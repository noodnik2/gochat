package config

import (
	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/spf13/viper"
)

type AdapterType string

type ScriberType string

const (
	GeminiAdapter AdapterType = "Gemini"
	OpenAIAdapter AdapterType = "OpenAI"
	NoScriber     ScriberType = "None"
	TemplateScriber ScriberType = "Template"
)

type Config struct {
	Adapter AdapterType
	adapter.Gemini
	adapter.OpenAI
	Scriber ScriberType
	adapter.TemplateScribe
}

func Load() (cfg Config, errParse error) {
	viper.SetConfigName("config-local")
	viper.AddConfigPath("config")

	if errParse = viper.ReadInConfig(); errParse != nil {
		return
	}
	errParse = viper.Unmarshal(&cfg)
	return
}
