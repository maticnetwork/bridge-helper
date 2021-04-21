# Bridge API

## Introduction

This micro service exposes some endpoints, which facilitate querying current state of deposit/ withdraw transaction, performed from root chain/ child chain, where only a set of `txHashes` need to be supplied, as POST request payload.

**This is only micro service to be interacted from outside**.

## Prerequisite

- For running this micro service, make sure you've Golang (>=1.13)
- Download all module dependencies by running

```bash
go get
```
- Create one `.env` file in this directory

```bash
touch .env
```
- It must have following fields

```
RootRPC=wss://root.node
ChildRPC=wss://child.node
StateIDManager=http://localhost:7001
CheckPointTracker=http://localhost:7002
POSExitChecker=http://localhost:7003
PORT=7000
ExitNFT=E2Ab047326B38e4DDb6791551e8d593D30E02724
DB_USER=user
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=database
MinPayloadSize=1
MaxPayloadSize=30
```

- PostgreSQL _( >=12.4 )_ also needs to be installed

```bash
sudo apt-get install postgresql postgresql-contrib # for debian based distros
```

- Start PostgreSQL server

```bash
sudo systemctl start postgresql
```

- Set up postgres & create database, as specified in `.env`
- **No need to worry about database migration, it'll be taken care of during application startup**

## Building

Compile to executable binary

```bash
go build -o bridge-api
```

## Running

```bash
./bridge-api
```

## Endpoints

Name | Payload | Response | Type | Info
--- | --- | --- | --- | --- 
`/v1/approval` | `{"txHashes": ["0x....", "0x...."]}` | `{"0x...": {"code": 5, "msg": "Approved"}, "0x...": {"code": 7, "msg": "Pending"}}`| POST | Given a non-empty array of `Token.approve()`'s txHashes _( on root chain )_, it can respond with their current status
`/v1/deposit` | `{"txHashes": ["0x....", "0x...."]}` | `{"0x...": {"code": 0, "msg": "Deposited"}, "0x...": {"code": 1, "msg": "En Route"}}`| POST | Given a non-empty array of `depositFor/depositEtherFor`'s txHashes _( on root chain )_, it can respond with their current status
`/v1/pos-burn` ~~/v1/pos-withdraw~~ | `{"txHashes": ["0x....", "0x...."]}` | `{"0x...": {"code": -5, "msg": "Exited"}, "0x...": {"code": -5, "msg": "Exited"}}`| POST | Given a non-empty array of token burn txHashes _( on child chain )_, it can respond with their current status
`/v1/pos-exit` ~~/v1/exit~~ | `{"txHashes": ["0x....", "0x...."]}` | `{"0x...": {"code": -5, "msg": "Exited"}, "0x...": {"code": -6, "msg": "Failed"}}`| POST | Given a non-empty array of `RootChain*.exit(...)` invokation txHashes _( on root chain )_, it can respond with their current status
`/v1/plasma-burn` | `{"txHashes": ["0x....", "0x...."]}` | `{"0x...": {"code": -4, "msg": "Checkpointed"}, "0x...": {"code": -2, "msg": "Failed"}}`| POST | Given a non-empty array of token burn txHashes _( on child chain )_, it can respond with their current status [ **Upto whether plasma burn tx checkpointed or not** ]
`/v1/plasma-confirm` | `{"txHashes": [{"burnTxHash": "0x....", "confirmTxHash": "0x..."}, {"burnTxHash": "0x....", "confirmTxHash": "0x..."}]}` | `{"0x...": {"code": -10, "msg": "Exited"}, "0x...": {"code": -8, "msg": "Exitable in 0"}}`| POST | Given a non-empty array of token burn tx hash on child chain & confirm withdraw tx hash on root chain, it can return latest status of withdraw [ **Can track after burn tx is checkpointed; upto exit completion** ]
`/v1/plasma-exit` | `{"txHashes": ["0x....", "0x...."]}` | `{"0x...": {"code": -10, "msg": "Exited"}, "0x...": {"code": -12, "msg": "Pending"}}`| POST | Given a non-empty array of `WithdrawManager.processExits(...)` txHashes _( on root chain )_, it can respond with their current status [ **Very similar to `/v1/pos-exit`** ]

> Note : **/v1/pos-withdraw** & **/v1/exit** to be removed in near future

## Deposit Status Codes [ **Plasma & POS** ]

Given that, payload of deposit status checking endpoint(s), is well formatted, we're going to return `http.Ok` with JSON data in body of form

```json
{
    "code": 0,
    "msg": "Deposited"
}
```

Below table demonstrates which status code means what. Consuming client application is supposed to take next step, depending on response `code`

Endpoint | Code | Message | Interpretation
--- | --- | --- | ---
`/v1/approval` | 7 | Pending | Status of `Token.approve(...)` tx on root chain
`/v1/approval` | 6 | Failed | Status of `Token.approve(...)` tx on root chain
`/v1/approval` | 5 | Approved | Status of `Token.approve(...)` tx on root chain
`/v1/deposit` | 4 | Pending | Status of `RootChain*.{depositFor(...), depositEtherFor(...)}` tx on root chain
`/v1/deposit` | 3 | Bad Deposit Hash | Status of `RootChain*.{depositFor(...), depositEtherFor(...)}` tx on root chain
`/v1/deposit` | 2  | Failed | Status of `RootChain*.{depositFor(...), depositEtherFor(...)}` tx on root chain
`/v1/deposit` | 1 | En Route | Status of `RootChain*.{depositFor(...), depositEtherFor(...)}` tx on root chain [ **Going to be synced any moment** ]
`/v1/deposit` | 0 | Deposited | Status of `RootChain*.{depositFor(...), depositEtherFor(...)}` tx on root chain [ **Successful Deposit** ]


