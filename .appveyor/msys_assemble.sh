#!/usr/bin/env bash

BUILD_FOLDER=${1:-build}

# Get deps
lines=($(ldd  $BUILD_FOLDER/LipoVision.exe | grep lib | cut -d' ' -f3))

for dep in ${lines[*]}
do
    cp $dep $BUILD_FOLDER
done

# Generaate GTK loaders
/mingw64/bin/gdk-pixbuf-query-loaders > /mingw64/lib/gdk-pixbuf-2.0/2.10.0/loaders.cache

mkdir -p $BUILD_FOLDER/lib/gdk-pixbuf-2.0/
mkdir -p $BUILD_FOLDER/share/icons/

cp -r /mingw64/lib/gdk-pixbuf-2.0/2.10.0/ $BUILD_FOLDER/lib/gdk-pixbuf-2.0/
cp -r /mingw64/share/icons/Adwaita/ $BUILD_FOLDER/share/icons/