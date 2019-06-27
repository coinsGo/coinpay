package services

import (
	"log"
	"time"

	"github.com/fanguanghui/coinapi/config"
	"github.com/fanguanghui/coinpay/application/models"
	"github.com/jinzhu/gorm"
)

func InputUpdate(row models.CoinInput, tx config.Tx) error {
	defer models.Db.Close()
	models.Db = models.Connect()
	ORM := models.Db.Hander.Begin()

	ctime := uint(time.Now().Unix())
	status := models.GetInputStatus(row.CoinId, tx.Confirms, tx.Valid)

	//更新存币记录
	info := map[string]interface{}{
		"confirms":    tx.Confirms,
		"update_time": ctime,
		"status":      status,
	}
	if err := ORM.Table("coin_inputs").Where("id = ?", row.Id).Updates(info).Error; err != nil {
		log.Println("更新存币记录", err)
		ORM.Rollback()
		return err
	}

	//通知商户
	if !tx.Valid || tx.Confirms > row.Confirms {
		info := models.CoinNotice{
			UserId:     row.UserId,
			TypeId:     models.NoticeType_Input,
			Assid:      row.Id,
			Stamp:      ctime,
			CreateTime: ctime,
			UpdateTime: ctime,
			Status:     models.NoticeStatus_Waiting,
		}
		if err := ORM.Table("coin_notices").Create(&info).Error; err != nil {
			ORM.Rollback()
			return err
		}
	}

	//记录流水
	if status == models.InputStatus_Failure {
		//待确认资金
		info := models.CoinWater{
			UserId: row.UserId,
			CoinId: row.CoinId,
			TypeId: models.WaterPended_InFail,
			Assid:  row.Id,
			Amount: -row.TxAmount,
			Remark: "存币失败扣减待确认数量",
		}
		if err := ORM.Table("coin_waters").Create(&info).Error; err != nil {
			ORM.Rollback()
			return err
		}

	} else if status == models.InputStatus_Complete {
		//待确认资金
		info := models.CoinWater{
			UserId: row.UserId,
			CoinId: row.CoinId,
			TypeId: models.WaterPended_InSucc,
			Assid:  row.Id,
			Amount: -row.TxAmount,
			Remark: "存币成功扣减待确认数量",
		}
		if err := ORM.Table("coin_waters").Create(&info).Error; err != nil {
			ORM.Rollback()
			return err
		}
		//可用资金
		info = models.CoinWater{
			UserId: row.UserId,
			CoinId: row.CoinId,
			TypeId: models.WaterAssets_Input,
			Assid:  row.Id,
			Amount: row.TxAmount,
			Remark: "存币成功增加可用数量",
		}
		if err := ORM.Table("coin_waters").Create(&info).Error; err != nil {
			ORM.Rollback()
			return err
		}
	}

	//钱包资金
	if status == models.InputStatus_Failure || status == models.InputStatus_Complete {
		assets := 0.0
		if status == models.InputStatus_Complete {
			assets = row.TxAmount
		}
		info := map[string]interface{}{
			"assets":      gorm.Expr("assets + ?", assets),
			"pended":      gorm.Expr("pended - ?", row.TxAmount),
			"update_time": uint(time.Now().Unix()),
		}
		err := ORM.Table("coin_wallets").Where("user_id = ? and coin_id = ?", row.UserId, row.CoinId).Updates(info).Error
		if err != nil {
			ORM.Rollback()
			return err
		}
	}

	return ORM.Commit().Error
}
