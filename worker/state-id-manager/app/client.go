package app

import (
	"state-id-manager/receiver"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Connects to RPC endpoint on child chain
func getClient() (*ethclient.Client, error) {
	return ethclient.Dial(get("RPC"))
}

// Obtains an instance of state receiver contract on child chain
func getStateReceiver(client *ethclient.Client) *receiver.Receiver {
	_receiver, err := receiver.NewReceiver(common.HexToAddress(get("StateReceiver")), client)
	if err != nil {
		return nil
	}

	return _receiver
}
