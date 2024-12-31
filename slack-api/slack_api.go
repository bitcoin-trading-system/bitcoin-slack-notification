package slackapi

import (
	"errors"

	"github.com/bitcoin-trading-system/bitcoin-slack-notification/config"
)

type SlackAPI struct {
	Config config.Config
}

func NewSlackAPI(cfg *config.Config) (*SlackAPI, error) {
	if cfg.SlackConfig.AccessToken == "" {
		return nil, errors.New("slack access token is empty")
	}

	return &SlackAPI{Config: *cfg}, nil
}