## Withdraw Status Codes [ POS ]

Given that, payload of pos-withdraw status checking endpoint, is well formatted, we're going to return `http.Ok` with JSON data in body of form

```json
{
    "code": -10,
    "msg": "Exited"
}
```

Below table demonstrates which status code means what. Consuming client application is supposed to take next step, depending on response `code`

Endpoint | Code | Message | Interpretation
--- | --- | --- | ---
`/v1/pos-burn` ~~/v1/pos-withdraw~~ | -1 | Pending | Token burning tx on child chain is yet to be confirmed
`/v1/pos-burn` ~~/v1/pos-withdraw~~ | -2 | Failed | Token burning tx's execution on child has failed
`/v1/pos-burn` ~~/v1/pos-withdraw~~ | -3 | Burnt | Token burning tx on child chain is successful [ **To be checkpointed** ]
`/v1/pos-burn` ~~/v1/pos-withdraw~~ | -4 | Checkpointed | Token burning tx on child chain has been checkpointed, good to go for `exit(...)`
`/v1/pos-burn` ~~/v1/pos-withdraw~~ | -5 | Exited | Token burning tx on child chain has exited on root chain using POS bridge i.e. `exit(...)` has been called
`/v1/pos-exit` ~~/v1/exit~~ | -12 | Pending | `RootChain*.exit(...)` transaction on root chain, in pending state
`/v1/pos-exit` ~~/v1/exit~~ | -11 | Failed | `RootChain*.exit(...)` transaction execution on root chain failed
`/v1/pos-exit` ~~/v1/exit~~ | -10 | Exited | `RootChain*.exit(...)` transaction execution on root chain, completed

> Note : Always prioritize `/v1/pos-exit`'s response compared to `/v1/pos-burn`, after tx is checkpointed.

## Withdraw Status Codes [ Plasma ]

Given that payload for these endpoints are well formed, it'll respond with `http.Ok` & respond with data of below form.

```json
{
    "code": -10,
    "msg": "Exited"
}
```

Below table demonstrates which status code means what. Consuming client application is supposed to take next step, depending on response `code`

Endpoint | Code | Message | Interpretation
--- | --- | --- | ---
`/v1/plasma-burn` | -1 | Pending | Token burning tx on child chain is yet to be confirmed
`/v1/plasma-burn` | -2 | Failed | Token burning tx's execution on child has failed
`/v1/plasma-burn` | -3 | Burnt | Token burning tx on child chain is successful [ **To be checkpointed** ]
`/v1/plasma-burn` | -4 | Checkpointed | Token burning tx on child chain has been checkpointed, good to go for `ERC20Predicate.startExitWithBurntTokens(...)` on root chain
`/v1/plasma-confirm` | -5 | Pending | Confirm withdraw tx on root chain, still in pending state
`/v1/plasma-confirm` | -6 | Bad Plasma Exit Hash | Feeded confirm withdraw tx hash on root chain, is bad
`/v1/plasma-confirm` | -7 | Failed | Confirm withdraw tx's execution on root chain failed
`/v1/plasma-confirm` | -8 | Exitable in 1603909681 | Plasma withdraw has not yet exceeded challenge period [ **Using `/\s\d+$/g;` regex, challenge period's end timestamp in second can be extracted** ]
`/v1/plasma-confirm` | -9 | Ready to Exit | Plasma withdraw can now be exited on root chain i.e. `WithdrawManager.processExits(...)` can be called
`/v1/plasma-confirm` | -10 | Exited | Plasma withdraw for burn tx & confirm withdraw tx combination has successfully completed [ **Process exit successfully called** ]
`/v1/plasma-exit` | -12 | Pending | `WithdrawManager.processExits(...)` transaction on root chain, in pending state
`/v1/plasma-exit` | -11 | Failed | `WithdrawManager.processExits(...)` transaction execution on root chain failed 
`/v1/plasma-exit` | -10 | Exited | `WithdrawManager.processExits(...)` transaction execution on root chain, completed  [ **Plasma Exit completed** ]

> Note : When -8 is received from `/v1/plasma-exit`, **"Exitable in 0"** can also be returned if timestamp can't be determined

## One endpoint for tracking Withdraw Status

> Please use this REST API for tracking status of all withdraw operations i.e. Plasma & PoS

Method : **POST**

End Point : **/v2/withdraw**

Payload :

```json
{
    "withdrawTxObjectArray": [
        {
            "txHash": "0x...",
            "isPoS": false,
            "relatedTxHash": "0x...",
            "exitTxHash": "0x..."
        },
		{
            "txHash": "0x...",
            "isPoS": true,
            "exitTxHash": "0x..."
        }
    ]
}
```

Response :

```json
{
    "withdrawTxStatus": {
        "0x..." : {
            "code": -3,
            "msg": "Burnt",
            "isPoS": false
        },
        "0x..." : {
            "code": -2,
            "msg": "Failed",
            "isPoS": true
        }
    },
    "action": "Transaction in progress/Action Required", 
    "count": 1
}
```

---

**Important** 👇

> If you ever see, `-13` in Plasma exit status, you need to understand the Plasma confirmTxHash & exitTxHash pair you provided, that doesn't show user's funds got released. Because service attempted to figure out whether NFT minted during Plasma confirm withdraw, still exists or not & it found it does, which was ideally supposed to not exists. 

> You'll arrive in this situation, when you attempted to call `processExit` but Plasma exit queue was lengthy & gas you were willing to spend, got exhausted. So your funds didn't release.

> You'll **not** see this in `/v1/plasma-exit` API

---

> `Action Required` is higher in priority than `Transaction in progress`

> `count` in response in nothing but sum of all tx(s) which haven't reached finality yet.
