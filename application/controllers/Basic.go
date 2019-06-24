package controllers

import (
	"github.com/fanguanghui/coinapi/wallet"
)

var WalletFactory wallet.Factory

func Setup(env string) {
	WalletFactory = wallet.Factory{Environment: env}
}
