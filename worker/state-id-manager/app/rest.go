package app

import (
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// Run - REST API runner function, exposing only one
// GET endpoint, for obtaining, latest `lastStateId`
// value, which this micro service fetches by talking to
// StateReceiver contract
func Run(file string) {
	err := read(file)
	if err != nil {
		log.Fatalln("[!] ", err)
		return
	}

	stateID := &LastStateID{
		ID: big.NewInt(0),
	}
	mutex := &sync.Mutex{}

	go getLastStateID(stateID, mutex)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		mutex.Lock()
		_id := stateID.ID.String()
		mutex.Unlock()

		c.JSON(200, gin.H{
			"id": _id,
		})
	})

	router.Run(strings.Join([]string{":", get("PORT")}, ""))
}
