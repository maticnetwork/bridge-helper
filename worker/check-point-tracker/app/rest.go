package app

import (
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// Run - Exposing REST API for querying what's latest checkpoint status
// and also given child chain block number, it can be checked whether block has
// been checkpointed or not
func Run(file string) {
	err := read(file)
	if err != nil {
		log.Fatalln("[!] ", err)
		return
	}

	checkPointedBlockRange := &CheckPointedBlockRange{
		Start: big.NewInt(0),
		End:   big.NewInt(0),
	}
	mutex := &sync.Mutex{}

	go trackCheckPointing(checkPointedBlockRange, mutex)

	router := gin.Default()

	// GET endpoint for obtaining latest checkpoint's included block range
	router.GET("/", func(c *gin.Context) {
		mutex.Lock()
		start := checkPointedBlockRange.Start.String()
		end := checkPointedBlockRange.End.String()
		mutex.Unlock()

		c.JSON(200, gin.H{
			"start": start,
			"end":   end,
		})
	})

	// When given with blockNumber on child chain, it can respond with `code`
	// to denote whether this child chain block has been checkpointed or not
	router.POST("/", func(c *gin.Context) {
		var payload Payload

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{
				"msg": "Bad Payload",
			})
			return
		}

		_tmp := big.NewInt(0)
		_tmp, ok := _tmp.SetString(payload.BlockNumber, 10)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "Bad Payload",
			})
			return
		}

		// this is critical section code, accessed after
		// acquiring lock
		mutex.Lock()
		if checkPointedBlockRange.End.Cmp(_tmp) >= 0 {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "Check Pointed",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "Not Check Pointed",
			})
		}
		mutex.Unlock()
	})

	router.Run(strings.Join([]string{":", get("PORT")}, ""))
}
