package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	BaseConfig BaseConfig
	SlackConfig SlackConfig
}

type BaseConfig struct {
	Port string `toml:"port"`
}

type SlackConfig struct {
	AccessToken string
}

func NewConfig(tomlFilePath, envFilePath string) (Config, error) {
	var cfg Config

	if err := loadToml(&cfg, tomlFilePath); err != nil {
		return cfg, err
	}

	if err := loadEnv(&cfg, envFilePath); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func loadToml(cfg *Config, tomlFilePath string) error {
	_, err := toml.DecodeFile(tomlFilePath, &cfg)
	return err
}

func loadEnv(cfg *Config, envFilePath string) error {
	if err := godotenv.Load(envFilePath); err != nil {
		return err
	}

	cfg.SlackConfig.AccessToken = os.Getenv("SLACK_ACCESS_TOKEN")

	return nil
}
