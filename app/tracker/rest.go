package tracker

import (
	"app/nft"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Minimum & Maximum payload size i.e. these many tx hash can be sent a time & asked to
// find out their status in POST call
//
// Being read from .env file
func getPayloadSize() (int, int) {
	min := get("MinPayloadSize")

	_min, err := strconv.ParseUint(min, 10, 32)
	if err != nil {
		return 1, 10
	}

	max := get("MaxPayloadSize")

	_max, err := strconv.ParseUint(max, 10, 32)
	if err != nil {
		return 1, 10
	}

	return int(_min), int(_max)
}

// Run - Deposit & withdraw flow tracker service's main power horse
// which exposes some REST API(s) to be used for tracking whole deposit & withdraw life cycle
// for both Plasma & PoS bridges
func Run(file string) {
	err := read(file)
	if err != nil {
		log.Fatalln("[!] ", err)
	}

	// reading payload size specified in .env file
	min, max := getPayloadSize()

	rootClient, err := getClient(true)
	if err != nil {
		log.Fatalln("[!] ", err)
	}
	childClient, err := getClient(false)
	if err != nil {
		log.Fatalln("[!] ", err)
	}
	_nft, err := nft.NewNft(common.HexToAddress(get("ExitNFT")), rootClient)
	if err != nil {
		log.Fatalln("[!] ", err)
	}

	// connecting to database, to be used for persisting status of transactions ( using txHash )
	db := connectToDB()
	// Performing auto migration
	migrateDB(db)

	router := gin.Default()

	// Allowing requests from all origins
	router.Use(cors.Default())

	v1 := router.Group("/v1")

	{

		// Given a non-empty set of approval tx hashes, returns status for each of them
		v1.POST("/approval", func(c *gin.Context) {
			var bulkPayload BulkPayload

			if err := c.ShouldBindJSON(&bulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(bulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(bulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := make(map[common.Hash]*TransactionState)

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, h := range bulkPayload.unique() {

				wg.Add(1)
				go func(h common.Hash) {

					_tmp := getApprovalStatus(rootClient, db, h)

					mutex.Lock()
					_statuses[h] = _tmp
					mutex.Unlock()

					wg.Done()
				}(h)

			}
			wg.Wait()

			c.JSON(200, gin.H{
				"approvalTxStatus": _statuses,
				"action":           checkIfApprovalTxInProgress(_statuses),
				"count":            countOfPendingApprovalTx(_statuses),
			})
		})

		// Given a non-empty set of `depositFor`/ `depositEtherFor` tx hashes ( on root chain )
		// it can respond with their current statuses
		v1.POST("/deposit", func(c *gin.Context) {
			var bulkPayload BulkPayload

			if err := c.ShouldBindJSON(&bulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(bulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(bulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := make(map[common.Hash]*TransactionState)

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, h := range bulkPayload.unique() {

				wg.Add(1)
				go func(h common.Hash) {

					_tmp := getDepositStatus(rootClient, db, h)

					mutex.Lock()
					_statuses[h] = _tmp
					mutex.Unlock()

					wg.Done()
				}(h)

			}
			wg.Wait()

			c.JSON(200, gin.H{
				"depositTxStatus": _statuses,
				"action":          checkIfDepositTxInProgress(_statuses),
				"count":           countOfPendingDepositTx(_statuses),
			})
		})

		// Given a non-empty set of burn tx hashes on child chain, it can track
		// all of their respective status & returns same
		//
		// Note : Please stop using /v1/pos-withdraw for tracking pos withdraw tx status
		// to be replaced by this one, in near future
		v1.POST("/pos-burn", func(c *gin.Context) {
			var bulkPayload BulkPayload

			if err := c.ShouldBindJSON(&bulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(bulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(bulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := gin.H{}

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, h := range bulkPayload.unique() {

				wg.Add(1)
				go func(h common.Hash) {

					_tmp := getPOSBurnStatus(childClient, db, h)

					mutex.Lock()
					_statuses[h.Hex()] = _tmp
					mutex.Unlock()

					wg.Done()
				}(h)

			}
			wg.Wait()

			c.JSON(200, _statuses)
		})

		// Given a non-empty set of burn tx hashes on child chain, it can track
		// all of their respective status & returns same
		//
		// @todo To be removed in near future, please consider using `/v1/pos-burn` instead of this one
		v1.POST("/pos-withdraw", func(c *gin.Context) {
			var bulkPayload BulkPayload

			if err := c.ShouldBindJSON(&bulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(bulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(bulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := gin.H{}

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, h := range bulkPayload.unique() {

				wg.Add(1)
				go func(h common.Hash) {

					_tmp := getPOSBurnStatus(childClient, db, h)

					mutex.Lock()
					_statuses[h.Hex()] = _tmp
					mutex.Unlock()

					wg.Done()
				}(h)

			}
			wg.Wait()

			c.JSON(200, _statuses)
		})

		// Given a non-empty set of exit tx hashes on root chain, it can check their status
		//
		// Note: Consider using this endpoint for checking exit tx status on root chain,
		// /v1/exit to be removed in near future
		v1.POST("/pos-exit", func(c *gin.Context) {
			var bulkPayload BulkPayload

			if err := c.ShouldBindJSON(&bulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(bulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(bulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := gin.H{}

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, h := range bulkPayload.unique() {

				wg.Add(1)
				go func(h common.Hash) {

					_tmp := getPOSExitStatus(rootClient, db, h)

					mutex.Lock()
					_statuses[h.Hex()] = _tmp
					mutex.Unlock()

					wg.Done()
				}(h)

			}
			wg.Wait()

			c.JSON(200, _statuses)
		})

		// Given a non-empty set of exit tx hashes on root chain, it can check their status
		//
		// @todo To be removed in near future, please consider using `/v1/pos-exit` instead of this one
		v1.POST("/exit", func(c *gin.Context) {
			var bulkPayload BulkPayload

			if err := c.ShouldBindJSON(&bulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(bulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(bulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := gin.H{}

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, h := range bulkPayload.unique() {

				wg.Add(1)
				go func(h common.Hash) {

					_tmp := getPOSExitStatus(rootClient, db, h)

					mutex.Lock()
					_statuses[h.Hex()] = _tmp
					mutex.Unlock()

					wg.Done()
				}(h)

			}
			wg.Wait()

			c.JSON(200, _statuses)
		})

		// Given token burn tx hash on child chain, it can track it's state
		// upto checkpointing state.
		//
		// Once status code is equivalent to checkpointed, client is good
		// to go for calling `ERC20Predicate.startExitWithBurntTokens(...)`
		//
		// Next step to be tracked using `/v1/plasma-confirm` endpoint
		v1.POST("/plasma-burn", func(c *gin.Context) {
			var bulkPayload BulkPayload

			if err := c.ShouldBindJSON(&bulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(bulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(bulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := gin.H{}

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, h := range bulkPayload.unique() {

				wg.Add(1)
				go func(h common.Hash) {

					_tmp := getCheckPointStatus(childClient, db, h)

					mutex.Lock()
					_statuses[h.Hex()] = _tmp
					mutex.Unlock()

					wg.Done()
				}(h)

			}
			wg.Wait()

			c.JSON(200, _statuses)
		})

		// Given a non-empty array of child chain burn tx hashes & root chain
		// confirm withdraw tx hashes, it can check status of plasma withdraw
		// upto whether withdraw fully completed or not
		v1.POST("/plasma-confirm", func(c *gin.Context) {
			var plasmaExitBulkPayload PlasmaExitBulkPayload

			if err := c.ShouldBindJSON(&plasmaExitBulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(plasmaExitBulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(plasmaExitBulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := gin.H{}

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, v := range plasmaExitBulkPayload.TransactionHashes {

				wg.Add(1)
				go func(v CheckExitable) {

					_tmp := getPlasmaConfirmStatus(rootClient, db, v.BurnTxHash, v.ConfirmTxHash, _nft)

					mutex.Lock()
					_statuses[v.ConfirmTxHash.Hex()] = _tmp
					mutex.Unlock()

					wg.Done()
				}(v)

			}
			wg.Wait()

			c.JSON(200, _statuses)
		})

		// Given root chain tx hash, obtained after performing `WithdrawManager.processExits(...)`
		// it'll check their status & return so
		v1.POST("/plasma-exit", func(c *gin.Context) {
			var bulkPayload BulkPayload

			if err := c.ShouldBindJSON(&bulkPayload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			// Expecting at least 1 txHash
			if len(bulkPayload.TransactionHashes) < min {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			// If more than 10 tx hashes are asked to be tracked
			// we're simply going to not take this request up
			if len(bulkPayload.TransactionHashes) > max {
				c.JSON(400, gin.H{
					"msg": "Heavy Payload",
				})
				return
			}

			_statuses := gin.H{}

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			// Starting `n (<=10)` workers concurrently, who are going to find out status
			// of a specific txHash, wait until all workers complete
			// and then return json response
			for _, h := range bulkPayload.unique() {

				wg.Add(1)
				go func(h common.Hash) {

					_tmp := getPlasmaExitStatus(rootClient, db, h)

					mutex.Lock()
					_statuses[h.Hex()] = _tmp
					mutex.Unlock()

					wg.Done()
				}(h)

			}
			wg.Wait()

			c.JSON(200, _statuses)
		})

	}

	v2 := router.Group("/v2")

	{

		v2.POST("/withdraw", func(c *gin.Context) {

			var payload WithdrawTransactions

			if err := c.ShouldBindJSON(&payload); err != nil {
				c.JSON(400, gin.H{
					"msg": "Bad Payload",
				})
				return
			}

			if !(len(payload.Transactions) > 0) {
				c.JSON(400, gin.H{
					"msg": "Empty Payload",
				})
				return
			}

			_statuses := make(map[common.Hash]*WithdrawTransactionStatus)

			mutex := sync.Mutex{}
			var wg sync.WaitGroup

			for _, v := range payload.Transactions {

				wg.Add(1)
				go func(tx *WithdrawTransaction) {

					// --- Function to be used for storing withdraw tx status in shared map
					storeWithdrawTxStatus := func(state *TransactionState) {

						mutex.Lock()
						defer mutex.Unlock()

						// Critical section code
						_statuses[tx.BurnTxHash] = &WithdrawTransactionStatus{
							Code:    state.Code,
							Message: state.Message,
							IsPOS:   tx.IsPOS,
						}

					}
					// -- This function deals with critical section of code
					// which is why its updations are protected using Mutex

					// Scheduling worker done call from now
					//
					// to be invoked when returning from this function call stack
					defer wg.Done()

					// burn hash must be supplied
					if isEmptyTxHash(tx.BurnTxHash) {
						return
					}

					switch tx.IsPOS {
					case true:

						// If POS exit hash is available, check status using that hash
						if !isEmptyTxHash(tx.ExitTxHash) {
							storeWithdrawTxStatus(getPOSExitStatus(rootClient, db, tx.ExitTxHash))
							return
						}

						// If only POS burn tx hash is available, try to
						// check status using that tx hash, whether checkpointed or not
						//
						// @note This is being replaced with `getPOSBurnStatus` so that wallet-web
						// client can obtain whether pos-exited or not status, even when no exit
						// hash is provided with, only burn hash is what we have
						//
						// If you have exit tx hash too, please consider sending in payload along
						// with burn tx hash
						//
						// Converting exit denoting status code to `-10` to match both of
						// `/v1/pos-exit` & `/v1/plasma-exit`
						_txStatus := getPOSBurnStatus(childClient, db, tx.BurnTxHash)
						if _txStatus.Code == -5 {
							_txStatus.Code = -10
						}

						storeWithdrawTxStatus(_txStatus)

					case false:

						// If Plasma exit hash is provided with, then check using its status
						if !isEmptyTxHash(tx.ExitTxHash) {
							storeWithdrawTxStatus(getReliablePlasmaExitStatus(rootClient, db, tx.BurnTxHash, tx.ConfirmWithdrawTxHash, _nft, tx.ExitTxHash))
							return
						}

						// If Plasma confirm withdraw tx hash is given, check using burn tx hash & confirm
						// withdraw tx hash
						if !isEmptyTxHash(tx.ConfirmWithdrawTxHash) {
							storeWithdrawTxStatus(getPlasmaConfirmStatus(rootClient, db, tx.BurnTxHash, tx.ConfirmWithdrawTxHash, _nft))
							return
						}

						// If only Plasma burn tx hash is available, try to
						// check status using that, whether checkpointed or not
						storeWithdrawTxStatus(getCheckPointStatus(childClient, db, tx.BurnTxHash))

					}

				}(v)

			}
			wg.Wait()

			c.JSON(200, gin.H{
				"withdrawTxStatus": _statuses,
				"action":           calculateActionRequiredForWithdraw(_statuses),
				"count":            calculateCountOfPendingWithdrawTx(_statuses),
			})

		})

	}

	router.Run(strings.Join([]string{":", get("PORT")}, ""))
}

// Calculating what should be higher priority activity for user
// depending upon computed tx status codes
//
// `Action Required` is higher in priority than `Transaction In Progress`
func calculateActionRequiredForWithdraw(statuses map[common.Hash]*WithdrawTransactionStatus) string {

	// --- Template messages/ labels to be sent in response
	const ActionRequired = "Action Required"
	const TxInProgress = "Transaction In Progress"
	// ---

	var action string

	for _, v := range statuses {

		if v.Code == -4 || v.Code == -9 {
			action = ActionRequired
			break
		}

		if v.Code == -1 || v.Code == -3 || v.Code == -5 || v.Code == -8 || v.Code == -12 {
			action = TxInProgress
		}

	}

	return action

}

// Finds how many withdraw txs are in pending state & returns that count to client
func calculateCountOfPendingWithdrawTx(statuses map[common.Hash]*WithdrawTransactionStatus) int {
	count := 0

	for _, v := range statuses {
		switch v.Code {
		case -1, -3, -4, -5, -8, -9, -12:
			count++
		}
	}

	return count
}

// Checking whether given tx hash is empty or not
//
// To be required while checking txHash provided in payload of `/v2/withdraw` endpoint
//
// @note Case insensitive matching performed, so that we can catch `0[x|X]...`
func isEmptyTxHash(hash common.Hash) bool {

	reg, err := regexp.Compile("(?i)^(0x0{64})$")
	if err != nil {
		return false
	}

	return reg.MatchString(hash.Hex())

}

// Given statuses obtained for approval tx(s), it'll check if atleast one of
// them present in a non-final ( this is a debatable topic ) state or not
//
// If yes, it'll return `Transaction In Progress`
func checkIfApprovalTxInProgress(statuses map[common.Hash]*TransactionState) string {

	var action string

	for _, v := range statuses {

		if v.Code == 7 {
			action = "Transaction In Progress"
			break
		}

	}

	return action

}

// Counting how many approval tx(s) are in pending state
func countOfPendingApprovalTx(statuses map[common.Hash]*TransactionState) int {

	var count int

	for _, v := range statuses {

		if v.Code == 7 {
			count++
		}

	}

	return count

}

// Given statuses obtained for deposit tx(s), it'll check if atleast one of
// them present in a non-final ( this is a debatable topic ) state or not
//
// If yes, it'll return `Transaction In Progress`
func checkIfDepositTxInProgress(statuses map[common.Hash]*TransactionState) string {

	var action string

	for _, v := range statuses {

		if v.Code == 4 || v.Code == 1 {
			action = "Transaction In Progress"
			break
		}

	}

	return action

}

// Counting how many deposit tx(s) are in pending state
func countOfPendingDepositTx(statuses map[common.Hash]*TransactionState) int {

	var count int

	for _, v := range statuses {

		if v.Code == 4 || v.Code == 1 {
			count++
		}

	}

	return count

}
