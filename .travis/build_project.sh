#!/usr/bin/env bash



if [[ $TRAVIS_OS_NAME == 'osx' ]]; then 
    .travis/macos_build.sh $1 $2
else
    cd cmd/lipovision
    go build -ldflags '-s -w' -o ../../build/lipovision -tags gtk_3_10
fi