package app

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Tracks checkpointing status by listening for `NewHeaderBlock(address,uint256,uint256,uint256,uint256,bytes32)`
// event on rootchain contract
func trackCheckPointing(_storage *CheckPointedBlockRange, _mutex *sync.Mutex) {
	client := getClient()
	if client == nil {
		return
	}

	root := getRootChain(client)
	if root == nil {
		return
	}

	// Closure to be used for attempting to read
	// last checkpointed child chain block number
	getLastCheckPoint := func() {

		// Trying to read from chain, last checkpointed Matic block number
		if lastChildBlock, err := root.GetLastChildBlock(nil); err == nil {

			// -- Critical section of code, acquiring exclusive lock
			_mutex.Lock()
			_storage.End.Set(lastChildBlock)
			_mutex.Unlock()
			// -- ends here, releasing lock

			log.Printf("[+] Fetched last checkpointed block number [ %s ]\n", lastChildBlock.String())
		} else {

			log.Printf("[!] Failed to fetch last checkpointed block number : %s\n", err.Error())

		}

	}

	// Reading for first time, during application initialization
	getLastCheckPoint()
	lastTimeRead := time.Now().UTC()

	// subscribing to listening for specific event emitted by RootChain contract
	// when ever new checkpoint is submitted, this event to be emitted on root chain
	logs := make(chan types.Log)
	subs, err := client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(get("RootChain"))},
		Topics:    [][]common.Hash{{common.HexToHash("0xba5de06d22af2685c6c7765f60067f7d2b08c2d29f53cdf14d67f6d1c9bfb527")}},
	}, logs)
	if err != nil {
		log.Fatalln("[!] ", err)
		return
	}

	// scheduling unsubscription
	defer subs.Unsubscribe()

	for {

		select {
		case err := <-subs.Err():

			// Intentionally crashing program when subscription gets cancelled so that systemd
			// will restart service and we'll again get to subscribe to this event
			//
			// Make sure systemd is configured to restart this service when craashed
			log.Fatalln("[!] ", err)

		case _log := <-logs:

			_parsed, err := root.ParseNewHeaderBlock(_log)
			if err != nil {
				log.Println("[!] ", err)
				continue
			}

			// executing critical section code
			// by acquiring a lock
			_mutex.Lock()
			_storage.Start = _parsed.Start
			_storage.End = _parsed.End
			_mutex.Unlock()

			// updating when last time we received checkpoint info
			// over channel, from RPC node
			lastTimeRead = time.Now().UTC()

			log.Println("[+] Updated Checkpoint info : ", _parsed.Start.String(), " <-> ", _parsed.End.String())

		case <-time.After(time.Minute * time.Duration(30)):

			// If due to some reasons for more than 30 minutes we've not received any checkpoint
			// info from RPC node, we'll attempt to read it from chain directly
			//
			// This ensures every 30 minutes we'll get an opportunity to refresh
			// checkpointed child chain block number
			if time.Now().UTC().Sub(lastTimeRead) >= time.Duration(30)*time.Minute {

				getLastCheckPoint()
				lastTimeRead = time.Now().UTC()

			}

		}

	}

}
