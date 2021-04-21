package tracker

import (
	"app/nft"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// Given both child chain's burn tx hash & respective confirm tx hash on root chain,
// it can check what's status of this plasma exit tx
func getPlasmaConfirmStatus(client *ethclient.Client, db *gorm.DB, burnTxHash common.Hash, confirmTxHash common.Hash, _nft *nft.Nft) *TransactionState {
	if status := getRootChaintxStatusFromDB(db, confirmTxHash); status != nil {
		if status.Code == -6 || status.Code == -7 || status.Code == -10 {
			return &TransactionState{
				Code:    status.Code,
				Message: status.Message,
			}
		}
	}

	receipt := getTransactionReceipt(client, confirmTxHash)
	if receipt == nil {
		return &TransactionState{
			Code:    -5,
			Message: "Pending",
		}
	}

	// Picking out `ExitStarted(address,uint256,address,uint256,bool)` from log entry
	//
	// If present, that will ensure, we're given with correct root chain tx hash, which is generated
	// as result of executing `ERC20Predicate.startExitWithBurntTokens(...)`
	//
	// Otherwise, we're going to stop checking further
	_log := pickOutTransactionLog(receipt.Logs, "0xaa5303fdad123ab5ecaefaf69137bf8632257839546d43a3b3dd148cc2879d6f")
	if _log == nil {
		putRootChainTxStatusInDB(db, confirmTxHash, -6, "Bad Plasma Exit Hash")

		return &TransactionState{
			Code:    -6,
			Message: "Bad Plasma Exit Hash",
		}
	}

	// Tx execution failed
	if receipt.Status == 0 {
		putRootChainTxStatusInDB(db, confirmTxHash, -7, "Failed")

		return &TransactionState{
			Code:    -7,
			Message: "Failed",
		}
	}

	// If Plasma exit has happened, then this NFT
	// must not exist
	exists, err := _nft.Exists(nil, _log.Topics[2].Big())
	if err != nil {

		log.Printf("[!] Failed to check if Plasma withdraw NFT exists : %s\n", err.Error())

		// This scenario was not expected
		// so it's returning some `static` result
		// where it assumes this NFT still
		// exists ( which might not be correct )
		// & says this plasma exit will be eligible
		// for exit in some
		//
		// But when exactly, is not known by service
		return &TransactionState{
			Code:    -8,
			Message: "Exitable in 0",
		}

	}

	// Yes Plasma exit has happened
	if !exists {

		putRootChainTxStatusInDB(db, confirmTxHash, -10, "Exited")

		return &TransactionState{
			Code:    -10,
			Message: "Exited",
		}

	}

	// Attempts to determine how much time left before
	// plasma exit can be invoked
	return checkWhetherExitable(burnTxHash, confirmTxHash)

}
