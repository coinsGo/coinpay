package models

import (
	"time"
)

const (
	OutputStatus_Waiting  int = 1
	OutputStatus_Pending  int = 2
	OutputStatus_Complete int = 3
)

type CoinOutput struct {
	Id          int     `gorm:"primary_key" json:"id"`
	SendId      int     `json:"send_id"`
	UserId      int     `json:"user_id"`
	CoinId      int     `json:"coin_id"`
	TxId        string  `json:"tx_id"`
	TxFee       float64 `json:"tx_fee"`
	TxAmount    float64 `json:"tx_amount"`
	PlatformFee float64 `json:"platform_fee"`
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	Confirms    int     `json:"confirms"`
	CreateTime  int     `json:"create_time"`
	UpdateTime  int     `json:"update_time"`
	Status      int     `json:"status"`
}

type CoinOutputCfm struct {
	Confirms   int `json:"confirms"`
	UpdateTime int `json:"update_time"`
	Status     int `json:"status"`
}

//添加申请提币记录
func OutputApply(data map[string]interface{}) (insertId int) {

	defer Db.Close()
	Db = Connect()

	ctime := int(time.Now().Unix())
	info := &CoinOutput{
		0,
		data["sendId"].(int),
		data["userId"].(int),
		data["coinId"].(int),
		"",
		data["txFee"].(float64),
		data["txAmount"].(float64),
		0.0,
		"",
		data["toAddress"].(string),
		0,
		ctime,
		ctime,
		1,
	}
	obj := Db.Hander.Table("coin_outputs").Create(info).Value
	insertId = obj.(*CoinInput).Id

	return
}

//更新提币确认数状态
func OutputUpdateTx(id, confirms, maxConfirm int) (result bool) {

	defer Db.Close()
	Db = Connect()

	utime := int(time.Now().Unix())
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
func OutputUpdateCfm(id, confirms, maxConfirm int) (result bool) {

	defer Db.Close()
	Db = Connect()

	utime := int(time.Now().Unix())
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
