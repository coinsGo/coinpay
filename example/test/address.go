package test

import (
	"log"

	"github.com/go-develop/coinpay/application/models"
)

func Address() {

	//添加地址记录
	data := map[string]interface{}{
		"userId":  1,
		"coinId":  1,
		"address": "erferferergergerg",
		"account": "",
	}
	id := models.AddressSave(data)
	log.Println(id)

	//判断地址是否存在
	b := models.GetAddressExist(1, "erferferergergerg")
	log.Println(b)

}
