package main

import (
	"github.com/fanguanghui/coinpay/application/controllers"
	"github.com/fanguanghui/coinpay/config"
)

func main() {

	config.Setup(config.Env_DEVELOPMENT)
	controllers.Setup(config.Env_DEVELOPMENT)

	controllers.Pending(config.CoinId_USDT)
}
