package models

import (
	"time"
)

type CoinAddress struct {
	Id         uint   `gorm:"primary_key" json:"id"`
	UserId     uint   `json:"user_id"`
	CoinId     uint   `json:"coin_id"`
	Address    string `json:"address"`
	Account    string `json:"account"`
	Balance    string `json:"balance"`
	CreateTime uint   `json:"create_time"`
	UpdateTime uint   `json:"update_time"`
}

//添加钱包地址
func AddressSave(data map[string]interface{}) (insertId uint) {

	defer Db.Close()
	Db = Connect()
	time := uint(time.Now().Unix())
	info := &CoinAddress{
		0,
		data["userId"].(uint),
		data["coinId"].(uint),
		data["address"].(string),
		data["account"].(string),
		"0",
		time,
		time,
	}
	obj := Db.Hander.Table("coin_address").Create(info).Value
	insertId = uint(obj.(*CoinAddress).Id)

	return
}

//判断地址是否存在
func GetAddressExist(coinId uint, address string) bool {

	defer Db.Close()
	Db = Connect()
	var count int
	obj := Db.Hander.Table("coin_address").Where("coin_id = ? and address = ?", coinId, address)
	obj.Count(&count)

	if count > 0 {
		return true
	} else {
		return false
	}
}

//查询
func GetAddress(coinId uint, address string) (result CoinAddress) {

	defer Db.Close()
	Db = Connect()
	Db.Hander.Table("coin_address").
		Where("coin_id = ? and address = ?", coinId, address).
		First(&result)
	return
}
