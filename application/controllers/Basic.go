package controllers

import (
	"github.com/go-develop/coinapi/wallet"
)

var WalletFactory wallet.Factory

func Setup(env string) {
	WalletFactory = wallet.Factory{Environment: env}
}
