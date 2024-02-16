package config

import (
	"strings"

	"github.com/noodnik2/gochat/internal/adapter/chatter"
	"github.com/noodnik2/gochat/internal/adapter/scriber"
	"github.com/spf13/viper"
)

type ScriberAdapterType string

type ChatterAdapterType string

const (
	GeminiChatterAdapter   ChatterAdapterType = "Gemini"
	OpenAIChatterAdapter   ChatterAdapterType = "OpenAI"
	NoScriberAdapter       ScriberAdapterType = "None"
	TemplateScriberAdapter ScriberAdapterType = "Template"
)

type ChatterConfig struct {
	Adapter       ChatterAdapterType
	DefaultPrompt string
	Prompts       map[string]string
	Adapters      struct {
		chatter.Gemini
		chatter.OpenAI
	}
}

type ScriberConfig struct {
	Adapter  ScriberAdapterType
	Adapters struct {
		scriber.TemplateScribe
	}
}

type Config struct {
	Chatter ChatterConfig
	Scriber ScriberConfig
}

func Load() (Config, error) {
	viper.SetEnvPrefix("gochat")
	viper.SetConfigName("config-local")
	viper.AddConfigPath("config")

	// the following enables you to override configuration values with environment variables
	// prefixed with "GOCHAT_"; e.g., "make run-chat GOCHAT_CHATTER_ADAPTER=Gemini"
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config

	errParse := viper.ReadInConfig()
	if errParse != nil {
		return cfg, errParse
	}

	errParse = viper.Unmarshal(&cfg)
	if errParse != nil {
		return cfg, errParse
	}

	return cfg, nil
}
