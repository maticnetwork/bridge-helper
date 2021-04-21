package app

import "math/big"

// CheckPointedBlockRange - Shared structure for holding recent checkpoint status i.e.
// which child block range is checkpointed in latest phase
type CheckPointedBlockRange struct {
	Start *big.Int
	End   *big.Int
}

// Payload - Data to be sent in POST request, before performing
// a check on whether this child chain block is check pointed or not
type Payload struct {
	BlockNumber string `json:"blockNumber" binding:"required"`
}
