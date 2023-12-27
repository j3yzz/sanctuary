package config

import (
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

func Provide() Config {
	var cfg Config

	var k = koanf.New("")

	if err := k.Load(file.Provider("./.env"), dotenv.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	return cfg
}
