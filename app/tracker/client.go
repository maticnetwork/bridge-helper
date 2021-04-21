package tracker

import "github.com/ethereum/go-ethereum/ethclient"

// Connects to RPC endpoint on root/ child chain
func getClient(isRoot bool) (*ethclient.Client, error) {
	if isRoot {
		return ethclient.Dial(get("RootRPC"))
	}

	return ethclient.Dial(get("ChildRPC"))
}
