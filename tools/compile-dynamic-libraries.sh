#!/bin/sh

VERSION="7.1.3"

BASE="./c-lib"
BUILD=$BASE/build
BUILD_LINUX=$BUILD-linux
BUILD_WINDOWS=$BUILD-windows
BUILD_MACOS=$BUILD-macos

set -eux

function prepare() {
  rm -r $1 || true
  mkdir -p $1
}

# Linux native.

prepare $BUILD_LINUX

( 
  cd $BUILD_LINUX                                        
  cmake -DBUILD_SHARED_LIBS="ON" ../                     
  make webp                                              
  cp libwebp.so.$VERSION ../../lib/dynamic/webp/blobs/libwebp_amd64.so 
)

# Windows with zig.

prepare $BUILD_WINDOWS

( 
  cd $BUILD_WINDOWS                                                                                           
  cmake -DBUILD_SHARED_LIBS="ON" -DCMAKE_TOOLCHAIN_FILE=../../tools/cross-windows.cmake ../ 
  make webp                                                                                                   
  cp libwebp.dll ../../lib/dynamic/webp/blobs/libwebp_amd64.dll                                                          
)

# macOS with osxcross.

prepare $BUILD_MACOS

(
  cd $BUILD_MACOS                                                                                         
  cmake -DBUILD_SHARED_LIBS="ON" -DCMAKE_TOOLCHAIN_FILE=../../tools/cross-macos.cmake ../ 
  make webp                                                                                               
  cp libwebp.$VERSION.dylib ../../lib/dynamic/webp/blobs/libwebp_amd64.dylib                                         
)

