package app

import (
	"log"
	"sync"
	"time"
)

// This function is supposed to be run in a diffrent thread of
// execution, which will wake up every 3 minutes & query child chain's
// StateReceiver contract, to get latest `lastStateId` value
//
// This value can be used by checking whether a certain root chain
// deposit transaction has successfully been synced in or not
func getLastStateID(stateID *LastStateID, mutex *sync.Mutex) {
	client, err := getClient()
	if err != nil {
		log.Fatalln("[!] ", err)
		return
	}

	receiver := getStateReceiver(client)
	if receiver == nil {
		log.Fatalln("[!] Failed to get instance of StateReceiver")
		return
	}

	updateStateID := func() {
		id, err := receiver.LastStateId(nil)
		if err != nil {
			log.Println("[!] ", err)
			return
		}

		mutex.Lock()
		if stateID.ID.Cmp(id) == -1 {
			stateID.ID = id
			log.Println("[+] Updated `lastStateId` : ", stateID.ID.String())
		}
		mutex.Unlock()
	}

	for {
		updateStateID()
		time.Sleep(time.Minute * time.Duration(3))
	}

}
