#!/bin/bash
#GO111MODULE="on" GOPROXY=$GOPROXY go get -u github.com/swaggo/swag/cmd/swag
cd ../backend
swag init --parseDependency --parseInternal