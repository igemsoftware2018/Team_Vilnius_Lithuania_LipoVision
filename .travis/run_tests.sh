#!/usr/bin/env bash

if [[ $TRAVIS_OS_NAME == 'osx' ]]; then 
    go test  -tags gtk_3_10 -coverprofile c.out ./...
else
    xvfb-run go test  -tags gtk_3_10 -coverprofile c.out ./...
fi