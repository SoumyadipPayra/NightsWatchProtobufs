#!/bin/bash

# Create directories if they don't exist
mkdir -p gogenproto

# Compile the proto files
protoc -I . \
  --go_out . \
  --go_opt=module=github.com/nightswatchprotobufs \
  --go-grpc_out . \
  --go-grpc_opt=module=github.com/nightswatchprotobufs \
  --grpc-gateway_out . \
  --grpc-gateway_opt=module=github.com/nightswatchprotobufs \
  ./protos/server/server.proto
