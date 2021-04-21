package tracker

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

// Retrieves tx status, performed on root chain, given tx hash ( for deposit/ withdraw op )
func getRootChaintxStatusFromDB(db *gorm.DB, txHash common.Hash) *RootChain {
	var rootChainTx RootChain

	if err := db.Model(&RootChain{}).Where("txhash = ?", txHash.Hex()).First(&rootChainTx).Error; err != nil {
		return nil
	}

	return &rootChainTx
}

// Updates tx status, performed on root chain, given tx hash ( for deposit/ withdraw op )
//
// If not present in db, creates entry
func putRootChainTxStatusInDB(db *gorm.DB, txHash common.Hash, code int, msg string) {
	if getRootChaintxStatusFromDB(db, txHash) == nil {

		if err := db.Create(&RootChain{
			TransactionHash: txHash.Hex(),
			Code:            code,
			Message:         msg,
		}).Error; err != nil {
			log.Println("[!] ", err)
		}

		return
	}

	if err := db.Model(&RootChain{}).Where("txhash = ?", txHash.Hex()).Select("code", "msg").Updates(&RootChain{
		Code:    code,
		Message: msg,
	}).Error; err != nil {

		log.Println("[!] ", err)

	}
}

// Retrieves tx status, performed on child chain, given tx hash ( for deposit/ withdraw op )
func getChildChaintxStatusFromDB(db *gorm.DB, txHash common.Hash) *ChildChain {
	var childChainTx ChildChain

	if err := db.Model(&ChildChain{}).Where("txhash = ?", txHash.Hex()).First(&childChainTx).Error; err != nil {
		return nil
	}

	return &childChainTx
}

// Updates tx status, performed on child chain, given tx hash ( for deposit/ withdraw op )
//
// If not present in db, creates entry
func putChildChainTxStatusInDB(db *gorm.DB, txHash common.Hash, code int, msg string) {
	if getChildChaintxStatusFromDB(db, txHash) == nil {

		if err := db.Create(&ChildChain{
			TransactionHash: txHash.Hex(),
			Code:            code,
			Message:         msg,
		}).Error; err != nil {
			log.Println("[!] ", err)
		}

		return
	}

	if err := db.Model(&ChildChain{}).Where("txhash = ?", txHash.Hex()).Select("code", "msg").Updates(&ChildChain{
		Code:    code,
		Message: msg,
	}).Error; err != nil {

		log.Println("[!] ", err)

	}
}
