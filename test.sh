#!/usr/bin/env bash

echo "mode: set" > coverage.out

for d in $(go list ./... | grep -v vendor); do
    go test -coverprofile=profile.out $d
    if [ -f profile.out ]; then
        cat profile.out | grep -v "mode: set" >> coverage.out
        rm profile.out
    fi
done
