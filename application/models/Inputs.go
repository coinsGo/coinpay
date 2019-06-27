package models

import (
	"time"
)

const (
	InputStatus_Pending  uint = 1
	InputStatus_Complete uint = 2
	InputStatus_Failure  uint = 3
)

type CoinInput struct {
	Id          uint    `gorm:"primary_key" json:"id"`
	UserId      uint    `json:"user_id"`
	CoinId      uint    `json:"coin_id"`
	TxId        string  `json:"tx_id"`
	TxAmount    float64 `json:"tx_amount"`
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	Confirms    uint64  `json:"confirms"`
	CollectId   uint    `json:"collect_id"`
	CreateTime  uint    `json:"create_time"`
	UpdateTime  uint    `json:"update_time"`
	Status      uint    `json:"status"`
}

type CoinInputCfm struct {
	Confirms   uint64 `json:"confirms"`
	UpdateTime uint   `json:"update_time"`
	Status     uint   `json:"status"`
}

//添加存币记录
func InputSave(data map[string]interface{}) (insertId uint) {

	defer Db.Close()
	Db = Connect()

	coinId := data["coinId"].(uint)
	confirms := data["confirms"].(uint64)
	status := data["status"].(uint)
	ctime := uint(time.Now().Unix())

	info := &CoinInput{
		0,
		data["userId"].(uint),
		coinId,
		data["txId"].(string),
		data["txAmount"].(float64),
		data["fromAddress"].(string),
		data["toAddress"].(string),
		confirms,
		0,
		ctime,
		ctime,
		status,
	}
	obj := Db.Hander.Table("coin_inputs").Create(info).Value
	insertId = obj.(*CoinInput).Id

	return
}

//更新存币确认数状态
func UpdateInputCfm(data map[string]interface{}) (result bool) {

	defer Db.Close()
	Db = Connect()

	id := data["id"].(uint)
	confirms := data["confirms"].(uint64)
	status := data["status"].(uint)

	info := &CoinInputCfm{
		confirms,
		uint(time.Now().Unix()),
		status,
	}
	err := Db.Hander.Table("coin_inputs").Where("id = ?", id).Updates(info).Error
	if err == nil {
		result = true
	}
	return
}

//判断地址是否存在
func GetInputExist(coinId uint, txId string) bool {

	defer Db.Close()
	Db = Connect()

	var count int
	obj := Db.Hander.Table("coin_inputs").Where("coin_id = ? and tx_id = ?", coinId, txId)
	obj.Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//获取存币记录
func GetPendingInputs(coinId uint) (result []CoinInput) {

	defer Db.Close()
	Db = Connect()

	Db.Hander.Table("coin_inputs").Where("coin_id = ? and status = 1", coinId).Find(&result)
	return
}

//计算存币记录状态
func GetInputStatus(coinId uint, confirms uint64, valid bool) uint {
	status := InputStatus_Pending
	if !valid {
		status = InputStatus_Failure
	} else {
		if CfgConfirms[coinId] > 0 {
			if CfgConfirms[coinId] <= confirms {
				status = InputStatus_Complete
			}
		}
	}
	return status
}
