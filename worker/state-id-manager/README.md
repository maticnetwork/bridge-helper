# State ID Manager

## Introduction

This micro service has only one responsibility which is querying child chain's StateReceiver contract, for `lastStateId` value & deliver it when some one sends GET request at `/`


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
RPC=wss://child.node
StateReceiver=0000000000000000000000000000000000001001
PORT=7001
```

> Note : Please use websocket endpoint as value of **RPC**

## Building

Compile to executable binary

```bash
go build -o state-id-manager
```

## Running

```bash
./state-id-manager
```

## Endpoints

Name | Payload | Response | Type | Info
--- | --- | --- | --- | --- | ---
`/` | - | `{"id": "2500"}`| GET | Provides us with latest value of `lastStateId`, to be used for checking whether a certain root chain transaction has been synced in or not
