#!/usr/bin/env bash

if [[ $TRAVIS_OS_NAME == 'osx' ]]; then 
  REPORTER_DL_LINK="https://codeclimate.com/downloads/test-reporter/test-reporter-latest-darwin-amd64"
else
  REPORTER_DL_LINK="https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64"
fi

curl -L $REPORTER_DL_LINK > ./cc-test-reporter
chmod +x ./cc-test-reporter