package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/urfin-djuice/layout/pkg/config"
)

type Config struct {
	Example ExampleConfig `mapstructure:",squash"`
	Loop    LoopConfig    `mapstructure:",squash"`
}

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Skip it.")
	}

	conf := &Config{}

	err := config.Unmarshal(conf)
	if err != nil {
		panic(fmt.Sprintf("Can't create config: %v", err))
	}

	return conf
}
