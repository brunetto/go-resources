#!/usr/bin/env bash 
docker run --rm -t -v ${PWD}:/work -v $(go env GOMODCACHE):/go/pkg/mod -w /work golang go run main.go