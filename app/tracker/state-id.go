package tracker

import (
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

// Fetches `lastStateId` of child chain contract
// by querying `state-id-manager` contract
func getLastStateID() *big.Int {
	resp, err := http.Get(get("StateIDManager"))
	if err != nil {
		log.Println("[!] ", err)
		return nil
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[!] ", err)
		return nil
	}

	lastStateID := decodeToLastStateID(data)
	if lastStateID == nil {
		log.Println("[!] Failed to decode response from `state-id-manager`")
		return nil
	}

	_tmp := big.NewInt(0)
	_tmp, ok := _tmp.SetString(lastStateID.ID, 10)
	if !ok {
		log.Println("[!] Failed to parse `lastStateId` to big.Int")
		return nil
	}

	return _tmp
}
