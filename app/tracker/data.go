package tracker

import (
	"encoding/json"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

// BulkPayload - An array of tx hashes can be sent with request body, which
// will eventually return JSON response where keys will be txHashes & values
// will be their respective status
type BulkPayload struct {
	TransactionHashes []common.Hash `json:"txHashes" binding:"required"`
}

// Given txHash, checks whether that's present in this slice or not
func exists(buffer []common.Hash, txHash common.Hash) bool {
	for _, v := range buffer {
		if v == txHash {
			return true
		}
	}

	return false
}

// Returns a slice of non-duplicate txHashes
//
// This function is implemented so that we don't end up
// wasting resources in computing current status of tx ( using txHash )
// if user tries to fool service, by sending a JSON array of same txHashes
// for `n` times
func (b *BulkPayload) unique() []common.Hash {
	buffer := make([]common.Hash, 0)

	for _, v := range b.TransactionHashes {
		if !exists(buffer, v) {
			buffer = append(buffer, v)
		}
	}

	return buffer
}

// PlasmaExitBulkPayload - JSON encoded payload to be feeded to /v1/plasma-exit endpoint
// to check status of plasma withdraw
type PlasmaExitBulkPayload struct {
	TransactionHashes []CheckExitable `json:"txHashes" binding:"required"`
}

// POSExited - Payload to be sent when querying `pos-exit-checker`
// for checking exit status of transaction
type POSExited struct {
	TransactionHash string `json:"txHash"`
}

// JSON - Converts to JSON encoded byte array
func (p *POSExited) JSON() []byte {
	data, err := json.Marshal(p)
	if err != nil {
		log.Println("[!] Failed to marshal `POSExited` to JSON")
		return nil
	}

	return data
}

// TransactionState - Represents current state of any transaction, though note
// that transaction hash is not being kept inside this structure
type TransactionState struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// JSON - Converts to JSON encoded byte array
func (t *TransactionState) JSON() []byte {
	data, err := json.Marshal(t)
	if err != nil {
		log.Println("[!] Failed to marshal `TransactionState` to JSON")
		return nil
	}

	return data
}

// LastStateID - Holds last state what was synced, to child chain
// to be queried by talking to `state-id-manager` micro service
type LastStateID struct {
	ID string `json:"id"`
}

// Given json encoded byte data, unmarshalling it to LastStateID struct
func decodeToLastStateID(data []byte) *LastStateID {
	var lastStateID LastStateID

	err := json.Unmarshal(data, &lastStateID)
	if err != nil {
		return nil
	}

	return &lastStateID
}

// CheckPointed - Data to be sent in POST request, before performing
// a check on whether this child chain block has been check pointed or not
type CheckPointed struct {
	BlockNumber string `json:"blockNumber"`
}

// JSON - JSON encoded form, to be sent as HTTP request body
func (c *CheckPointed) JSON() []byte {
	data, err := json.Marshal(c)
	if err != nil {
		return nil
	}

	return data
}

// CheckExitable - Data to be sent in POST request to check
// whether this withdraw is still under challenge period or not
//
// Also `/v1/plasma-exit` will accept array of this type, for querying status of
// confirm withdraw tx on root chain
type CheckExitable struct {
	BurnTxHash    common.Hash `json:"burnTxHash" binding:"required"`
	ConfirmTxHash common.Hash `json:"confirmTxHash" binding:"required"`
}

// JSON - JSON encoded form, to be sent as HTTP request body
func (c *CheckExitable) JSON() []byte {
	data, err := json.Marshal(c)
	if err != nil {
		return nil
	}

	return data
}

// WithdrawTransaction - Data schema for withdraw transaction tracking request,
// to be received in this form
//
// When `isPoS` is true, it'll not have `ConfirmWithdrawTxHash` field
// else, it'll have all fields, which is why `ConfirmWithdrawTxHash` field is not
// strictly bound
type WithdrawTransaction struct {
	BurnTxHash            common.Hash `json:"txHash" binding:"required"`
	IsPOS                 bool        `json:"isPoS"`
	ConfirmWithdrawTxHash common.Hash `json:"relatedTxHash"`
	ExitTxHash            common.Hash `json:"exitTxHash"`
}

// WithdrawTransactions - All withdraw transactions required to be tracked
// are to be sent in this form
type WithdrawTransactions struct {
	Transactions []*WithdrawTransaction `json:"withdrawTxObjectArray" binding:"required"`
}

// WithdrawTransactionStatus - Reponse of withdraw tx status tracking request
type WithdrawTransactionStatus struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	IsPOS   bool   `json:"isPoS"`
}
