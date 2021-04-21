package tracker

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Fetches transaction receipt of specific transaction hash
// will only return something non-nil, given that transaction is not pending
func getTransactionReceipt(client *ethclient.Client, txHash common.Hash) *types.Receipt {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return nil
	}

	// Status is 0, in case of failed transaction execution
	// to be taken care of in next stage of processing
	return receipt
}
