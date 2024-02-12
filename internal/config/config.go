package config

import (
	"github.com/noodnik2/gochat/internal/adapter/chatter"
	"github.com/noodnik2/gochat/internal/adapter/scriber"
	"github.com/spf13/viper"
)

type AdapterType string

const (
	GeminiChatterAdapter   AdapterType = "Gemini"
	OpenAIChatterAdapter   AdapterType = "OpenAI"
	NoScriberAdapter       AdapterType = "None"
	TemplateScriberAdapter AdapterType = "Template"
)

type ChatterConfig struct {
	Adapter       AdapterType
	DefaultPrompt string
	Prompts       map[string]string
	Adapters      struct {
		chatter.Gemini
		chatter.OpenAI
	}
}

type ScriberConfig struct {
	Adapter  AdapterType
	Adapters struct {
		scriber.TemplateScribe
	}
}

type Config struct {
	Chatter ChatterConfig
	Scriber ScriberConfig
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
