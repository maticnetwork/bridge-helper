# Check Point Tracker

## Introduction

This micro service will keep listening for occurance of check point on root chain and update its internal data structure. It'll also expose two HTTP endpoints

- For querying current checkpoint status
- For checking whether certain child chain `blockNumber` has been checkpointed or not

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
RPC=wss://root.node
RootChain=2890bA17EfE978480615e330ecB65333b880928e
PORT=7002
```

## Building

Compile to executable binary

```bash
go build -o check-point-tracker
```

## Running

```bash
./check-point-tracker
```

## Endpoints

Name | Payload | Response | Type | Info
--- | --- | --- | --- | --- | ---
`/` | - | `{"start": "5693322", "end": "5693577"}`| GET | Returns latest check pointed block number range, child blocks
`/` | `{"blockNumber": "5693323"}` | `{"code": 1, "msg": "Check Pointed"}`| POST | Given child chain `blockNumber` can return whether it has been check pointed or not
