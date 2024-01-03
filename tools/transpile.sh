#!/bin/sh

set -eux

cd ./c-lib/

# Linux transpile.

NATIVE_CC=$(which gcc)
NATIVE_CXX=$(which c++)
NATIVE_AR=$(which ar)

make distclean || true
./autogen.sh
./configure \
    CC=$NATIVE_CC \
    CXX=$NATIVE_CXX \
    --host=x86_64-linux-gnu \
    --disable-libwebpmux \
    --disable-libwebpdemux \
    --disable-sse4.1 \
    --disable-sse2 \
    --disable-neon \
    --disable-neon-rtcd \
    --disable-threading \
    --disable-gl \
    --disable-sdl \
    --disable-png \
    --disable-jpeg \
    --disable-tiff  \
    --disable-gif  \
    --disable-wic \
    --disable-shared

rm cdb-linux.json || true

CC=$NATIVE_CC \
CXX=$NATIVE_CXX \
AR=$NATIVE_AR \
ccgo -compiledb cdb-linux.json make

CC=$NATIVE_CC \
CXX=$NATIVE_CXX \
AR=$NATIVE_AR \
ccgo -pkgname webp -trace-translation-units -o ../lib/transpiled/webp/libwebp_linux_amd64.go cdb-linux.json src/.libs/libwebp.a

gofmt -s -w ../lib/transpiled/webp/libwebp_linux_amd64.go

# Linux arm64.

LINUX_ARM_HOST=aarch64-linux-gnu
LINUX_ARM_CC=$(which $LINUX_ARM_HOST-gcc) 
LINUX_ARM_CXX=$(which $LINUX_ARM_HOST-c++) 
LINUX_ARM_AR=$(which $LINUX_ARM_HOST-ar) 

make distclean || true
./autogen.sh
./configure \
    CC=$LINUX_ARM_CC \
    CXX=$LINUX_ARM_CXX \
    --host=$LINUX_ARM_HOST \
    --disable-libwebpmux \
    --disable-libwebpdemux \
    --disable-sse4.1 \
    --disable-sse2 \
    --disable-neon \
    --disable-neon-rtcd \
    --disable-threading \
    --disable-gl \
    --disable-sdl \
    --disable-png \
    --disable-jpeg \
    --disable-tiff  \
    --disable-gif  \
    --disable-wic \
    --disable-shared

rm cdb-linux-arm64.json || true

CC=$LINUX_ARM_CC \
CXX=$LINUX_ARM_CXX \
AR=$LINUX_ARM_AR \
ccgo -compiledb cdb-linux-arm64.json make

CC=$LINUX_ARM_CC \
CXX=$LINUX_ARM_CXX \
AR=$LINUX_ARM_AR \
ccgo -pkgname webp -trace-translation-units -o ../lib/transpiled/webp/libwebp_linux_arm64.go cdb-linux-arm64.json src/.libs/libwebp.a

gofmt -s -w ../lib/transpiled/webp/libwebp_linux_arm64.go

# macOS transpile.

MACOS_HOST=x86_64-apple-darwin15
MACOS_CC=$(which $MACOS_HOST-cc) 
MACOS_CXX=$(which $MACOS_HOST-c++) 
MACOS_AR=$(which $MACOS_HOST-ar) 

make distclean || true
./autogen.sh
./configure \
    CC=$MACOS_CC \
    CXX=$MACOS_CXX \
    --host=$MACOS_HOST \
    --disable-libwebpmux \
    --disable-libwebpdemux \
    --disable-sse4.1 \
    --disable-sse2 \
    --disable-neon \
    --disable-neon-rtcd \
    --disable-threading \
    --disable-gl \
    --disable-sdl \
    --disable-png \
    --disable-jpeg \
    --disable-tiff  \
    --disable-gif  \
    --disable-wic \
    --disable-shared

rm cdb-darwin.json || true

CC=$MACOS_CC \
CXX=$MACOS_CXX \
AR=$MACOS_AR \
ccgo -compiledb cdb-darwin.json make

CC=$MACOS_CC \
CXX=$MACOS_CXX \
AR=$MACOS_AR \
ccgo -pkgname webp -trace-translation-units -o ../lib/transpiled/webp/libwebp_darwin_amd64.go cdb-darwin.json src/.libs/libwebp.a

gofmt -s -w ../lib/transpiled/webp/libwebp_darwin_amd64.go

# Windows transpile.

WINDOWS_HOST=x86_64-w64-mingw32
WINDOWS_CC=$(which x86_64-w64-mingw32-gcc)
WINDOWS_CXX=$(which x86_64-w64-mingw32-g++)
WINDOWS_AR=$(which x86_64-w64-mingw32-ar)

make distclean || true
./autogen.sh
./configure \
    CC=$WINDOWS_CC \
    CXX=$WINDOWS_CXX \
    --host=$WINDOWS_HOST \
    --disable-libwebpmux \
    --disable-libwebpdemux \
    --disable-sse4.1 \
    --disable-sse2 \
    --disable-neon \
    --disable-neon-rtcd \
    --disable-threading \
    --disable-gl \
    --disable-sdl \
    --disable-png \
    --disable-jpeg \
    --disable-tiff  \
    --disable-gif  \
    --disable-wic \
    --disable-shared

rm cdb-windows.json || true

CC=$WINDOWS_CC \
CXX=$WINDOWS_CXX \
AR=$WINDOWS_AR \
ccgo -compiledb cdb-windows.json make

CC=$WINDOWS_CC \
CXX=$WINDOWS_CXX \
AR=$WINDOWS_AR \
ccgo -windows -pkgname webp -trace-translation-units -o ../lib/transpiled/webp/libwebp_windows_amd64.go cdb-windows.json src/.libs/libwebp.a

gofmt -s -w ../lib/transpiled/webp/libwebp_windows_amd64.go


