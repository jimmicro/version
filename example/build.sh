#!/bin/bash

GIT_TAG=$(git describe --tags --always --dirty="-dev")
BUILD_TIME=$(TZ='Asia/Shanghai' date "+%Y-%m-%d-%H-%M-%S")
LDFLAGS="-s -w  -X 'github.com/jimmicro/version.GitTag=$GIT_TAG' \
    -X 'github.com/jimmicro/version.BuildTime=$BUILD_TIME'"

go build -ldflags="$LDFLAGS" -o main main.go