package models

import (
	"time"
)

type CoinAddress struct {
	Id         int     `gorm:"primary_key" json:"id"`
	UserId     int     `json:"user_id"`
	CoinId     int     `json:"coin_id"`
	Address    string  `json:"address"`
	Account    string  `json:"account"`
	Balance    float32 `json:"balance"`
	CreateTime int     `json:"create_time"`
	UpdateTime int     `json:"update_time"`
}

//添加钱包地址
func AddressSave(data map[string]interface{}) (insertId int) {

	defer Db.Close()
	Db = Connect()
	time := int(time.Now().Unix())
	info := &CoinAddress{
		0,
		data["userId"].(int),
		data["coinId"].(int),
		data["address"].(string),
		data["account"].(string),
		0,
		time,
		time,
	}
	obj := Db.Hander.Table("coin_address").Create(info).Value
	insertId = obj.(*CoinAddress).Id

	return
}

//判断地址是否存在
func GetAddressExist(coinId int, address string) bool {

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
