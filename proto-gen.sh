#!/usr/bin/bash

if [ "$1" = "go" ]; then
    protoc \
        -I=. \
        --go_out=. \
        --go-grpc_out=. \
        protos/*.proto
else
    echo "Unknown argument: '$1'"
fi
