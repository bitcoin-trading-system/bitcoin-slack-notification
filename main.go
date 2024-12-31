package main

import (
	"flag"

	"github.com/bitcoin-trading-system/bitcoin-slack-notification/config"
	"github.com/bitcoin-trading-system/bitcoin-slack-notification/router"
)

func main() {
	tomlFilePath := flag.String("toml", "toml/local.toml", "tomlファイルのパス")
	envFilePath := flag.String("env", "env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg, err := config.NewConfig(*tomlFilePath, *envFilePath)
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
