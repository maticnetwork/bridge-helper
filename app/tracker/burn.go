package tracker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// Given burn transactionHash on child chain, it can check what's current status
// of this transaction is
//
// This needs to be performed first, before asset can be withdrawn from child
// chain to root chain
func getBurnStatus(client *ethclient.Client, db *gorm.DB, txHash common.Hash) *TransactionState {
	if status := getChildChaintxStatusFromDB(db, txHash); status != nil {
		if status.Code == -3 {
			return &TransactionState{
				Code:    status.Code,
				Message: status.Message,
			}
		}
	}

	receipt := getTransactionReceipt(client, txHash)
	if receipt == nil {
		return &TransactionState{
			Code:    -1,
			Message: "Pending",
		}
	}

	if receipt.Status == 0 {
		putChildChainTxStatusInDB(db, txHash, -2, "Failed")

		return &TransactionState{
			Code:    -2,
			Message: "Failed",
		}
	}

	putChildChainTxStatusInDB(db, txHash, -3, "Burnt")

	return &TransactionState{
		Code:    -3,
		Message: "Burnt",
	}
}
