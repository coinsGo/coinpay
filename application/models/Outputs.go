package models

import (
	"time"
)

const (
	OutputStatus_Waiting  uint = 1
	OutputStatus_Pending  uint = 2
	OutputStatus_Complete uint = 3
	OutputStatus_Failure  uint = 4
)

type CoinOutput struct {
	Id          uint   `gorm:"primary_key" json:"id"`
	SendId      uint   `json:"send_id"`
	UserId      uint   `json:"user_id"`
	CoinId      uint   `json:"coin_id"`
	TxId        string `json:"tx_id"`
	TxFee       string `json:"tx_fee"`
	TxAmount    string `json:"tx_amount"`
	PlatformFee string `json:"platform_fee"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	Confirms    uint   `json:"confirms"`
	CreateTime  uint   `json:"create_time"`
	UpdateTime  uint   `json:"update_time"`
	Status      uint   `json:"status"`
}

type CoinOutputCfm struct {
	Confirms   uint `json:"confirms"`
	UpdateTime uint `json:"update_time"`
	Status     uint `json:"status"`
}

//添加申请提币记录
func OutputApply(data map[string]interface{}) (insertId uint) {

	defer Db.Close()
	Db = Connect()

	ctime := uint(time.Now().Unix())
	info := &CoinOutput{
		0,
		data["sendId"].(uint),
		data["userId"].(uint),
		data["coinId"].(uint),
		"",
		data["txFee"].(string),
		data["txAmount"].(string),
		"0",
		"",
		data["toAddress"].(string),
		0,
		ctime,
		ctime,
		OutputStatus_Waiting,
	}
	obj := Db.Hander.Table("coin_outputs").Create(info).Value
	insertId = obj.(*CoinInput).Id

	return
}

//更新提币确认数状态
func OutputUpdateTx(id, confirms, maxConfirm uint) (result bool) {

	defer Db.Close()
	Db = Connect()

	utime := uint(time.Now().Unix())
	status := InputStatus_Pending
	if maxConfirm > 0 && confirms >= maxConfirm {
		status = InputStatus_Complete
	}
	data := &CoinOutputCfm{
		confirms,
		utime,
		status,
	}
	err := Db.Hander.Table("coin_outputs").Where("id = ?", id).Updates(data).Error
	if err == nil {
		result = true
	}
	return
}

//更新提币确认数状态
func OutputUpdateCfm(id, confirms, maxConfirm uint) (result bool) {

	defer Db.Close()
	Db = Connect()

	utime := uint(time.Now().Unix())
	status := InputStatus_Pending
	if maxConfirm > 0 && confirms >= maxConfirm {
		status = InputStatus_Complete
	}
	data := &CoinOutputCfm{
		confirms,
		utime,
		status,
	}
	err := Db.Hander.Table("coin_outputs").Where("id = ?", id).Updates(data).Error
	if err == nil {
		result = true
	}
	return
}

//判断提币记录唯一性
func GetOutputExist(sendId, userId int) bool {

	defer Db.Close()
	Db = Connect()

	var count int
	obj := Db.Hander.Table("coin_outputs").Where("send_id = ? and user_id = ?", sendId, userId)
	obj.Count(&count)

	if count > 0 {
		return true
	} else {
		return false
	}
}

//获取提币记录列表
func GetOutputsByStatus(status int) (result []CoinInput) {

	defer Db.Close()
	Db = Connect()

	Db.Hander.Table("coin_outputs").Where("status = ?", status).Find(&result)
	return
}
