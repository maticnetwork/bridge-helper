package tracker

import (
	"app/nft"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// Checking status of `WithdrawManager.processExits(...)` call on root chain. First checked in database, if
// found in confirmed state, we're not talking to blockchain, rather cached status is returned
//
// If tx on root chain is still in pending state, then we'll check with blockchain whether status has updated or not
func getPlasmaExitStatus(client *ethclient.Client, db *gorm.DB, txHash common.Hash) *TransactionState {
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

// Checking status of `WithdrawManager.processExits(...)` call on root chain. First checked in database, if
// found in confirmed state, we're not talking to blockchain, rather cached status is returned
//
// If tx on root chain is still in pending state, then we'll check with blockchain whether status has updated or not
//
// After that we'll also check whether respective NFT minted during confirm tx was performed
// does exist anymore or not
//
// If it does, we need to notify client with status code `-13` that their `processExit` call
// wasn't successful
//
// @note This function is nothing but updated & improved version of `getPlasmaExitStatus`
// so that we also take NFT existance under consideration
func getReliablePlasmaExitStatus(client *ethclient.Client, db *gorm.DB, burnTxHash common.Hash, confirmTxHash common.Hash, _nft *nft.Nft, exitTxHash common.Hash) *TransactionState {
	// If record is present in database & confirmed, then we're simply going to read status from database
	// and return to client
	if status := getRootChaintxStatusFromDB(db, exitTxHash); status != nil {
		if status.Code == -10 || status.Code == -11 {
			return &TransactionState{
				Code:    status.Code,
				Message: status.Message,
			}
		}
	}

	receipt := getTransactionReceipt(client, exitTxHash)
	if receipt == nil {
		return &TransactionState{
			Code:    -12,
			Message: "Pending",
		}
	}

	if receipt.Status == 0 {
		putRootChainTxStatusInDB(db, exitTxHash, -11, "Failed")

		return &TransactionState{
			Code:    -11,
			Message: "Failed",
		}
	}

	// We'll reach here only if Plasma exit tx hash was correct ( exists in network )
	//
	// Now we'll try to see if NFT minted during Plasma confirm tx, still exists or not
	//
	// Ideally it must not exists. But if it does, it denotes user called `processExit`
	// but didn't exit his/ her locked tokens
	//
	// @note This situation happens when a lot of pending tx(s) present in Plasma
	// exit queue, gas provided with tx, gets exhausted
	confirmTxStat := getPlasmaConfirmStatus(client, db, burnTxHash, confirmTxHash, _nft)

	var retStatus *TransactionState

	switch confirmTxStat.Code {

	case -8:
	case -9:
		// Some times, we might reach here
		// when Plasma exit didn't happen for user
		//
		// Following `code`, denotes that
		// scenario, client application is expected to
		// take necessary steps when they see this status code
		//
		// Necessary steps should be asking user again call
		// Plasma `processExit`

		retStatus = &TransactionState{
			Code:    -13,
			Message: "Plasma exit called, but not exited",
		}

	case -10:
		// This is what we expect to see ideally

		putRootChainTxStatusInDB(db, exitTxHash, -10, "Exited")

		retStatus = confirmTxStat

	default:

		// Reaching here is not expected under normal circumstances
		//
		// Kept to handle issues encountered during `getPlasmaConfirmStatus`
		// function call
		retStatus = &TransactionState{
			Code:    -12,
			Message: "Pending",
		}

	}

	return retStatus

}
