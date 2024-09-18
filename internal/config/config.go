package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token         string
	CommandPrefix string
}

func LoadConfig() (*Config, error) {
	var err error = godotenv.Load()

	if err != nil {
		return nil, err
	}

	config := &Config{
		Token:         os.Getenv("TOKEN"),
		CommandPrefix: os.Getenv("COMMAND_PREFIX"),
	}

	return config, nil
}
