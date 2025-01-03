package handler

import (
	"github.com/bitcoin-trading-system/bitcoin-slack-notification/config"
	slackapi "github.com/bitcoin-trading-system/bitcoin-slack-notification/slack-api"
)

type Handler struct {
	Config   config.Config
	SlackAPI slackapi.SlackAPI
}

func NewHandler(cfg config.Config) (*Handler, error) {
	slackAPI, err := slackapi.NewSlackAPI(&cfg)
	if err != nil {
		return nil, err
	}

	return &Handler{Config: cfg, SlackAPI: *slackAPI}, nil
}
