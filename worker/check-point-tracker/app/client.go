package app

import (
	"check-point-tracker/root"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Connects to root blockchain node
func getClient() *ethclient.Client {
	client, err := ethclient.Dial(get("RPC"))
	if err != nil {
		log.Fatalln("[!] ", err)
		return nil
	}

	return client
}

// Obtains an instance of RootChain contract
func getRootChain(client *ethclient.Client) *root.Root {
	_root, err := root.NewRoot(common.HexToAddress(get("RootChain")), client)
	if err != nil {
		log.Fatalln("[!] ", err)
		return nil
	}

	return _root
}
