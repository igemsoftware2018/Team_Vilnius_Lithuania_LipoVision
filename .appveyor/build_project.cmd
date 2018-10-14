set PATH=%PATH%;C:\msys64\mingw64\bin

mkdir build
mkdir pkg

cd assets
windres -o ..\cmd\lipovision\lipovision-res.syso lipovision.rc

cd ..\cmd\lipovision
go build -ldflags "-w -H windowsgui" -o ../../build/LipoVision.exe

cd ../..
c:\msys64\usr\bin\env MSYSTEM=MINGW64 c:\msys64\usr\bin\bash %APPVEYOR_BUILD_FOLDER%\.appveyor\msys_assemble.sh %APPVEYOR_BUILD_FOLDER%\build