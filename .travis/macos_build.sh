#!/usr/bin/env bash
# To be executed in root project folder

ASSETS=assets
OPENCV_BUILD=${1:-"/usr/local/opt/opencv/lib"}
OPENCV_POSTFIX=${2:-".3.4.2.dylib"}

BUNDLE_FOLDER=LipoVision.app
BIN_FOLDER=$BUNDLE_FOLDER/Contents/MacOS
RES_FOLDER=$BUNDLE_FOLDER/Contents/Resources

# Create bundle
mkdir -p $BIN_FOLDER
mkdir -p $RES_FOLDER

# Copy plist
cp $ASSETS/Info.plist $BUNDLE_FOLDER/Contents

# Make icons
bash .travis/iconset_make.sh $ASSETS
cp $ASSETS/lipovision.icns $RES_FOLDER

pushd cmd/lipovision
go build -ldflags '-s -w' -o ../../$BIN_FOLDER/LipoVision -tags gtk_3_10
popd

cp template-intersection.png $BIN_FOLDER/
cp $OPENCV_BUILD/libopencv_*$OPENCV_POSTFIX $BIN_FOLDER
cp /usr/local/opt/fontconfig/lib/lib* $BIN_FOLDER
cp /usr/local/opt/freetype/lib/lib* $BIN_FOLDER

dylibs=($BIN_FOLDER/libopencv_*)
for lib in "${dylibs[@]}"
do
    install_name_tool -change $OPENCV_BUILD/$(basename $lib) \
    @executable_path/$(basename $lib) $BIN_FOLDER/LipoVision
done

mv LipoVision.app/ build/LipoVision.app/