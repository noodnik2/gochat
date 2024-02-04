package config

import (
	"github.com/noodnik2/gochat/internal/adapter"
	"github.com/spf13/viper"
)

type AdapterType string

const (
	GeminiAdapter AdapterType = "Gemini"
	OpenAIAdapter AdapterType = "OpenAI"
)

type Config struct {
	Adapter AdapterType
	adapter.Gemini
	adapter.OpenAI
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
