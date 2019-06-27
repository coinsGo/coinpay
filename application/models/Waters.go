package models

import (
	"time"
)

const (
	WaterAssets_Input  uint = 1 //存币
	WaterAssets_Output uint = 2 //提币
)

const (
	WaterFreeze_OutWait uint = 100 //提币冻结
	WaterFreeze_OutSucc uint = 101 //提币成功
	WaterFreeze_OutFail uint = 102 //提币失败
)

const (
	WaterPended_InWait uint = 200 //存币待确认
	WaterPended_InSucc uint = 201 //存币成功
	WaterPended_InFail uint = 202 //存币失败
)

type CoinWater struct {
	Id         uint    `gorm:"primary_key" json:"id"`
	UserId     uint    `json:"user_id"`
	CoinId     uint    `json:"coin_id"`
	TypeId     uint    `json:"type_id"`
	Assid      uint    `json:"assid"`
	Amount     float64 `json:"amount"`
	CreateTime uint    `json:"create_time"`
	Remark     string  `json:"remark"`
}

//保存资金流水记录
func WaterSave(data map[string]interface{}) (insertId uint) {

	defer Db.Close()
	Db = Connect()

	info := &CoinWater{
		0,
		data["userId"].(uint),
		data["coinId"].(uint),
		data["typeId"].(uint),
		data["assid"].(uint),
		data["amount"].(float64),
		uint(time.Now().Unix()),
		data["remark"].(string),
	}
	obj := Db.Hander.Table("coin_waters").Create(info).Value
	insertId = obj.(*CoinWater).Id
	return
}
