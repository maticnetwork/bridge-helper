package tracker

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// Given burn txHash, checks whether it has been checkpointed or not,
// by querying `check-point-manager` micro service
func getCheckPointStatus(client *ethclient.Client, db *gorm.DB, txHash common.Hash) *TransactionState {
	if status := getChildChaintxStatusFromDB(db, txHash); status != nil {
		if status.Code == -4 {
			return &TransactionState{
				Code:    status.Code,
				Message: status.Message,
			}
		}
	}

	// first checking burn status
	// if response code is not equivalent to burnt
	// we're responding with response received from `getBurnStatus`
	_state := getBurnStatus(client, db, txHash)
	if _state.Code != -3 {
		return _state
	}

	// fetching child chain `blockNumber` in which
	// transaction was mined; given txHash, we're fetching
	// respective receipt
	receipt := getTransactionReceipt(client, txHash)
	if receipt == nil {
		return _state
	}

	// sending request to `check-point-tracker` service
	// with `blockNumber` as payload, so that we can check
	// whether this child chain block has been checkpointed or not
	//
	// If yes, we can also say burn tx has been checkpointed
	resp, err := http.Post(get("CheckPointTracker"),
		"application/json",
		bytes.NewReader((&CheckPointed{
			BlockNumber: receipt.BlockNumber.String(),
		}).JSON()))
	if err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -3,
			Message: "Burnt",
		}
	}

	defer resp.Body.Close()

	var _tmp TransactionState

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -3,
			Message: "Burnt",
		}
	}

	err = json.Unmarshal(data, &_tmp)
	if err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -3,
			Message: "Burnt",
		}
	}

	if _tmp.Code == 0 {
		return &TransactionState{
			Code:    -3,
			Message: "Burnt",
		}
	}

	putChildChainTxStatusInDB(db, txHash, -4, "Checkpointed")

	return &TransactionState{
		Code:    -4,
		Message: "Checkpointed",
	}
}
