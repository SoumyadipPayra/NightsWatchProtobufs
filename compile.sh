#!/bin/bash

# Create directories if they don't exist
mkdir -p gogenproto

# Compile the proto files
protoc -I . \
  --go_out ./gogenproto \
  --go_opt paths=source_relative \
  --go-grpc_out ./gogenproto \
  --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./gogenproto \
  --grpc-gateway_opt paths=source_relative \
  ./server/*.proto
