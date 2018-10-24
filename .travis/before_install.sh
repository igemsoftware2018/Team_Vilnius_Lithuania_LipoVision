#!/usr/bin/env bash

if [[ $TRAVIS_OS_NAME == 'osx' ]]; then 
  brew install s3cmd && brew link s3cmd
  brew install gtk+ && brew link gtk+
  brew install gtk+3 && brew link gtk+3
else
  wget http://netix.dl.sourceforge.net/project/s3tools/s3cmd/2.0.2/s3cmd-2.0.2.tar.gz
  tar xvfz s3cmd-2.0.2.tar.gz
  cd s3cmd-2.0.2
  sudo python setup.py install
  cd ..
fi