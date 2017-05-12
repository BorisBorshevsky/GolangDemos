#!/usr/bin/env bash

set -e
echo "" > coverage.txt
ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --compilers=2

for d in $(go list ./... | grep -v vendor); do
    VAR=$( basename $d )
    if [ -f $GOPATH/src/$d/$VAR.coverprofile ]; then
        cat $GOPATH/src/$d/$VAR.coverprofile >> coverage.txt
        rm $GOPATH/src/$d/$VAR.coverprofile
    fi
done