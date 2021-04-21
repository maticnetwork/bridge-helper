package tracker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// From one given transaction's log entries, we're going to find out
// log entry with first topic `0x103fed9db65eac19c4d870f49ab7520fe03b99f1838e5996caf47e9e43308392`
// i.e. 'StateSynced(uint256,address,bytes)', which will denote its a state sync event emission
// and we can find out stateID from emitted log
func pickOutTransactionLog(logs []*types.Log, topic string) *types.Log {
	for _, v := range logs {
		if v.Topics[0].Hex() == topic {
			return v
		}
	}

	return nil
}

// Given transaction hash, it can return what's current state of deposit transaction
// First approval needs to be performed, so call above function first with `approve` transaction hash
// then call this one with `depositFor`/ `depositEtherFor` transaction hash
func getDepositStatus(client *ethclient.Client, db *gorm.DB, txHash common.Hash) *TransactionState {
	if status := getRootChaintxStatusFromDB(db, txHash); status != nil {
		if status.Code == 0 || status.Code == 2 || status.Code == 3 {
			return &TransactionState{
				Code:    status.Code,
				Message: status.Message,
			}
		}
	}

	receipt := getTransactionReceipt(client, txHash)
	if receipt == nil {
		return &TransactionState{
			Code:    4,
			Message: "Pending",
		}
	}

	// find out that transaction log which has topic `StateSynced(uint256,address,bytes)`, if any
	_log := pickOutTransactionLog(receipt.Logs, "0x103fed9db65eac19c4d870f49ab7520fe03b99f1838e5996caf47e9e43308392")
	if _log == nil {
		putRootChainTxStatusInDB(db, txHash, 3, "Bad Deposit Hash")

		return &TransactionState{
			Code:    3,
			Message: "Bad Deposit Hash",
		}
	}

	// deposit transaction has failed
	if receipt.Status == 0 {
		putRootChainTxStatusInDB(db, txHash, 2, "Failed")

		return &TransactionState{
			Code:    2,
			Message: "Failed",
		}
	}

	lastStateID := getLastStateID()
	// If we're unable to communicate with `state-id-manager` in this moment
	// we're going to assume, fund is on its way to child chain, will reach destination
	//
	// But this assumption may not be correct, fund may have already reached child chain
	if lastStateID == nil {
		return &TransactionState{
			Code:    1,
			Message: "En Route",
		}
	}

	if lastStateID.Cmp(_log.Topics[1].Big()) >= 0 {
		putRootChainTxStatusInDB(db, txHash, 0, "Deposited")

		return &TransactionState{
			Code:    0,
			Message: "Deposited",
		}
	}

	// In this case truly its `en route`
	return &TransactionState{
		Code:    1,
		Message: "En Route",
	}
}
