# POS Exit Status Checker

## Introduction

Light weight micro service for checking whether POS withdraw has exited on root chain or not, using `matic.js`. Also checks whether plasma confirm tx on root chain is still under challenge period or not i.e. is it okay to call `WithdrawManager.processExits(...)` or not

## Prerequisite

- For running this micro service, make sure you've NodeJS (>=12.\*) & NPM (>=6.*)
- Download all dependencies by running

```bash
npm install
```
- Create one `.env` file in this directory

```bash
touch .env
```
- It must have following fields

```
Network=testnet
Version=mumbai
RootRPC=wss://root.node
ChildRPC=wss://child.node
HOST=127.0.0.1
PORT=7003
```

> ***RPC** : Can be websocket/ http endpoint

## Running

```bash
node index.js
```

## Endpoints

Name | Payload | Response | Type | Info
--- | --- | --- | --- | ---
`/` | `{"txHash": "0x...."}` | `{"code": 1, "msg": "Exited"}`| POST | Given child chain's burn transaction hash, it'll check whether this POS withdraw has exited on root chain or not
`/exit-time` | `{"burnTxHash": "0x....", "confirmTxHash": "0x...."}` | `{"code": 1, "msg": "unix timestamp"}`| POST | Given child chain's burn tx hash & associated confirm tx performed on root chain, it can check whether this plasma withdraw still in challenge period or not
