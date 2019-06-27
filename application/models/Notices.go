package models

import (
	"time"
)

const (
	NoticeType_Input  uint = 1 //存币
	NoticeType_Output uint = 2 //提币
)

const (
	NoticeStatus_Waiting uint = 1 //待回调
	NoticeStatus_Success uint = 2 //成功
	NoticeStatus_Failure uint = 3 //失败
)

type CoinNotice struct {
	Id         uint   `gorm:"primary_key" json:"id"`
	UserId     uint   `json:"user_id"`
	TypeId     uint   `json:"type_id"`
	Assid      uint   `json:"assid"`
	Stamp      uint   `json:"stamp"`
	Times      uint   `json:"times"`
	Backurl    string `json:"backurl"`
	Params     string `json:"params"`
	Response   string `json:"response"`
	CreateTime uint   `json:"create_time"`
	UpdateTime uint   `json:"update_time"`
	Status     uint   `json:"status"`
}

//添加回调记录
func NoticeSave(data map[string]interface{}) (insertId uint) {

	defer Db.Close()
	Db = Connect()
	ctime := uint(time.Now().Unix())
	info := &CoinNotice{
		0,
		data["userId"].(uint),
		data["typeId"].(uint),
		data["assid"].(uint),
		ctime,
		0,
		"",
		"",
		"",
		ctime,
		ctime,
		NoticeStatus_Waiting,
	}
	obj := Db.Hander.Table("coin_notices").Create(info).Value
	insertId = uint(obj.(*CoinNotice).Id)

	return
}

//判断记录是否存在
func GetNoticeExist(typeId, assid uint) bool {

	defer Db.Close()
	Db = Connect()
	var count int
	obj := Db.Hander.Table("coin_notices").Where("type = ? and assid = ? and status = 1", typeId, assid)
	obj.Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//查询单条记录
func GetNotice(coinId int, address string) (result CoinAddress) {

	defer Db.Close()
	Db = Connect()
	Db.Hander.Table("coin_notices").Where("coin_id = ? and address = ?", coinId, address).First(&result)
	return
}
