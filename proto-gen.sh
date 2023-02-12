#!/usr/bin/bash

if [ "$1" = "go" ]; then
    protoc \
        -I=. \
        --go_out=. \
        --go-grpc_out=. \
        protos/*.proto
elif [ "$1" = "node" ]; then
    protoc \
        -I=. \
        --js_out=import_style=commonjs:./web \
        --grpc-web_out=import_style=typescript,mode=grpcwebtext:./web \
        protos/*.proto
else
    echo "Unknown argument: '$1'"
fi
