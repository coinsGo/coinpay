package controllers

import (
	"log"

	"github.com/fanguanghui/coinpay/application/models"
	"github.com/fanguanghui/coinpay/config"
)

func Pending(coinId int) {
	client := WalletFactory.Get(config.CoinName[config.CoinId_USDT])
	listTx := client.GetPendingTxs("")

	//var info map[string]interface{}

	for _, row := range listTx {
		for _, vout := range row.Vouts {
			addr := models.GetAddress(coinId, vout.Address)
			log.Println(addr)
		}

		//info["userId"] = 1

		//models.InputSave(info)
	}
}
