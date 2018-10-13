#!/usr/bin/env bash
# To be executed in root project folder

ASSETS=assets

BUNDLE_FOLDER=LipoVision.app
BIN_FOLDER=$BUNDLE_FOLDER/Contents/MacOS
RES_FOLDER=$BUNDLE_FOLDER/Contents/Resources

# Create bundle
mkdir -p $BIN_FOLDER
mkdir -p $RES_FOLDER

# Copy plist
cp $ASSETS/Info.plist $BUNDLE_FOLDER/Contents

# Make icons
bash scripts/iconset_make.sh $ASSETS
cp $ASSETS/lipovision.icns $RES_FOLDER

pushd cmd/lipovision
go build -ldflags '-s -w' -o ../../$BIN_FOLDER/LipoVision
popd

cp /usr/local/opt/opencv/lib/libopencv_*.3.4.2.dylib $BIN_FOLDER

dylibs=($BIN_FOLDER/libopencv_*)
for lib in "${dylibs[@]}"
do
    install_name_tool -change /usr/local/opt/opencv/lib/$(basename $lib) \
    @executable_path/$(basename $lib) $BIN_FOLDER/LipoVision
done

mv LipoVision.app build/