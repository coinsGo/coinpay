package models

import (
	"time"
)

const (
	InputStatus_Pending  int = 1
	InputStatus_Complete int = 2
)

type CoinInput struct {
	Id          int     `gorm:"primary_key" json:"id"`
	UserId      int     `json:"user_id"`
	CoinId      int     `json:"coin_id"`
	TxId        string  `json:"tx_id"`
	TxAmount    float64 `json:"tx_amount"`
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	Confirms    int     `json:"confirms"`
	CollectId   int     `json:"collect_id"`
	CreateTime  int     `json:"create_time"`
	UpdateTime  int     `json:"update_time"`
	Status      int     `json:"status"`
}

type CoinInputCfm struct {
	Confirms   int `json:"confirms"`
	UpdateTime int `json:"update_time"`
	Status     int `json:"status"`
}

//添加存币记录
func InputSave(data map[string]interface{}, maxConfirm int) (insertId int) {

	defer Db.Close()
	Db = Connect()

	ctime := int(time.Now().Unix())
	confirms := data["confirms"].(int)
	status := InputStatus_Pending
	if maxConfirm > 0 && confirms >= maxConfirm {
		status = InputStatus_Complete
	}
	info := &CoinInput{
		0,
		data["userId"].(int),
		data["coinId"].(int),
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
func UpdateInputCfm(id, confirms, maxConfirm int) (result bool) {

	defer Db.Close()
	Db = Connect()

	utime := int(time.Now().Unix())
	status := InputStatus_Pending
	if maxConfirm > 0 && confirms >= maxConfirm {
		status = InputStatus_Complete
	}
	data := &CoinInputCfm{
		confirms,
		utime,
		status,
	}
	err := Db.Hander.Table("coin_inputs").Where("id = ?", id).Updates(data).Error
	if err == nil {
		result = true
	}
	return
}

//判断地址是否存在
func GetInputExist(coinId int, txId string) bool {

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
func GetPendingInputs() (result []CoinInput) {

	defer Db.Close()
	Db = Connect()

	Db.Hander.Table("coin_inputs").Where("status = 1").Find(&result)
	return
}
