#!/usr/bin/env bash

# Launch central-server
go run cmd/central-server/central-server.go &

# Launch node
go run cmd/node/node.go &

# Launch live-cli
go run cmd/live-cli/live-cli.go &

# Wait 5 seconds
sleep 5

# Stop test
killall node live-cli central-server
