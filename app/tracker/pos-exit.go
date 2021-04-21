package tracker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// Checking status of `RootChain*.exit(...)` call on root chain. First checked in database, if
// found in confirmed state, we're not talking to blockchain, rather cached status is returned
//
// If tx on root chain is still in pending state, then we'll check with blockchain whether status has updated or not
func getPOSExitStatus(client *ethclient.Client, db *gorm.DB, txHash common.Hash) *TransactionState {
	// If record is present in database & confirmed, then we're simply going to read status from database
	// and return to client
	if status := getRootChaintxStatusFromDB(db, txHash); status != nil {
		if status.Code == -10 || status.Code == -11 {
			return &TransactionState{
				Code:    status.Code,
				Message: status.Message,
			}
		}
	}

	receipt := getTransactionReceipt(client, txHash)
	if receipt == nil {
		return &TransactionState{
			Code:    -12,
			Message: "Pending",
		}
	}

	if receipt.Status == 0 {
		putRootChainTxStatusInDB(db, txHash, -11, "Failed")

		return &TransactionState{
			Code:    -11,
			Message: "Failed",
		}
	}

	putRootChainTxStatusInDB(db, txHash, -10, "Exited")

	return &TransactionState{
		Code:    -10,
		Message: "Exited",
	}
}
