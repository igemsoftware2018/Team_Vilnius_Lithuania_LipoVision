#!/usr/bin/env bash

# Deps
pacman -Sy --needed mingw-w64-x86_64-gtk2 mingw-w64-x86_64-toolchain mingw-w64-x86_64-pkg-config

# Fix pkg-config
pushd /c/msys64/mingw64/lib/pkgconfig
sed -i.bak 's|/mingw64|C:/msys64/mingw64/|g' *.pc
popd