package tracker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// Checks whether a ERC20 token approval has completed or not, given transaction hash on root chain
//
// Due to the fact, token approval must be performed before calling `depositFor` on root chain contract,
// only then root contract can call `transferFrom` on token being deposited
func getApprovalStatus(client *ethclient.Client, db *gorm.DB, txHash common.Hash) *TransactionState {
	// If record is present in database & confirmed, then we're simply going to read status from database
	// and return to client
	if status := getRootChaintxStatusFromDB(db, txHash); status != nil {
		if status.Code == 5 || status.Code == 6 {
			return &TransactionState{
				Code:    status.Code,
				Message: status.Message,
			}
		}
	}

	receipt := getTransactionReceipt(client, txHash)
	if receipt == nil {
		return &TransactionState{
			Code:    7,
			Message: "Pending",
		}
	}

	if receipt.Status == 0 {
		putRootChainTxStatusInDB(db, txHash, 6, "Failed")

		return &TransactionState{
			Code:    6,
			Message: "Failed",
		}
	}

	putRootChainTxStatusInDB(db, txHash, 5, "Approved")

	return &TransactionState{
		Code:    5,
		Message: "Approved",
	}
}
