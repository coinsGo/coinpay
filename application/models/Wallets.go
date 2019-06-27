package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CoinWallet struct {
	Id         uint    `gorm:"primary_key" json:"id"`
	UserId     uint    `json:"user_id"`
	CoinId     uint    `json:"coin_id"`
	Assets     float64 `json:"assets"`
	Freeze     float64 `json:"freeze"`
	Pended     float64 `json:"pended"`
	CreateTime uint    `json:"create_time"`
	UpdateTime uint    `json:"update_time"`
	Status     uint    `json:"status"`
}

const (
	WalletStatus_Enable  uint = 1 //启用
	WalletStatus_Disable uint = 2 //停用
)

//保存钱包记录
func WalletSave(data map[string]interface{}) (insertId uint) {

	defer Db.Close()
	Db = Connect()

	ctime := uint(time.Now().Unix())
	info := &CoinWallet{
		0,
		data["userId"].(uint),
		data["coinId"].(uint),
		0,
		0,
		0,
		ctime,
		ctime,
		1,
	}
	obj := Db.Hander.Table("coin_wallets").Create(info).Value
	insertId = obj.(*CoinWater).Id
	return
}

//更新钱包资金
func UpdateWallet(user_id, coin_id uint, data map[string]interface{}) (result bool) {

	defer Db.Close()
	Db = Connect()

	info := map[string]interface{}{
		"assets":      gorm.Expr("assets + ?", data["assets"].(float64)),
		"freeze":      gorm.Expr("freeze + ?", data["freeze"].(float64)),
		"pended":      gorm.Expr("pended + ?", data["pended"].(float64)),
		"update_time": uint(time.Now().Unix()),
	}
	err := Db.Hander.Table("coin_wallets").
		Where("user_id = ? and coin_id = ?", user_id, coin_id).
		Updates(info).Error
	if err == nil {
		result = true
	}
	return

}
