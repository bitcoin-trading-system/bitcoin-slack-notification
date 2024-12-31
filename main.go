package main

import (
	"github.com/bitcoin-trading-system/bitcoin-slack-notification/config"
	"github.com/bitcoin-trading-system/bitcoin-slack-notification/router"
)

func main() {
	cfg, err := config.NewConfig("toml/local.toml", "env/.env.local")
	if err != nil {
		panic(err)
	}

	r, err := router.NewRouter(cfg)
	if err != nil {
		panic(err)
	}

	if err := r.Run(":" + cfg.BaseConfig.Port); err != nil {
		panic(err)
	}
}
