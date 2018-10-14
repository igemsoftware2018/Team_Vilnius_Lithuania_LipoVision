#!/usr/bin/env bash

BUILD_FOLDER=${1:-build}

# Get deps
lines=($(ldd  $BUILD_FOLDER/LipoVision.exe | grep lib | cut -d' ' -f3))

for dep in ${lines[*]}
do
    cp $dep $BUILD_FOLDER
done