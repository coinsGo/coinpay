package test

import (
	"log"

	"github.com/fanguanghui/coinpay/application/models"
)

func Inputs() {

	//创建存币记录
	data := map[string]interface{}{
		"userId":      1,
		"coinId":      1,
		"txId":        "1124rtyerger434",
		"txAmount":    1.0,
		"fromAddress": "456456refgdsf",
		"toAddress":   "456456refgdsf",
		"confirms":    0,
	}
	id := models.InputSave(data, 0)
	log.Println(id)

	//验证TXID唯一性
	b1 := models.GetInputExist(1, "1124rtyerger434")
	log.Println(b1)

	//更新确认数状态
	b2 := models.UpdateInputCfm(1, 1, 3)
	log.Println(b2)

	//获取待确认记录
	res := models.GetPendingInputs()
	log.Println(res)

}
