package main

import (
	"github.com/go-develop/coinpay/application/controllers"
	"github.com/go-develop/coinpay/application/models"
	"github.com/go-develop/coinpay/application/utils"
	"github.com/go-develop/coinpay/config"
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
