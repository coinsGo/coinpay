package main

import (
	"github.com/fanguanghui/coinpay/application/controllers"
	"github.com/fanguanghui/coinpay/application/models"
	"github.com/fanguanghui/coinpay/application/utils"
	"github.com/fanguanghui/coinpay/config"
)

func main() {

	config.Setup(config.Env_DEVELOPMENT)
	utils.Setup()

	models.Setup(config.Env_DEVELOPMENT)

	controllers.Setup(config.Env_DEVELOPMENT)

	controllers.ChainPending(config.CoinId_USDT)
	controllers.ChainBlock(config.CoinId_USDT)

	controllers.ChainTxid(config.CoinId_USDT)

}
