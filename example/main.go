package main

import (
	"log"

	"github.com/fanguanghui/coinpay/example/test"

	"github.com/fanguanghui/coinpay/config"
)

func main() {
	config.Setup(config.Env_DEVELOPMENT)

	log.Println(config.Global)
	log.Println(config.Envdata)

	test.Inputs()
}
