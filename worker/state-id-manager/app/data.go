package app

import (
	"math/big"
)

// LastStateID - Holds last state what was synced into child chain
// along with timestamp of last two updates
type LastStateID struct {
	ID *big.Int
}
