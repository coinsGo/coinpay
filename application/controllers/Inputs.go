package controllers

import (
	"log"

	"github.com/fanguanghui/coinpay/application/utils"

	"github.com/fanguanghui/coinapi/config"
	"github.com/fanguanghui/coinpay/application/models"
	"github.com/fanguanghui/coinpay/application/services"
)

//实时添加零确认记录
func ChainPending(coinId uint) {

	client := WalletFactory.Get(coinId)
	listTx := client.GetPendingTxs("")
	handleListTx(coinId, listTx)
}

//查询块补齐漏掉的记录
func ChainBlock(coinId uint) {

	client := WalletFactory.Get(coinId)
	height := models.CfgHeights[coinId]
	latest := client.GetBlockCount()
	if height >= latest {
		return
	}
	if height == 0 {
		height = latest
	} else {
		height = height + 1
	}
	listTx := client.GetBlockTxs(height)
	handleListTx(coinId, listTx)

	models.UpdateHeight(coinId, latest, height)
}

//根据TXID查确认数
func ChainTxid(coinId uint) {

	log.Println("根据Txid查询确认数开始")
	client := WalletFactory.Get(coinId)
	listInput := models.GetPendingInputs(coinId)

	for _, row := range listInput {
		tx := client.GetTxById(row.TxId)
		err := services.InputUpdate(row, tx)

		data := map[string]interface{}{
			"filename": "database",
			"size":     10,
		}
		utils.Log.WithFields(data).Info(err)
	}
}

/*
 * ------------------------------------------------------------------------------------------------------------------
 *	                                                     私有函数
 * ------------------------------------------------------------------------------------------------------------------
 */

//处理交易记录列表
func handleListTx(coinId uint, listTx []config.Tx) {
	var info map[string]interface{}
	for _, tx := range listTx {
		for _, vout := range tx.Vouts {
			addr := models.GetAddress(coinId, vout.ToAddress)
			if addr.Id > 0 {
				exist := models.GetInputExist(coinId, tx.Txid)
				if !exist {
					status := models.GetInputStatus(coinId, tx.Confirms, tx.Valid)
					info = map[string]interface{}{
						"userId":      addr.UserId,
						"coinId":      addr.CoinId,
						"txId":        tx.Txid,
						"txAmount":    vout.TxAmount,
						"fromAddress": tx.FromAddress,
						"toAddress":   vout.ToAddress,
						"confirms":    tx.Confirms,
						"status":      status,
					}
					assid := models.InputSave(info)
					if assid > 0 {
						info = map[string]interface{}{
							"userId": addr.UserId,
							"typeId": models.NoticeType_Input,
							"assid":  assid,
						}
						models.NoticeSave(info)
					}
				}
			}
		}
	}
}
