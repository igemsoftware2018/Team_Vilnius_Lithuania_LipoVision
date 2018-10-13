#!/usr/bin/env bash



if [[ $TRAVIS_OS_NAME == 'osx' ]]; then 
    .travis/macos_build.sh
else
    cd cmd/lipovision
    go build -ldflags '-s -w' -o ../../build/lipovision
fi