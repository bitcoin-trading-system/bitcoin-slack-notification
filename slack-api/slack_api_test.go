package slackapi

import "github.com/bitcoin-trading-system/bitcoin-slack-notification/config"

var SlackAPITestCfg config.Config

func init() {
	var err error
	SlackAPITestCfg, err = config.NewConfig("../toml/local.toml", "../env/.env.local")
	if err != nil {
		panic(err)
	}
}
