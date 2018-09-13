#!/usr/bin/env bash

if [[ $TRAVIS_OS_NAME == 'osx' ]]; then 
    ./travis_build_osx.sh
else
  ./travis_build_linux.sh
fi