# Packs API (WIP)

Packs API written in Go

## Pre-requirements

Go v1.17+

## Update dependencies

```
go mod tidy
```

## Env variables

```
# RPC Endpoint; used to sync previous blocks when is catching up
PACKS_API_RPC_ENDPOINTS=http://rpc1:26657,http://rpc2:26657

# WS Endpoint; used to sync live block as soon as they are available through RPC websocket  
PACKS_API_WS_ENDPOINTS=ws://rpc1:26657/websocket,ws://rpc2:26657/websocket

# Block height to start syncing from
PACKS_API_CURRENT_BLOCK=0

# If true service won't accept any blocks (usually for debugging)
PACKS_API_DISABLE_SYNC=false
```