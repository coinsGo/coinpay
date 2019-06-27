package models

type CoinConfig struct {
	Id         uint   `gorm:"primary_key" json:"id"`
	Sign       string `json:"sign"`
	Name       string `json:"name"`
	Confirms   uint   `json:"confirms"`
	CreateTime uint   `json:"create_time"`
	UpdateTime uint   `json:"update_time"`
	Status     uint   `json:"status"`
}

//配置确认数
func GetConfirms() map[uint]uint64 {

	defer Db.Close()
	Db = Connect()

	result := make(map[uint]uint64)

	var list []CoinConfig
	Db.Hander.Table("coin_config").Find(&list)
	for _, row := range list {
		result[row.Id] = uint64(row.Confirms)
	}
	return result
}
