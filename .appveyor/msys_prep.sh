#!/usr/bin/env bash

# Deps
pacman -Sy --needed \
    mingw-w64-x86_64-gtk3 \
    mingw-w64-x86_64-pkg-config \
    mingw-w64-x86_64-ntldd-git

# Fix pkg-config
pushd /c/msys64/mingw64/lib/pkgconfig
sed -i.bak 's|/mingw64|C:/msys64/mingw64/|g' *.pc
popd