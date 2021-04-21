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

// Given burn tx hash on child chain, it'll check whether
// this transaction has exited using POS bridge or not
// by talking to another micro service, `pos-exit-checker`
func getPOSBurnStatus(client *ethclient.Client, db *gorm.DB, txHash common.Hash) *TransactionState {
	if status := getChildChaintxStatusFromDB(db, txHash); status != nil {
		if status.Code == -5 || status.Code == -2 {
			return &TransactionState{
				Code:    status.Code,
				Message: status.Message,
			}
		}
	}

	// first checking whether tx is checkpointed or not
	_state := getCheckPointStatus(client, db, txHash)
	if _state.Code != -4 {
		return _state
	}

	resp, err := http.Post(get("POSExitChecker"),
		"application/json",
		bytes.NewReader((&POSExited{
			TransactionHash: txHash.Hex(),
		}).JSON()))
	if err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -4,
			Message: "Checkpointed",
		}
	}

	defer resp.Body.Close()

	var _tmp TransactionState

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -4,
			Message: "Checkpointed",
		}
	}

	err = json.Unmarshal(data, &_tmp)
	if err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -4,
			Message: "Checkpointed",
		}
	}

	if _tmp.Code == 0 {
		return &TransactionState{
			Code:    -4,
			Message: "Checkpointed",
		}
	}

	putChildChainTxStatusInDB(db, txHash, -5, "Exited")

	return &TransactionState{
		Code:    -5,
		Message: "Exited",
	}
}
