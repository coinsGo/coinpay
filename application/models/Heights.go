package models

import "time"

type CoinHeight struct {
	Id         uint   `gorm:"primary_key" json:"id"`
	Sign       string `json:"sign"`
	Name       string `json:"name"`
	Confirms   uint   `json:"confirms"`
	CreateTime uint   `json:"create_time"`
	UpdateTime uint   `json:"update_time"`
	Status     uint   `json:"status"`
}

//配置确认数
func GetHeights() map[uint]uint64 {

	defer Db.Close()
	Db = Connect()

	result := make(map[uint]uint64)

	var list []CoinConfig
	Db.Hander.Table("coin_heights").Find(&list)
	for _, row := range list {
		result[row.Id] = uint64(row.Confirms)
	}
	return result
}

func UpdateHeight(coinId uint, latest, height uint64) (result bool) {

	defer Db.Close()
	Db = Connect()

	info := map[string]interface{}{
		"latest":      latest,
		"height":      height,
		"update_time": uint(time.Now().Unix()),
	}
	err := Db.Hander.Table("coin_heights").Where("coin_id = ?", coinId).Updates(info).Error
	if err == nil {
		result = true
	}
	return
}
