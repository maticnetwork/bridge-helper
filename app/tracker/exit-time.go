package tracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

// Given child chain's burnTxHash & respective confirmTxHash, performed on root chain
// it can check whether withdraw tx has covered challenge period or not
func checkWhetherExitable(burnTxHash common.Hash, confirmTxHash common.Hash) *TransactionState {
	resp, err := http.Post(fmt.Sprintf("%s/%s", get("POSExitChecker"), "exit-time"),
		"application/json",
		bytes.NewReader((&CheckExitable{
			BurnTxHash:    burnTxHash,
			ConfirmTxHash: confirmTxHash,
		}).JSON()))

	if err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -8,
			Message: "Exitable in 0",
		}
	}

	defer resp.Body.Close()

	// HTTP status code must be 200 for valid response, otherwise we don't proceed
	if resp.StatusCode != 200 {

		return &TransactionState{
			Code:    -8,
			Message: "Exitable in 0",
		}

	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -8,
			Message: "Exitable in 0",
		}
	}

	var _tmp TransactionState

	if err = json.Unmarshal(data, &_tmp); err != nil {
		log.Println("[!] ", err)

		return &TransactionState{
			Code:    -8,
			Message: "Exitable in 0",
		}
	}

	if _tmp.Code == 0 {
		return &TransactionState{
			Code: -8,
			// _tmp.Message is unix timestamp in seconds, after that this endpoint can be
			// called & it'll see -9 status code
			Message: fmt.Sprintf("Exitable in %s", _tmp.Message),
		}
	}

	return &TransactionState{
		Code:    -9,
		Message: "Ready To Exit",
	}
}
