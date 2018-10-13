#!/usr/bin/env bash

cd $1
mkdir lipovision.iconset
sips -z 16 16     lipovision.png --out lipovision.iconset/icon_16x16.png
sips -z 32 32     lipovision.png --out lipovision.iconset/icon_16x16@2x.png
sips -z 32 32     lipovision.png --out lipovision.iconset/icon_32x32.png
sips -z 64 64     lipovision.png --out lipovision.iconset/icon_32x32@2x.png
sips -z 128 128   lipovision.png --out lipovision.iconset/icon_128x128.png
sips -z 256 256   lipovision.png --out lipovision.iconset/icon_128x128@2x.png
sips -z 256 256   lipovision.png --out lipovision.iconset/icon_256x256.png
sips -z 512 512   lipovision.png --out lipovision.iconset/icon_256x256@2x.png
sips -z 512 512   lipovision.png --out lipovision.iconset/icon_512x512.png
sips -z 512 512   lipovision.png --out lipovision.iconset/icon_512x512@2x.png
iconutil -c icns lipovision.iconset
rm -rf lipovision.iconset